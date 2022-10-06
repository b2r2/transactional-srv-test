package model

import "time"

// Client ...
type Client struct {
	ID        uint64     `db:"id"`
	Username  string     `db:"username"`
	Balance   *int64     `db:"balance"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

// ClientTableName ...
const ClientTableName = "client"
