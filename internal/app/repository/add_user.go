package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/b2r2/transactional-srv-test/internal/app/model"
)

// AddUser ...
func (r *repository) AddUser(ctx context.Context, username string, balance *int64) (id uint64, err error) {
	builder := sq.Insert(model.ClientTableName).PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")
	if balance != nil {
		builder = builder.Columns("username", "balance").Values(username, balance)
	} else {
		builder = builder.Columns("username").Values(username)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building query: %w", err)
	}
	if err = r.db.QueryRowxContext(ctx, query, args...).Scan(&id); err != nil {
		return 0, fmt.Errorf("error execution: %w", err)
	}
	return
}
