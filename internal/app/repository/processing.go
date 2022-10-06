package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/b2r2/transactional-srv-test/internal/app/model"
)

// BalanceProcessing ...
func (r *repository) BalanceProcessing(ctx context.Context, clientID uint64, newBalance int64) error {
	builder := sq.Update(model.ClientTableName).PlaceholderFormat(sq.Dollar).
		Set("balance", newBalance).Where(sq.Eq{"id": clientID})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}
