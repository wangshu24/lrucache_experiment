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

//Add add a new item to the list and return boolean for whether an old item was discarded or not
//TODO: added condition to check for existing key value pair, if exist, move to top of list and remove old one 
func (l *List[K , V]) Add(e *Entry[K, V])  bool {
	
	l.tail.next =  e
	l.tail = e
	l.len++
	evicted := l.len > l.cap
	if l.len > l.cap {
		l.root = l.root.next
		l.len--
	}
 	return evicted
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
