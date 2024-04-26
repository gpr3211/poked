package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.Cachebody == nil {
		t.Error("cache not created")
	}

}

func TestGetCache(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key not found")
	}
	if string(actual) != "val1" {
		t.Error("not found")
	}
}
