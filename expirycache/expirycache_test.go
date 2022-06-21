package expirycache

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewInt64(t *testing.T) {
	c := NewInt64(time.Second, func() time.Time {
		return time.Time{}
	}, func() (int64, error) {
		return 0, nil
	})
	assert.NotNil(t, c.fetch)
	assert.NotNil(t, c.now)
	assert.NotNil(t, c.cache)
	assert.Zero(t, c.lastUpdate)
}

func TestInt64_Get(t *testing.T) {
	i := 0
	samples := []int64{11, 22, 33}
	clock := time.Now()
	now := func() time.Time {
		return clock
	}
	c := NewInt64(time.Minute, now, func() (int64, error) {
		v := samples[i]
		i++
		return v, nil
	})

	v, _ := c.Get()
	assert.Equal(t, int64(11), v)

	clock = clock.Add(time.Second * 30)
	v, _ = c.Get()
	assert.Equal(t, int64(11), v)

	clock = clock.Add(time.Second * 30)
	v, _ = c.Get()
	assert.Equal(t, int64(22), v)

	clock = clock.Add(time.Second * 60)
	v, _ = c.Get()
	assert.Equal(t, int64(33), v)
}

func TestInt64_GetErr(t *testing.T) {
	clock := time.Now()
	now := func() time.Time {
		return clock
	}
	c := NewInt64(time.Minute, now, func() (int64, error) {
		return 0, errors.New("some err")
	})

	_, err := c.Get()
	assert.NotNil(t, err)

	*c.cache = 17
	v, err := c.Get()
	assert.NotNil(t, err)
	assert.Equal(t, int64(17), v)

	c.fetch = func() (int64, error) {
		return 18, nil
	}
	v, err = c.Get()
	assert.Nil(t, err)
	assert.Equal(t, int64(18), v)
}
