package transactional

import (
	"context"

	desc "github.com/b2r2/transactional-srv-test/pkg/api/v1/transactional"
)

// Balance ...
func (i *Implementation) Balance(ctx context.Context, req *desc.BalanceRequest) (*desc.BalanceResponse, error) {
	res, err := i.s.Balance(ctx, req.GetClientId())
	if err != nil {
		return nil, err
	}
	resp := &desc.BalanceResponse{Username: res.Username}
	if res.Balance != nil {
		resp.Balance = *res.Balance
	}
	return resp, nil
}
