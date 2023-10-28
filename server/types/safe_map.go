package types

import "sync"

type SafeMap[V any] struct {
	items map[string]V
	mu    sync.Mutex
}

func NewSafeMap[V any]() *SafeMap[V] {
	return &SafeMap[V]{items: make(map[string]V)}
}

func (sm *SafeMap[V]) Get(key string) V {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.items[key]
}

func (sm *SafeMap[V]) Set(key string, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.items[key] = value
}

func (sm *SafeMap[V]) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.items, key)
}

func (sm *SafeMap[V]) Update(key string, fn func(V)) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	item := sm.items[key]
	fn(item)
	sm.items[key] = item
}
