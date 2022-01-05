package main

// brew install go
// go run monotonically_incr_decr.go

import "fmt"

func mono_incr(mylist []int) bool {
	sorted_ascending := true
	for index, _ := range mylist {
		if index > len(mylist)-2 {
			break
		}
		if mylist[index] > mylist[index+1] {
			sorted_ascending = false
		}
	}
	return sorted_ascending
}

func mono_decr(mylist []int) bool {
	sorted_descending := true
	for index, _ := range mylist {
		if index > len(mylist)-2 {
			break
		}
		if mylist[index] < mylist[index+1] {
			sorted_descending = false
		}
	}
	return sorted_descending
}

func mono_incr_or_decr(mylist []int) bool {
	return mono_incr(mylist) || mono_decr(mylist)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 4, 5} // true
	nums2 := []int{7, 6, 3}          // true
	nums3 := []int{8, 4, 6}          // false

	fmt.Println(mono_incr_or_decr(nums1))
	fmt.Println(mono_incr_or_decr(nums2))
	fmt.Println(mono_incr_or_decr(nums3))
}
