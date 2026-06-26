package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nill")
	}
}

func TestReap(t *testing.T) {

	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyone := "key1"
	cache.Add(keyone, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyone)
	if !ok {
		t.Errorf("%s should not have been reaped", keyone)
	}

}
