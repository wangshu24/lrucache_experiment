package list

import (
	"errors"
	"time"
)

type Entry[K comparable, V any] struct {
	Next, Back *Entry[K, V]
	Key        K
	Value      V
	Bday       time.Time
	TTL        time.Duration
}

type List[K comparable, V any] struct {
	Tail, Root Entry[K, V]
	len, cap   int
}

type V interface{}

func NewList[K comparable, V any](cap int) *List[K, V] {
	l := new(List[K, V])
	var c int
	if cap > 0 {
		c = cap
	}
	l.cap = c
	l.Root.Next = &l.Root
	l.Tail.Back = &l.Root
	return l
}

// Add add a new item to the list and return boolean for whether an old item was discarded or not
// TODO: added condition to check for existing key value pair, if exist, move to top of list and remove old one
func (l *List[K, V]) Add(e Entry[K, V]) bool {
	if l.len == 0 {
		l.Root = e
		l.Tail = e
		l.len++
		return false
	}

	l.Tail.Next = &e
	l.Tail = e
	l.len++
	evicted := l.len > l.cap
	if l.len > l.cap {
		l.Root = *l.Root.Next
		l.len--
	}
	return evicted
}

func (l *List[K, V]) GetInd(index int) (*Entry[K, V], error) {
	if index > l.len || index > l.cap {
		return nil, errors.New("invalid index")
	}

	tmp := l.Root
	for i := 0; i < index; i++ {
		tmp = *tmp.Next
	}
	time := time.Now()
	endtime := tmp.Bday.Add(tmp.TTL)
	if time.After(endtime) {
		return nil, errors.New("stale cache")
	}
	tmp.Bday = time
	return &tmp, nil
}

func (l *List[K, V]) RemoveInd(index int) error {
	if index > l.len || index > l.cap {
		return errors.New("invalid index")
	}
	tmp := l.Root
	for i := 0; i < index; i++ {
		tmp = *tmp.Next
	}

	back := tmp.Back
	next := tmp.Next
	back.Next = next
	return nil
}

func (l *List[K, V]) PeekInd(index int) (*Entry[K, V], error) {
	if index > l.cap || index > l.len {
		return nil, errors.New("invalid index")
	}

	tmp := l.Root
	for i := 0; i < index; i++ {
		tmp = *tmp.Next
	}

	return &tmp, nil
}

func (l *List[K, V]) Len() int {
	return l.len
}

func (l *List[K, V]) Cap() int {
	return l.cap
}
