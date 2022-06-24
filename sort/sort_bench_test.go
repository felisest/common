package bsort

import (
	"fmt"

	"sort"
	"sync"
	"testing"
)

var sliceSize uint32 = 10000
var templateSlice []int32 = make([]int32, sliceSize)

var once sync.Once

func getTemplateSlice(size uint32) []int32 {

	once.Do(func() {
		templateSlice = GetSlice(size, int32(0))
	})

	var s_ret []int32 = make([]int32, size)
	copy(s_ret, templateSlice)

	return s_ret
}

func printSlice(name string, s []uint32) {
	fmt.Println(name)
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

var result []int32

func BenchmarkGoSort(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		sort.SliceStable(s, func(p, q int) bool {
			return s[p] < s[q]
		})
	}
	result = s
}

func BenchmarkInsertionSort(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		InsertionSort(s)
	}
	result = s
}

func BenchmarkMergeSort(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	var ss []int32

	for i := 0; i < b.N; i++ {

		ss = MergeSort(s)
	}
	result = ss
}

func BenchmarkMergeSortWiki(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	var ss []int32

	for i := 0; i < b.N; i++ {

		ss = MergeSortWiki(s)
	}
	result = ss
}

func BenchmarkBubbleSort(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		BubbleSort(s)
	}
	result = s
}

func BenchmarkQuickSort(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		QuickSort(s)
	}
	result = s
}

func BenchmarkBucketSort100(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		BucketSort(s, 100)
	}
	result = s
}

func BenchmarkBucketSort1000(b *testing.B) {

	b.StopTimer()
	s := getTemplateSlice(sliceSize)
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		BucketSort(s, 1000)
	}
	result = s
}
