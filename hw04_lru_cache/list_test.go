package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("add element to the front of empty list", func(t *testing.T) {
		l := NewList()

		value := 10

		front := l.PushFront(value) // [10]

		require.Equal(t, 1, l.Len())

		require.Equal(t, value, l.Front().Value)
		require.Equal(t, value, l.Back().Value)

		require.Equal(t, front, l.Front())
		require.Nil(t, front.Next)
		require.Nil(t, front.Prev)
	})

	t.Run("add element to the front of list with single element", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20

		l.PushFront(value1)          // [10]
		front := l.PushFront(value2) // [20, 10]

		require.Equal(t, 2, l.Len())

		require.Equal(t, value2, l.Front().Value)
		require.Equal(t, value1, l.Back().Value)

		require.Equal(t, value1, l.Front().Next.Value)
		require.Equal(t, value2, l.Back().Prev.Value)

		require.Equal(t, front, l.Front())
	})

	t.Run("add element to the front of list with multiple elements", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		l.PushFront(value1)          // [10]
		l.PushFront(value2)          // [20, 10]
		front := l.PushFront(value3) // [30, 20, 10]

		require.Equal(t, 3, l.Len())

		require.Equal(t, value3, l.Front().Value)
		require.Equal(t, value2, l.Front().Next.Value)
		require.Equal(t, value1, l.Front().Next.Next.Value)

		require.Equal(t, value1, l.Back().Value)
		require.Equal(t, value2, l.Back().Prev.Value)
		require.Equal(t, value3, l.Back().Prev.Prev.Value)

		require.Equal(t, front, l.Front())
	})

	t.Run("add element to the back of empty list", func(t *testing.T) {
		l := NewList()

		value := 10

		back := l.PushBack(value) // [10]

		require.Equal(t, 1, l.Len())

		require.Equal(t, value, l.Front().Value)
		require.Equal(t, value, l.Back().Value)

		require.Equal(t, back, l.Back())
		require.Nil(t, back.Next)
		require.Nil(t, back.Prev)
	})

	t.Run("add element to the back of list with single element", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20

		l.PushBack(value1)         // [10]
		back := l.PushBack(value2) // [10, 20]

		require.Equal(t, 2, l.Len())

		require.Equal(t, value1, l.Front().Value)
		require.Equal(t, value2, l.Back().Value)

		require.Equal(t, value2, l.Front().Next.Value)
		require.Equal(t, value1, l.Back().Prev.Value)

		require.Equal(t, back, l.Back())
	})

	t.Run("add element to the back of list with multiple elements", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		l.PushBack(value1)         // [10]
		l.PushBack(value2)         // [10, 20]
		back := l.PushBack(value3) // [10, 20, 30]

		require.Equal(t, 3, l.Len())

		require.Equal(t, value3, l.Back().Value)
		require.Equal(t, value2, l.Back().Prev.Value)
		require.Equal(t, value1, l.Back().Prev.Prev.Value)

		require.Equal(t, value1, l.Front().Value)
		require.Equal(t, value2, l.Front().Next.Value)
		require.Equal(t, value3, l.Front().Next.Next.Value)

		require.Equal(t, back, l.Back())
	})

	t.Run("remove element from the list center", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		l.PushFront(value1)             // [10]
		toRemove := l.PushFront(value2) // [20, 10]
		l.PushFront(value3)             // [30, 20, 10]
		l.Remove(toRemove)              // [30, 10]

		require.Equal(t, 2, l.Len())
		require.Equal(t, value1, l.Front().Next.Value)
		require.Equal(t, value3, l.Back().Prev.Value)
	})

	t.Run("remove element from the back of the list", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		toRemove := l.PushFront(value1) // [10]
		l.PushFront(value2)             // [20, 10]
		l.PushFront(value3)             // [30, 20, 10]
		l.Remove(toRemove)              // [30, 20]

		require.Equal(t, 2, l.Len())
		require.Equal(t, value2, l.Back().Value)
		require.Nil(t, l.Back().Next)
	})

	t.Run("remove element from the front of the list", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		l.PushFront(value1)             // [10]
		l.PushFront(value2)             // [20, 10]
		toRemove := l.PushFront(value3) // [30, 20, 10]
		l.Remove(toRemove)              // [20, 10]

		require.Equal(t, 2, l.Len())
		require.Equal(t, value2, l.Front().Value)
		require.Nil(t, l.Front().Prev)
	})

	t.Run("remove next to last (front) element from the list", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20

		l.PushFront(value1)             // [10]
		toRemove := l.PushFront(value2) // [20, 10]
		l.Remove(toRemove)              // [10]

		require.Equal(t, 1, l.Len())
		require.Equal(t, value1, l.Front().Value)
		require.Equal(t, value1, l.Back().Value)
	})

	t.Run("remove next to last (back) element from the list", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20

		toRemove := l.PushFront(value1) // [10]
		l.PushFront(value2)             // [20, 10]
		l.Remove(toRemove)              // [20]

		require.Equal(t, 1, l.Len())
		require.Equal(t, value2, l.Front().Value)
		require.Equal(t, value2, l.Back().Value)
	})

	t.Run("remove the last element from the list", func(t *testing.T) {
		l := NewList()

		value := 10

		toRemove := l.PushFront(value) // [10]
		l.Remove(toRemove)             // []

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("move to front current front", func(t *testing.T) {
		l := NewList()

		value := 10

		toMove := l.PushFront(value) // [10]
		l.MoveToFront(toMove)

		require.Equal(t, toMove, l.Front())
	})

	t.Run("move to front the element from the list center", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		l.PushFront(value1)           // [10]
		toMove := l.PushFront(value2) // [20, 10]
		l.PushFront(value3)           // [30, 20, 10]
		l.MoveToFront(toMove)         // [20, 30, 10]

		require.Equal(t, toMove, l.Front())

		require.Equal(t, value1, l.Front().Next.Next.Value)
		require.Equal(t, value2, l.Front().Value)
		require.Equal(t, value3, l.Front().Next.Value)

		require.Equal(t, value1, l.Back().Value)
		require.Equal(t, value2, l.Back().Prev.Prev.Value)
		require.Equal(t, value3, l.Back().Prev.Value)
	})

	t.Run("move to front current back element in list with two elements (swap)", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20

		toMove := l.PushFront(value1) // [10]
		l.PushFront(value2)           // [20, 10]
		l.MoveToFront(toMove)         // [10, 20]

		require.Equal(t, toMove, l.Front())

		require.Equal(t, value1, l.Front().Value)
		require.Equal(t, value2, l.Front().Next.Value)

		require.Equal(t, value1, l.Back().Prev.Value)
		require.Equal(t, value2, l.Back().Value)
	})

	t.Run("move to front current back element in list with multiple elements", func(t *testing.T) {
		l := NewList()

		value1 := 10
		value2 := 20
		value3 := 30

		toMove := l.PushFront(value1) // [10]
		l.PushFront(value2)           // [20, 10]
		l.PushFront(value3)           // [30, 20, 10]
		l.MoveToFront(toMove)         // [10, 30, 20]

		require.Equal(t, toMove, l.Front())

		require.Equal(t, value1, l.Front().Value)
		require.Equal(t, value2, l.Front().Next.Next.Value)
		require.Equal(t, value3, l.Front().Next.Value)

		require.Equal(t, value1, l.Back().Prev.Prev.Value)
		require.Equal(t, value2, l.Back().Value)
		require.Equal(t, value3, l.Back().Prev.Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
