package mapsecure

import (
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{})}
}

func (sm *SafeMap) Set(key string, value interface{}) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (interface{}, bool) {
	sm.mu.RLock()
	value, ok := sm.m[key]
	sm.mu.RUnlock()
	return value, ok
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}
