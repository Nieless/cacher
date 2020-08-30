package cacher

import (
	"sync"
	"time"
)

const expiryDiff = 1

// is type to store cache values
type Cache map[string]*CacheItem

// is type of cache value
type CacheItem struct {
	Item      interface{}
	ExpiresAt *time.Time
}

var (
	cache = Cache{}
	lock  = sync.Mutex{}
)

// SetCacheItem cache the item by taking key, value as an arguments
func SetCacheItem(key string, val interface{}, expiresAt *time.Time) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if expiresAt != nil {
		// return if cached item is about to expire
		if expiresAt.Sub(time.Now()).Minutes() < expiryDiff {
			return nil
		}
	}

	cache[key] = &CacheItem{
		Item:      val,
		ExpiresAt: expiresAt,
	}
	return cache[key].Item
}

// GetCachedItem returns the cached item by taking key as an argument
func GetCachedItem(key string) interface{} {
	lock.Lock()
	defer lock.Unlock()

	item := cache[key]
	if item == nil {
		return nil
	}

	if item.ExpiresAt == nil {
		return item.Item
	}

	// delete cached item if it is about to expire
	if item.ExpiresAt.Sub(time.Now()).Minutes() < expiryDiff {
		delete(cache, key)
		return nil
	}

	return item.Item
}

// DeleteCachedItem deletes the cached item by taking key as as an argument
func DeleteCachedItem(key string) interface{} {
	lock.Lock()
	defer lock.Unlock()

	if cache[key] == nil {
		return nil
	}

	deletedItem := cache[key]
	delete(cache, key)
	return deletedItem.Item
}
