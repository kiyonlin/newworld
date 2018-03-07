package queue

import (
	"fmt"
)

// Queue use []int to implement
type Queue []int

// Push an item to the given q
func (q *Queue) Push(v int) {
	fmt.Println(q)

	*q = append(*q, v)
}

// Pop a value from the given q
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty judge if the given q is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
