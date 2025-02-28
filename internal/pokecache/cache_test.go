package pokecache

import (
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	cache := NewCache(time.Minute)
	cases := []string{
		"test1",
		"test2",
		"test3",
	}
	for _, key := range cases {
		data := []byte(key + "_data")
		cache.Add(key, data)
		result, exists := cache.Get(key)
		if !exists {
			t.Errorf("Expected key %s to exist in cache", key)
		}
		if string(result) != string(data) {
			t.Errorf("Expected data %s, got %s", string(data), string(result))
		}
	}
}

func TestCacheExpiry(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)
	key := "test_key"
	data := []byte("test_data")
	cache.Add(key, data)
	time.Sleep(time.Millisecond * 200)
	_, exists := cache.Get(key)
	if exists {
		t.Error("Expected key to be expired")
	}
}

func TestCacheGet(t *testing.T) {
	cache := NewCache(time.Minute)
	key := "test_key"
	_, exists := cache.Get(key)
	if exists {
		t.Error("Expected non-existent key to return false")
	}
}
