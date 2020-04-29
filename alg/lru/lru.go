package lru

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type valueItem struct {
	key   interface{}
	value interface{}
	ts    time.Time
}

// Lru lru struct
type Lru struct {
	cap      int
	size     int
	duration time.Duration
	queue    *list.List
	dict     map[interface{}]*list.Element
	lock     sync.RWMutex
	ch       chan interface{}
}

// New new lru object
func New(cap int, duration time.Duration) *Lru {
	cache := &Lru{
		size:     0,
		cap:      cap,
		queue:    list.New(),
		dict:     make(map[interface{}]*list.Element, cap),
		duration: duration,
		lock:     sync.RWMutex{},
		ch:       make(chan interface{}, 128),
	}

	go cache.doErase()

	return cache
}

// Put put item
func (l *Lru) Put(key, value interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	elem, ok := l.dict[key]
	if ok {
		elem.Value = value
		l.queue.MoveToFront(elem)
		return nil
	}

	if l.size >= l.cap {
		item, ok := l.queue.Back().Value.(valueItem)
		if ok {
			delete(l.dict, item.key)
		}
		l.queue.Remove(l.queue.Back())
		l.size--
	}

	elem = l.queue.PushFront(valueItem{key: key, value: value, ts: time.Now()})
	l.dict[key] = elem
	l.size++

	return nil
}

// Get get item
func (l *Lru) Get(key interface{}) (value interface{}, err error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	elem, ok := l.dict[key]
	if !ok {
		return value, fmt.Errorf("not found")
	}

	item, ok := elem.Value.(valueItem)
	if !ok {
		return value, fmt.Errorf("inner error")
	}

	// expired
	if time.Since(item.ts) > l.duration {
		l.ch <- key
		return value, nil
	}

	value = item.value

	return
}

// Erase erase item
func (l *Lru) Erase(key interface{}) (err error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	elem, ok := l.dict[key]
	if !ok {
		return nil
	}

	delete(l.dict, key)
	l.queue.Remove(elem)
	l.size--

	return nil
}

// Len len of item
func (l *Lru) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.size
}

// Range visit each item
func (l *Lru) Range(f func(key, value interface{})) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	for k, v := range l.dict {
		item, ok := v.Value.(valueItem)
		if !ok {
			continue
		}

		if time.Since(item.ts) > l.duration {
			l.ch <- k
			continue
		}

		f(k, item.value)
	}
}

func (l *Lru) doErase() {
	for {
		select {
		case key := <-l.ch:
			l.Erase(key)
		}
	}
}
