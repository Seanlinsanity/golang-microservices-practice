package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// func TestBubbleSortWorstCase(t *testing.T) {
// 	elements := []int{7, 8, 9, 2, 5, 1}

// 	BubbleSort2(elements)
// 	assert.NotNil(t, elements)
// 	assert.EqualValues(t, 6, len(elements))
// 	assert.EqualValues(t, 1, elements[0])
// 	assert.EqualValues(t, 2, elements[1])
// 	assert.EqualValues(t, 5, elements[2])
// 	assert.EqualValues(t, 7, elements[3])
// 	assert.EqualValues(t, 8, elements[4])
// 	assert.EqualValues(t, 9, elements[5])
// }

// func TestBubbleSortBestCase(t *testing.T) {
// 	elements := []int{1, 2, 5, 7, 8, 9}

// 	BubbleSort2(elements)
// 	assert.NotNil(t, elements)
// 	assert.EqualValues(t, 6, len(elements))
// 	assert.EqualValues(t, 1, elements[0])
// 	assert.EqualValues(t, 2, elements[1])
// 	assert.EqualValues(t, 5, elements[2])
// 	assert.EqualValues(t, 7, elements[3])
// 	assert.EqualValues(t, 8, elements[4])
// 	assert.EqualValues(t, 9, elements[5])
// }

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElement(t *testing.T) {
	elements := getElements(5)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 4, elements[0])
	assert.EqualValues(t, 3, elements[1])
	assert.EqualValues(t, 2, elements[2])
	assert.EqualValues(t, 1, elements[3])
	assert.EqualValues(t, 0, elements[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort2_10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort2(elements)
	}
}

func BenchmarkSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort2_1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort2(elements)
	}
}

func BenchmarkSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

func BenchmarkBubbleSort200000(b *testing.B) {
	elements := getElements(200000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort2_200000(b *testing.B) {
	elements := getElements(200000)
	for i := 0; i < b.N; i++ {
		BubbleSort2(elements)
	}
}

func BenchmarkSort200000(b *testing.B) {
	elements := getElements(200000)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}
