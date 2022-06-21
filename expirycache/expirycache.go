package expirycache

import (
	"sync"
	"time"
)

type Int64 struct {
	expiry     time.Duration
	now        func() time.Time
	fetch      func() (int64, error)
	cache      *int64
	lastUpdate time.Time
	mu         sync.Mutex
}

// NewInt64 returns a cacher for an int64 value where expiry is the max duration for which a value will remain valid,
// now is the clock function and fetch is used to retrieve a new sample.
func NewInt64(expiry time.Duration, now func() time.Time, fetch func() (int64, error)) *Int64 {
	return &Int64{
		expiry: expiry,
		now:    now,
		fetch:  fetch,
		cache:  new(int64),
	}
}

func (i *Int64) update(now time.Time) error {
	v, err := i.fetch()
	if err != nil {
		return err
	}

	*(i.cache) = v
	i.lastUpdate = now
	return nil
}

// Get returns a possibly cached value for the fetch function.
func (i *Int64) Get() (int64, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	now := i.now()

	if expired := now.Sub(i.lastUpdate) >= i.expiry; !expired {
		return *(i.cache), nil
	}

	err := i.update(now)
	return *(i.cache), err
}
