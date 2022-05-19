package cache

import "time"

type Cache struct {
	key   string
	value string

	deadline time.Time
}

func   NewCache() Cache {
	return Cache{}
}

func (cache Cache) Get(key string) (string, bool) {
	if cache.key == key && cache.deadline.After(time.Now()) {
		return cache.value, true
	}
	return "", false

}

func (cache *Cache) Put(key, value string) {

	cache.key = key
	cache.value = value

}

func (cache Cache) Keys() []string {
	if cache.deadline.After(time.Now()) {
		return []string{}
	}
	return []string{cache.key}
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.key = key
	cache.value = value
	cache.deadline = deadline

}
