package transactional

import (
	"context"

	desc "github.com/b2r2/transactional-srv-test/pkg/api/v1/transactional"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Withdraw ...
func (i *Implementation) Withdraw(ctx context.Context, req *desc.WithdrawRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, i.s.Withdraw(ctx, req.GetClientId(), req.GetAmount())
}
