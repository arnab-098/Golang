package main

import "fmt"

type Heap struct {
	size, length int
	values       []int
}

func createHeap(size int) Heap {
	h := Heap{size: size, length: 0, values: make([]int, size)}
	return h
}

func displayHeap(h Heap) {
	for i := 0; i < h.length; i++ {
		fmt.Print(h.values[i], "\t")
	}
	fmt.Println()
}

func isFullHeap(h Heap) bool {
	return h.length == h.size
}

func heapify(pos int, h *Heap) {
	largest := pos
	left := 2*pos + 1
	right := 2*pos + 2

	if left < h.size && h.values[left] > h.values[largest] {
		largest = left
	}
	if right < h.size && h.values[right] > h.values[largest] {
		largest = right
	}

	if largest != pos {
		temp := h.values[largest]
		h.values[largest] = h.values[pos]
		h.values[pos] = temp

		heapify(largest, h)
	}
}

func insertInHeap(num int, h *Heap) error {
	if isFullHeap(*h) {
		return fmt.Errorf("Heap is full")
	}
	h.values[h.length] = num
	h.length++
	for pos := h.length - 1; h.values[(pos-1)/2] < h.values[pos]; pos = (pos - 1) / 2 {
		temp := h.values[pos]
		h.values[pos] = h.values[(pos-1)/2]
		h.values[(pos-1)/2] = temp
	}
	return nil
}
