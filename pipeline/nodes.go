package pipeline

import "sort"

// ArraySource read array source and return data in channel
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// InMemSort get channel int and sort in memory.
// return the ordered channel int
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// Read in memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		// Sort
		sort.Ints(a)

		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
