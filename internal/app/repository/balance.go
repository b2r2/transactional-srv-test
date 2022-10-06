package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/b2r2/transactional-srv-test/internal/app/model"
)

// Balance ...
func (r *repository) Balance(ctx context.Context, clientID uint64) (*model.Client, error) {
	query, args, err := sq.Select("*").From(model.ClientTableName).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": clientID}).ToSql()
	if err != nil {
		return nil, err
	}
	var c []*model.Client
	if err = r.db.SelectContext(ctx, &c, query, args...); err != nil {
		return nil, err
	}
	if len(c) == 0 {
		return nil, sql.ErrNoRows
	}
	return c[0], nil
}
