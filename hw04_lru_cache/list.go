package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	elem := &ListItem{Value: v}

	if l.len == 0 {
		l.front = elem
		l.back = elem
	}

	if l.len == 1 {
		l.front = elem
		l.front.Next = l.back
		l.back.Prev = l.front
	}

	if l.len > 1 {
		elem.Next = l.front
		l.front.Prev = elem
		l.front = elem
	}

	l.len++

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	elem := &ListItem{Value: v}

	if l.len == 0 {
		l.back = elem
		l.front = elem
	}

	if l.len == 1 {
		l.back = elem
		l.back.Prev = l.front
		l.front.Next = l.back
	}

	if l.len > 1 {
		elem.Prev = l.back
		l.back.Next = elem
		l.back = elem
	}

	l.len++

	return l.back
}

func (l *list) Remove(i *ListItem) {
	if l.len == 1 {
		l.front = nil
		l.back = nil
	} else if l.len == 2 && i == l.Back() {
		l.back = l.front
		l.back.Prev = nil
		l.back.Next = nil
	} else if l.len == 2 && i == l.Front() {
		l.front = l.back
		l.front.Prev = nil
		l.front.Next = nil
	} else if i == l.Front() {
		l.front = l.front.Next
		l.front.Prev = nil
	} else if i == l.back {
		l.back = l.back.Prev
		l.back.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	}

	if i == l.back && l.len == 2 {
		temp := l.back
		l.back = l.front
		l.front = temp

		l.back.Next = nil
		l.back.Prev = l.front

		l.front.Next = l.back
		l.front.Prev = nil
		return
	}

	if i == l.back && l.len > 2 {
		l.back = l.back.Prev
		l.back.Next = nil

		l.front.Prev = i
		i.Next = l.front

		l.front = i
		l.front.Prev = nil
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	l.front.Prev = i
	i.Next = l.front
	i.Prev = nil
	l.front = i
}

func NewList() List {
	return &list{}
}
