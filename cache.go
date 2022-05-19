package cache

import "time"

type Cache struct {
	key      string
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{}
}

func (cache Cache) Get(key string) (string, bool) {
	if cache.key == key && cache.deadline.After(time.Now()) {
		return cache.value, true
	}
	return "", false

}

func (cache *Cache) Put(key, value string) {
	if cache.key == key {
		cache.value = value

		return
	} else {
		cache.key = key
		cache.value = value
		*cache = NewCache()
	}

}

func (cache *Cache) Keys() []string {
	var keys []string
	if cache.deadline.After(time.Now()) {
		keys = append(keys, cache.key)
	}
	return []string{}
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.key = key
	cache.value = value
	cache.deadline = deadline

}
