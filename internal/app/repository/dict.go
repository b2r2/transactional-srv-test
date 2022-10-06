package repository

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/b2r2/transactional-srv-test/internal/app/model"
)

// ListEventsTypeIDs ...
func (r *repository) ListEventsTypeIDs(ctx context.Context) (m []*model.DictEntry, err error) {
	builder := sq.Select("id", "code", "title").
		PlaceholderFormat(sq.Dollar).
		From("events_types").OrderBy("id ASC")
	query, _, err := builder.ToSql()
	if err != nil {
		return
	}
	return r.query(query)

}

// ListEventsStatusIDs ...
func (r *repository) ListEventsStatusIDs(ctx context.Context) (m []*model.DictEntry, err error) {
	builder := sq.Select("id", "code", "title").
		PlaceholderFormat(sq.Dollar).
		From("events_statuses").OrderBy("id ASC")
	query, _, err := builder.ToSql()
	if err != nil {
		return
	}
	return r.query(query)
}

// SetDictionaryIDs ...
func (r *repository) SetDictionaryIDs(dest interface{}, dicts []*model.DictEntry) error {
	codes := make(map[string]uint32)
	for _, ds := range dicts {
		codes[ds.Code] = ds.ID
	}
	dstValue := reflect.ValueOf(dest)
	if dstValue.Kind() == reflect.Ptr {
		dstValue = dstValue.Elem()
	}
	dstType := dstValue.Type()
	for i := 0; i < dstType.NumField(); i++ {
		name := strings.ToUpper(dstType.Field(i).Name)
		id, ok := codes[name]
		if !ok {
			return fmt.Errorf("code not found")
		}
		field := dstValue.Field(i)
		if field.CanSet() {
			field.SetUint(uint64(id))
		}
	}
	return nil
}

func (r *repository) query(query string) (m []*model.DictEntry, err error) {
	rows, err := r.db.Queryx(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for i := 0; rows.Next(); i++ {
		var d model.DictEntry
		err = rows.Scan(&d.ID, &d.Code, &d.Title)
		if err != nil {
			return
		}
		m = append(m, &d)
	}
	return
}
