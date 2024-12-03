package utils

import (
	"cmp"
	"fmt"
)

type Item struct {
	value    string
	priority int
	index    int
}

func (it *Item) Compare(it1 *Item) int {
	return -cmp.Compare(it.priority, it1.priority)
}

func (it *Item) setIndex(index int) {
	it.index = index
}

func ExampleHeap() {
	pq := NewHeap[*Item]((*Item).Compare)
	pq.SetIndex((*Item).setIndex)

	pq.Push(&Item{value: "banana", priority: 3})
	pq.Push(&Item{value: "apple", priority: 2})
	pq.Push(&Item{value: "pear", priority: 4})

	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("%d:%s ", item.priority, item.value)
	}
	fmt.Println()

	// Output: 4:pear 3:banana 2:apple
}

func ExampleHeap2() {
	pq := NewHeap[*Item]((*Item).Compare)
	pq.SetIndex((*Item).setIndex)

	pq.Push(&Item{value: "banana", priority: 3})
	pq.Push(&Item{value: "apple", priority: 2})
	pq.Push(&Item{value: "pear", priority: 4})

	orange := &Item{value: "orange", priority: 1}
	pq.Push(orange)
	orange.priority = 5
	pq.Fix(orange.index)

	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	fmt.Println()

	// Output: 05:orange 04:pear 03:banana 02:apple
}
