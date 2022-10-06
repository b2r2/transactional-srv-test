package model

// DictEntry ...
type DictEntry struct {
	ID    uint32 `db:"id"`
	Code  string `db:"code"`
	Title string `db:"title"`
}

// Events ...
type Events struct {
	ID       uint64 `db:"id"`
	ClientID uint64 `db:"client_id"`
	Balance  int64  `db:"balance"`
	Amount   uint64 `db:"amount"`
	TypeID   uint32 `db:"type_id"`
	StatusID uint32 `db:"status_id"`
}

// EventsTable ...
const EventsTable = "events"

// NewEvents ...
func NewEvents(clientID, amount uint64, balance int64, typeID, statusID uint32) Events {
	return Events{
		ClientID: clientID,
		Balance:  balance,
		Amount:   amount,
		TypeID:   typeID,
		StatusID: statusID,
	}
}
