package cache

import "time"

type Cache struct {
	key      string
	value    string
	expValue string
	deadline time.Time
	expired  bool
}

func NewCache(key string,
	value string,
	expValue string,
	deadline time.Time,
	expired bool) Cache {
	return Cache{key: key, value: value, expValue: expValue, deadline: deadline, expired: expired}
}

func (cache Cache) Get(key string) (string, bool) {
	if cache.key == key && cache.expired {
		return cache.value, true
	}
	return "", false

}

func (cache *Cache) Put(key, value string) {

	cache.key = key
	cache.value = value
	cache.expired = false
}

func (cache Cache) Keys() []string {
	if cache.expired {
		return []string{}
	}
	return []string{cache.key}
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.key = key
	cache.value = value
	cache.deadline = deadline
	cache.expired = deadline.After(time.Now())

}
