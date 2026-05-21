package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const defaultHoldTTL = 2 * time.Minute

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{rdb: rdb}
}

func sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func (s *RedisStore) Book(b Booking) error {
	booking, err := s.hold(b)
	if err != nil {
		return err
	}

	s
	return nil
}
func (s *RedisStore) ListBookings(movieID string) []Booking {
	return []Booking{}
}

func (s *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := fmt.Sprintf("seat:%s:%s", b.MovieID, b.SeatID)
	b.ID = id
	val, err := json.Marshal(b)
	if err != nil {
		return Booking{}, err
	}

	res := s.rdb.SetArgs(ctx, key, val, redis.SetArgs{
		TTL:  defaultHoldTTL,
		Mode: "NX",
	})

	ok := res.Val() == "OK"
	if !ok {
		return Booking{}, ErrSeatAlreadyExists
	}

	s.rdb.Set(ctx, sessionKey(id), val, defaultHoldTTL)

	booking := Booking{
		MovieID:   b.MovieID,
		SeatID:    b.SeatID,
		ID:        id,
		Status:    "held",
		UserID:    b.UserID,
		ExpiresAt: now.Add(defaultHoldTTL),
	}
	return booking, nil

}
