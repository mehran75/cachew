package memorycache

import (
	"sync"
	"time"
)

type dataWrapper[T any] struct {
	data     T
	expireAt time.Time
}

type Cache[Key comparable, Value any] struct {
	data map[Key]dataWrapper[Value]
	lock sync.RWMutex
}
