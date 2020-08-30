package cacher

import (
	"testing"
	"time"
)

func TestSetCacheItem(t *testing.T) {
	expiredAtWithDurationMoreThenMinute := time.Now().Add(2 * time.Minute)
	expiredAtWithDurationLessThenMinute := time.Now().Add(59 * time.Second)

	tests := []struct {
		name               string
		key                string
		value              interface{}
		expiredAt          *time.Time
		expectedCachedItem interface{}
	}{
		{"+ve:ShouldCacheItemWithExpiredAt", "name", "neel001", &expiredAtWithDurationMoreThenMinute, "neel001"},
		{"+ve:ShouldCacheItemWithoutExpiredAt", "name", "neel002", nil, "neel002"},
		{"-ve:ShouldNotCacheItemWithoutExpiredAtLessThanMinute", "name", "neel003", &expiredAtWithDurationLessThenMinute, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SetCacheItem(tt.key, tt.value, tt.expiredAt)
			if tt.expectedCachedItem != res {
				t.Errorf("expected value of cached item %v, got %v", tt.expectedCachedItem, res)
			}
		})
	}
}

func TestGetCacheItem(t *testing.T) {
	expiredAtWithDurationMoreThenMinute := time.Now().Add(2 * time.Minute)
	key1 := "name"
	value1 := "neel001"
	SetCacheItem(key1, value1, &expiredAtWithDurationMoreThenMinute)

	key2 := "age"
	value2 := 26
	SetCacheItem(key2, value2, nil)

	expiredAtWithDurationLessThenMinute := time.Now().Add(59 * time.Second)
	key3 := "random"
	value3 := "randomValue"
	SetCacheItem(key3, value3, &expiredAtWithDurationLessThenMinute)

	nonExistingKey := "nonExistingKey"

	tests := []struct {
		name               string
		key                string
		expectedCachedItem interface{}
	}{
		{"+ve:ShouldGetCachedItemWhileExpiredAtSet", key1, value1},
		{"+ve:ShouldGetCachedItemWhileExpiredAtSetNull", key2, value2},
		{"-ve:ShouldNotGetCachedItemWhileExpiredAtSetLessThanMinute", key3, nil},
		{"-ve:ShouldNotGetNonExistingCachedItem", nonExistingKey, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GetCachedItem(tt.key)
			if tt.expectedCachedItem != res {
				t.Errorf("expected value of cached item %v, got %v", tt.expectedCachedItem, res)
			}
		})
	}
}

func TestDeleteCacheItem(t *testing.T) {

	expiredAtWithDurationMoreThenMinute := time.Now().Add(2 * time.Minute)
	cachedItem := SetCacheItem("name", "neel001", &expiredAtWithDurationMoreThenMinute)

	tests := []struct {
		name               string
		key                string
		expectedCachedItem interface{}
	}{
		{"+ve:ShouldDeleteCachedItem", "name", cachedItem},
		{"+ve:ShouldNotFailWhenUsedNonExistingKey", "name1", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DeleteCachedItem(tt.key)
			if tt.expectedCachedItem != res {
				t.Errorf("expected value of cached item %v, got %v", tt.expectedCachedItem, res)
			}
		})
	}
}
