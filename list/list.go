package list

import (
	"errors"
	"time"
)

type Entry[K comparable, V any] struct {
	next, back *Entry[K, V]
	key K
	value V
	ttl    time.Time
}

type List[K comparable, V any] struct {
	tail, root *Entry[K, V]
	len, cap      int
	
}

type V interface {}

func (l *List[K , V]) Add(e *Entry[K, V])  bool {
	
	l.tail.next =  e
	l.len++
	
 	return l.len > l.cap
}

func (l *List[K, V]) Get(index int) ( *Entry[K,V] , error) {
	if index > l.len || index > l.cap { 
		return nil, errors.New("invalid index")
	}	
	
	tmp := l.root
	for i:=0; i < index;i++{
		tmp = tmp.next
	}
	return tmp, nil
}

func (l *List[K, V]) Remove(index int) error {
	if index > l.len || index > l.cap {
		return errors.New("invalid index")
	}
	tmp := l.root 
	for i:=0; i < index; i++ {
		tmp = tmp.next
	}

	back := tmp.back
	next := tmp.next
	back.next = next
	return nil
}

func (l *List[K, V]) Len() int {
	return l.len
}

func (l *List[K, V]) Cap() int {
	return l.cap
}
