package bsort

/*
#cgo CFLAGS: -I./libs
#cgo LDFLAGS: -L./libs/release -lcsort_static -Wl,-rpath=./lib/release -lstdc++
//void bubble(void* in_buff, int size);
void bubble_uint(void* in_buff, size_t size, size_t type_size);
void bubble_int(void* in_buff, size_t size, size_t type_size);
*/
import "C"

import (
	"unsafe"
)

func BubbleSortUint[T ~uint8 | ~uint16 | ~uint32 | ~uint64](s []T) {
	C.bubble_uint(unsafe.Pointer(&s[0]), C.ulong(len(s)), C.ulong(unsafe.Sizeof(s[0])))
}

func BubbleSortInt[T ~int8 | ~int16 | ~int32 | ~int64](s []T) {
	C.bubble_int(unsafe.Pointer(&s[0]), C.ulong(len(s)), C.ulong(unsafe.Sizeof(s[0])))
}
