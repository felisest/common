package main

import (
	"fmt"

	bsort "felis.est/examples/bsort"
)

func main() {

	s := []uint32{3200395169, 2209190777, 3179077564, 397518597, 3038765280, 1362336113, 4161221610, 2444693757, 2806669049, 3538105223}

	for _, v := range s {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\n")

	bsort.BubbleSortUint(s)

	for _, v := range s {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("\n")
}
