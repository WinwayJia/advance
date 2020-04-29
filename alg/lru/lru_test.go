package lru

import (
	"testing"
	"time"
)

const (
	cacheItemSize = 4
)

func TestLru(t *testing.T) {
	cache := New(cacheItemSize, time.Second*3)

	for i := 0; i <= 10; i++ {
		cache.Put(i, i)
	}

	if cache.Len() != cacheItemSize {
		t.Fatalf("invalid len: %d %d\n", cache.Len(), cacheItemSize)
	}
	cache.Range(func(k, v interface{}) {
		t.Logf("xxxx key: %v value: %v\n", k, v)
	})

	for i := 0; i < cacheItemSize; i++ {
		key := 10 - i
		value, err := cache.Get(key)
		if err != nil {
			t.Fatalf("Get error, %v %v\n", key, err)
		}
		t.Logf("key: %v value: %v\n", key, value)
		cache.Erase(key)
	}

	if cache.Len() != 0 {
		t.Fatalf("invalid len: %d\n", cache.Len())
	}
}

func TestLruExpired(t *testing.T) {
	cache := New(cacheItemSize, time.Second*3)
	cache.Put(1, 2)
	if cache.Len() != 1 {
		t.Fatalf("invalid len\n")
	}

	v, err := cache.Get(1)
	t.Logf("value: %v, err: %v\n", v, err)

	time.Sleep(3 * time.Second)
	v, err = cache.Get(1)
	t.Logf("value: %v, err: %v\n", v, err)

	time.Sleep(time.Second) // ensure async erase occured
	if cache.Len() != 0 {
		t.Fatalf("invalid len %d\n", cache.Len())
	}
}

func TestParallel(t *testing.T) {
	cache := New(cacheItemSize, time.Second*3)

	go func() {
		cache.Put(1, 2)
	}()
}
