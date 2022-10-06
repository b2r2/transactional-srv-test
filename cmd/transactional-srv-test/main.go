package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/b2r2/transactional-srv-test/internal/app/repository"
	"github.com/b2r2/transactional-srv-test/internal/app/service"
	"github.com/b2r2/transactional-srv-test/internal/app/transactional"
	"github.com/b2r2/transactional-srv-test/internal/config"
	pb "github.com/b2r2/transactional-srv-test/pkg/api/v1/transactional"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func run(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("run(): %w", err)
		}
	}()

	db := config.Get().GetDBConnection()
	defer config.Get().Close()

	s, err := service.NewService(ctx, repository.NewRepository(ctx, db))
	if err != nil {
		return err
	}

	return listen(ctx, transactional.New(s))
}

func main() {
	ctx := context.Background()
	if err := config.Load(); err != nil {
		log.Fatalf("can't get load config: %v", err)
	}
	if err := run(ctx); err != nil {
		log.Fatalf("failed run application: %v", err)
	}
}

func listen(ctx context.Context, service *transactional.Implementation) error {
	addr := config.GetEnv("HTTP_ADDR", "127.0.0.1:8845")
	tcpAddr := config.GetEnv("TCP_ADDR", "127.0.0.1:8848")
	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		return fmt.Errorf("can't get tcp connection: %w", err)
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	pb.RegisterTransactionalServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	var group errgroup.Group
	group.Go(func() error {
		return grpcServer.Serve(lis)
	})

	group.Go(func() error {
		mux := runtime.NewServeMux(
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true}}),
		)

		if err = pb.RegisterTransactionalServiceHandlerServer(ctx, mux, service); err != nil {
			return err
		}

		return http.ListenAndServe(addr, mux)
	})
	log.Println("start application")
	return group.Wait()
}
