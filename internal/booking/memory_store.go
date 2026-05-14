package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (m *MemoryStore) Book(b Booking) error {
	if _, exists := m.bookings[b.ID]; exists {
		return ErrSeatAlreadyExists
	}
	m.bookings[b.ID] = b
	return nil
}
func (m *MemoryStore) ListBookings(movieID string) []Booking {
	var bookings []Booking
	for _, b := range m.bookings {
		if b.ID == movieID {
			bookings = append(bookings, b)
		}
	}
	return bookings
}
