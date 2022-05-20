package cache

import (
	"time"
)

type cacheValue struct {
	value string
	valid time.Time
}

type Cache struct {
	cache     map[string]cacheValue
	timelimit int
}

func NewCache() Cache {
	return Cache{cache: make(map[string]cacheValue), timelimit: 300}
}

func (this *Cache) Valid(key string) bool {
	cvalue, exists := this.cache[key]

	if exists && cvalue.valid.Sub(time.Now()).Seconds() > 0.0 {
		return true
	}

	return false
}

func (this *Cache) Get(key string) (string, bool) {
	cvalue, exists := this.cache[key]

	if exists && cvalue.valid.Sub(time.Now()).Seconds() > 0.0 {
		return cvalue.value, true
	}

	return "", false
}

func (this *Cache) Put(key, value string) {
	this.cache[key] = cacheValue{value: value, valid: time.Now().Add(time.Second * time.Duration(this.timelimit))}
}

func (this *Cache) Keys() []string {
	keys := make([]string, 0, len(this.cache))

	for key := range this.cache {
		if this.Valid(key) {
			keys = append(keys, key)
		}
	}

	return keys
}

func (this *Cache) PutTill(key, value string, deadline time.Time) {
	this.cache[key] = cacheValue{value: value, valid: deadline}
}
