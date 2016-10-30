package demo1

import "sync"

type MemHashTable struct {
	m map[string][]byte
	lock sync.RWMutex
}

func NewMemHashTable() HashTable {
	return &MemHashTable{m: make(map[string][]byte)}
}

func (h *MemHashTable) Get(key string) ([]byte, error){
	h.lock.RLock()
	defer h.lock.RUnlock()
	val, ok := h.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (h *MemHashTable) Set(key string, val []byte) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.m[key] = val
	return nil
}