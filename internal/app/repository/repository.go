package repository

import (
	"context"

	"github.com/b2r2/transactional-srv-test/internal/app/model"
	"github.com/jmoiron/sqlx"
)

// IRepository ...
type IRepository interface {
	AddUser(ctx context.Context, username string, balance *int64) (uint64, error)
	BalanceProcessing(ctx context.Context, clientID uint64, newBalance int64) error
	Balance(ctx context.Context, clientID uint64) (*model.Client, error)
	AddEvent(ctx context.Context, event model.Events) (uint32, error)
	RemoveEvent(ctx context.Context, id uint32) error
	ListEventsTypeIDs(ctx context.Context) (m []*model.DictEntry, err error)
	ListEventsStatusIDs(ctx context.Context) (m []*model.DictEntry, err error)
	SetDictionaryIDs(dest interface{}, dicts []*model.DictEntry) error
}

type repository struct {
	db *sqlx.DB
}

// NewRepository ...
func NewRepository(ctx context.Context, db *sqlx.DB) IRepository {
	return &repository{db: db}
}
