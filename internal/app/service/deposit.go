package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/b2r2/transactional-srv-test/internal/app/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Deposit ...
func (s *Service) Deposit(ctx context.Context, clientID, amount uint64) error {
	client, err := s.Balance(ctx, clientID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return status.Error(codes.NotFound, "client not found")
		}
		return status.Error(codes.Internal, err.Error())
	}
	balance := *client.Balance
	eventID, err := s.r.AddEvent(ctx, model.NewEvents(client.ID,
		amount,
		balance,
		s.EventsTypeIDs.Deposit,
		s.EventsStatusIDs.Processing))
	if err != nil {
		return err
	}
	nb := balance + int64(amount)
	err = s.r.BalanceProcessing(ctx, client.ID, nb)
	if err != nil {
		_, errEvent := s.r.AddEvent(ctx, model.NewEvents(client.ID,
			amount,
			balance,
			s.EventsTypeIDs.Deposit,
			s.EventsStatusIDs.Deferred))
		if errEvent != nil {
			return fmt.Errorf("%s: %w", err.Error(), errEvent)
		}
		return err
	}
	return s.r.RemoveEvent(ctx, eventID)
}
