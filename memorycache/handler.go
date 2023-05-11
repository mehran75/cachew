package memorycache

import (
	"errors"
	"time"
)

func NewCache[Key comparable, Value any]() *Cache[Key, Value] {
	return &Cache[Key, Value]{
		data: make(map[Key]dataWrapper[Value]),
	}
}

func (c *Cache[Key, Value]) Set(key Key, value Value, expiresAfter time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = dataWrapper[Value]{
		data:     value,
		expireAt: time.Now().Add(expiresAfter),
	}
}

func (c *Cache[Key, Value]) Get(key Key) (*Value, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	data, ok := c.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	if data.expireAt.Before(time.Now()) {
		return nil, errors.New("the key has expired")
	}

	return &data.data, nil
}

func (c *Cache[Key, Value]) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data = make(map[Key]dataWrapper[Value])
}
