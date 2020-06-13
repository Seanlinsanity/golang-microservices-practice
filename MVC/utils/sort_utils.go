package utils

import (
	"sort"
)

//BubbleSort algorithm
func BubbleSort(elements []int) {
	//[1, 9, 5, 2]
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

//BubbleSort2 algorithm
func BubbleSort2(elements []int) {
	//[1, 9, 5, 2]
	for j := 0; j < len(elements); j++ {
		for i := 0; i < len(elements)-1-j; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

//Sort custom algorithm
func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
		return
	}
	sort.Ints(elements)
}
