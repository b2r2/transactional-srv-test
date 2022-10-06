package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/b2r2/transactional-srv-test/internal/app/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Balance ...
func (s *Service) Balance(ctx context.Context, id uint64) (*model.Client, error) {
	client, err := s.r.Balance(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "client not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return client, nil
}
