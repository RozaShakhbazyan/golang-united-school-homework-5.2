package cache

import "time"

type Cache struct {
	key      string
	value    string
	deadline time.Time
}

var keys []string

func NewCache() Cache {
	return Cache{}
}

func (cache Cache) Get(key string) (string, bool) {
	if cache.key == key {
		return cache.value, true
	} else {
		return "", false
	}

}

func (cache *Cache) Put(key, value string) {
	if cache.key == key {
		cache.value = value

	} else {
		cache.key = key
		cache.value = value
		*cache = NewCache()
	}
	keys = append(keys, cache.key)

}

func (cache *Cache) Keys() []string {

	if cache.deadline.After(time.Now()) {
		keys = append(keys, cache.key)
	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.key = key
	cache.value = value
	cache.deadline = deadline

}
