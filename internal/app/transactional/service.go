package transactional

import (
	"github.com/b2r2/transactional-srv-test/internal/app/service"
	desc "github.com/b2r2/transactional-srv-test/pkg/api/v1/transactional"
)

// Implementation ...
type Implementation struct {
	s *service.Service
	desc.UnimplementedTransactionalServiceServer
}

// New ...
func New(s *service.Service) *Implementation {
	return &Implementation{
		s:                                       s,
		UnimplementedTransactionalServiceServer: desc.UnimplementedTransactionalServiceServer{},
	}
}
