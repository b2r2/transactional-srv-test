package service

import (
	"context"
)

// AddUser ...
func (s *Service) AddUser(ctx context.Context, username string, balance *int64) (uint64, error) {
	return s.r.AddUser(ctx, username, balance)
}
