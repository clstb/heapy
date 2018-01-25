package main

import "fmt"

func main() {
	a := []int{10, 100, 1, 10, 100, 10} // your slice
	size := len(a)
	b := make([]int, size)
	copy(b, a)

	for i := size / 2; i >= 1; i-- {
		maxHeapify(&a, i, size)
	}
	fmt.Println("Max-Heap: ", a)

	for i := size / 2; i >= 0; i-- {
		minHeapify(&b, i, size)
	}
	fmt.Println("Min-Heap: ", b)
}

func maxHeapify(a *[]int, i int, size int) {
	l := 2 * i
	r := 2*i + 1
	var largest int
	if (l < size) && ((*a)[l] > (*a)[i]) {
		largest = l
	} else {
		largest = i
	}
	if (r < size) && ((*a)[r] > (*a)[largest]) {
		largest = r
	}
	if largest != i {
		tmp := (*a)[largest]
		(*a)[largest] = (*a)[i]
		(*a)[i] = tmp
		maxHeapify(a, largest, size)
	}
}

func minHeapify(a *[]int, i int, size int) {
	l := 2 * i
	r := 2*i + 1
	var smallest = i
	if (l < size) && ((*a)[l] < (*a)[i]) {
		smallest = l
	}
	if (r < size) && ((*a)[r] < (*a)[smallest]) {
		smallest = r
	}
	if smallest != i {
		tmp := (*a)[smallest]
		(*a)[smallest] = (*a)[i]
		(*a)[i] = tmp
		minHeapify(a, smallest, size)
	}
}
