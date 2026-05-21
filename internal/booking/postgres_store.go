package booking

import "database/sql"

type postgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) BookingStore {
	return postgresStore{db: db}
}

func (s *postgresStore) Book(b Booking) error {

}

func (s *postgresStore) ListBookings(movieID string) []Booking {

}
