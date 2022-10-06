package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/b2r2/transactional-srv-test/internal/app/model"
)

// AddEvent ...
func (r *repository) AddEvent(ctx context.Context, event model.Events) (uint32, error) {
	query, args, err := squirrel.Insert(model.EventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns("type_id", "status_id", "client_id", "amount", "balance").
		Values(event.TypeID, event.StatusID, event.ClientID, event.Amount, event.Balance).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}
	var id uint32
	err = r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	return id, err
}

// RemoveEvent ...
func (r *repository) RemoveEvent(ctx context.Context, id uint32) error {
	query, args, err := squirrel.Delete(model.EventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}
