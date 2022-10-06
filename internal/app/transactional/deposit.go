package transactional

import (
	"context"
	"fmt"

	desc "github.com/b2r2/transactional-srv-test/pkg/api/v1/transactional"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Deposit ...
func (i *Implementation) Deposit(ctx context.Context, req *desc.DepositRequest) (*emptypb.Empty, error) {
	if req.GetAmount() == 0 {
		return nil, fmt.Errorf("amount must be > 0")
	}
	return &emptypb.Empty{}, i.s.Deposit(ctx, req.GetClientId(), req.GetAmount())
}
