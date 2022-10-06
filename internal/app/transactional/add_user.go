package transactional

import (
	"context"

	desc "github.com/b2r2/transactional-srv-test/pkg/api/v1/transactional"
)

// AddUser ...
func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	id, err := i.s.AddUser(ctx, req.GetUsername(), req.Balance)
	return &desc.AddUserResponse{ClientId: id}, err
}
