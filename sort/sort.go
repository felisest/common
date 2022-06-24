package bsort

import (
	"fmt"
	"math/rand"
	"unsafe"

	"golang.org/x/exp/constraints"

	"time"
)

func PrintSlice(name string, s []uint32) {
	fmt.Println(name)
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func GetSlice[T ~uint32 | ~uint64 | ~int32 | ~int64](size uint32, ofType T) []T {

	s := make([]T, size)
	rand.Seed(time.Now().UnixNano())

	for i := uint32(0); i < size; i++ {

		switch any(ofType).(type) {

		case uint32:
			s[i] = T(rand.Uint32())
		case uint64:
			s[i] = T(rand.Uint64())
		case int32:
			s[i] = T(rand.Int31())
		case int64:
			s[i] = T(rand.Int63())
		}
	}
	return s
}

func SliceEqual[T comparable](s1 []T, s2 []T) bool {

	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func Swap[T any](first *T, second *T) {
	*first, *second = *second, *first
}

func InsertionSort[T constraints.Ordered](s []T) {

	for j := 1; j < len(s); j++ {
		key := s[j]
		i := j - 1
		for ; i >= 0 && s[i] > key; i-- {
			s[i+1] = s[i]
		}
		s[i+1] = key
	}
}

func MergeSort[T constraints.Ordered](s []T) []T {

	s_len := len(s)

	if s_len > 2 {

		sl := MergeSort(s[:s_len/2])
		sr := MergeSort(s[s_len/2:])

		sm := make([]T, len(sl)+len(sr))

		for i, il, ir := 0, 0, 0; i < len(sm); i++ {

			if ir < len(sr) && il < len(sl) {

				if sl[il] < sr[ir] {

					sm[i] = sl[il]
					il++

				} else {

					sm[i] = sr[ir]
					ir++
				}

			} else if ir >= len(sr) {

				sm[i] = sl[il]
				il++

			} else if il >= len(sl) {

				sm[i] = sr[ir]
				ir++
			}
		}

		return sm

	} else if s_len > 1 && s[0] > s[1] {
		//Swap(&s[0], &s[1])
		s[0], s[1] = s[1], s[0]
	}

	return s
}

func MergeSortWiki[T constraints.Ordered](list []T) []T {
	count := len(list)

	switch {
	case count > 2:
		lb := MergeSortWiki(list[:count/2])
		rb := MergeSortWiki(list[count/2:])
		list = append(lb, rb...)
		lastIndex := len(list) - 1

		for i := 0; i < lastIndex; i++ {
			mv := list[i]
			mi := i

			for j := i + 1; j < lastIndex+1; j++ {
				if mv > list[j] {
					mv = list[j]
					mi = j
				}
			}

			if mi != i {
				list[i], list[mi] = list[mi], list[i]
			}
		}

	case len(list) > 1 && list[0] > list[1]:
		list[0], list[1] = list[1], list[0]
	}

	return list
}

func BubbleSort[T constraints.Ordered](s []T) {

	size := len(s)

	for j := size - 1; j > 0; j-- {

		for i := 0; i < j; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
}

func QuickSort[T constraints.Ordered](s []T) {

	size := len(s)

	if size <= 1 {
		return
	}

	pivot := s[(size/2)-1:][0]

	s[(size/2)-1], s[size-1] = s[size-1], s[(size/2)-1]

	lb := 0
	rb := size - 2

	for ; ; lb++ {

		if s[lb] >= pivot {

			for ; s[rb] >= pivot && lb != rb; rb-- {
			}
			s[lb], s[rb] = s[rb], s[lb]
		}

		if lb == rb {
			break
		}
	}

	if s[lb] >= s[size-1] {
		s[lb], s[size-1] = s[size-1], s[lb]
	}

	QuickSort(s[:lb+1])
	QuickSort(s[lb+1:])
}

func BucketSort[T interface {
	~int8 | ~int16 | ~int32 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint
}](s []T, num_buckets uint) error {

	if num_buckets > uint(len(s)) {
		return fmt.Errorf("too many backets")
	}

	var step uint64
	var max T
	var min T

	buckets := make([][]T, num_buckets)

	var t T
	switch any(t).(type) {

	case uint8, uint16, uint32, uint64, uint:
		max = ^T(0)
		step = uint64(max) / uint64(num_buckets)
		min = 0

	case int8, int16, int32, int64, int:
		max = 1<<(unsafe.Sizeof(t)*8-1) - 1

		step = uint64(max*2) / uint64(num_buckets)
		min = max
	}

	for i := 0; uint(i) < num_buckets; i++ {
		buckets[i] = make([]T, 0, 100)
	}

	var b_num uint64

	for _, v := range s {
		b_num = uint64((v + min)) / step
		buckets[b_num] = append(buckets[b_num], v)
	}

	i := 0
	for _, b := range buckets {
		InsertionSort(b)
		for _, v := range b {
			s[i] = v
			i++
		}
	}

	return nil
}
