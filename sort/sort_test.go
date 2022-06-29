package bsort

import (
	"sort"
	"testing"
)

func TestIntSort(t *testing.T) {

	size := uint32(10000)

	template := getTemplateSlice(size)
	var template_sorted []int32
	template_sorted = make([]int32, size)

	copy(template_sorted, template)

	sort.SliceStable(template_sorted, func(p, q int) bool {
		return template_sorted[p] < template_sorted[q]
	})

	// Insertion sort
	sorted := make([]int32, size)
	copy(sorted, template)

	InsertionSort(sorted)

	if !SliceEqual(sorted, template_sorted) {
		t.Error("Insertion sort: slices not equal")
	}

	// Merge sort
	copy(sorted, template)

	res := MergeSort(sorted)

	if !SliceEqual(res, template_sorted) {
		t.Error("Merge sort: slices not equal")
	}

	// Merge sort wiki
	copy(sorted, template)

	res = MergeSortWiki(sorted)

	if !SliceEqual(res, template_sorted) {
		t.Error("Merge Wiki sort: slices not equal")
	}

	// Quick sort
	qsorted := make([]int32, size)
	copy(qsorted, template)

	QuickSort(qsorted)

	if !SliceEqual(qsorted, template_sorted) {
		t.Error("Quick sort: slices not equal")
	}

	// Bucket sort
	// Insertion sort
	bsorted := make([]int32, size)
	copy(bsorted, template)

	BucketSort(bsorted, 100)

	if !SliceEqual(bsorted, template_sorted) {
		t.Error("Bucket sort: slices not equal")
	}

	// Bubble sort
	sorted = make([]int32, size)
	copy(sorted, template)

	BubbleSort(sorted)

	if !SliceEqual(sorted, template_sorted) {
		t.Error("Bubble sort: slices not equal")
	}

	// Internal Bubble sort
	sorted = make([]int32, size)
	copy(sorted, template)

	BubbleSortInt(sorted)

	if !SliceEqual(sorted, template_sorted) {
		t.Error("Internal Bubble sort: slices not equal")
	}	

}
