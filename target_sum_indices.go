package main

// Given input_list=[1,4,6,5],target=10
// returns: [1,2] (because 4(index=1)+6(index=2) = 10 (target))
// Complexity: O(n^2)
// TBD: See if this can be done in <O(n^2)

import (
	"fmt"
)

type MyTargetLists struct {
	input_list        []int
	target            int
	target_index_list []int
}

func (tl *MyTargetLists) set_target(t int) {
	tl.target = t
}

func (tl *MyTargetLists) set_input_list(mylist []int) {
	tl.input_list = mylist
}

// algorithm
func (tl *MyTargetLists) get_index_list() []int {
	for ii := 0; ii < len(tl.input_list); ii++ {
		for jj := 0; jj < len(tl.input_list); jj++ {
			if tl.input_list[ii]+tl.input_list[jj] == tl.target {
				tl.target_index_list = append(tl.target_index_list, ii)
				tl.target_index_list = append(tl.target_index_list, jj)
				return tl.target_index_list
			}
		}
	}
	return tl.target_index_list
}

func main() {
	mtl := MyTargetLists{}
	mtl.set_target(3)
	mtl.set_input_list([]int{10, 20, 1, 2, 45, 7})
	fmt.Println(mtl.get_index_list())
}
