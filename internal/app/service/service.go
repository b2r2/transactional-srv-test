package service

import (
	"context"

	"github.com/b2r2/transactional-srv-test/internal/app/repository"
)

// Service ...
type Service struct {
	r               repository.IRepository
	EventsTypeIDs   struct{ Deposit, Withdraw uint32 }
	EventsStatusIDs struct{ Deferred, Processing uint32 }
}

// NewService ...
func NewService(ctx context.Context, r repository.IRepository) (*Service, error) {
	s := &Service{r: r}
	eventsTypeIDs, err := r.ListEventsTypeIDs(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.SetDictionaryIDs(&s.EventsTypeIDs, eventsTypeIDs); err != nil {
		return nil, err
	}

	eventsStatusIDs, err := r.ListEventsStatusIDs(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.SetDictionaryIDs(&s.EventsStatusIDs, eventsStatusIDs); err != nil {
		return nil, err
	}
	return s, nil
}
