package cache

import "sync"

/*
NOTE: naming convenstion sometime we don't use name with same package
cache.Caches IMHO it's ok for interface
cache.MemCache IMHO it's ok for struct type
*/

type memCache struct {
	m map[string]string
}

func (m *memCache) Set(k string, v string) {
	m.m[k] = v
}

func (m *memCache) Get(k string) string {
	return m.m[k]
}

func New() *memCache {
	return &memCache{
		m: make(map[string]string),
	}
}

type memCacheSync struct {
	m sync.Map
}

func (m *memCacheSync) Set(k string, v string) {
	m.m.Store(k, v)
}

func (m *memCacheSync) Get(k string) string {
	v, ok := m.m.Load(k)
	if !ok {
		// need to do something here
		// return error or log - depend on your design. (IMHO: better to return error )
		return ""
	}

	return v.(string)
}

func NewSync() *memCacheSync {
	return &memCacheSync{
		m: sync.Map{},
	}
}
