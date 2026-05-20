package booking

import "sync"

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: make(map[string]Booking),
		RWMutex:  sync.RWMutex{},
	}
}

func (m *ConcurrentStore) Book(b Booking) error {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.bookings[b.ID]; exists {
		return ErrSeatAlreadyExists
	}
	m.bookings[b.ID] = b
	return nil
}
func (m *ConcurrentStore) ListBookings(movieID string) []Booking {
	m.RLock()
	defer m.RUnlock()
	var bookings []Booking
	for _, b := range m.bookings {
		if b.ID == movieID {
			bookings = append(bookings, b)
		}
	}
	return bookings
}
