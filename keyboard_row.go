package main

import (
	"fmt"
	"strings"
)

func get_row_for_char(mychar string) int {
	kb_characters := make(map[int]string)
	kb_characters[0] = "qwertyuiop"
	kb_characters[1] = "asdfghjkl"
	kb_characters[2] = "zxcvbnm"

	for ii := 0; ii < 3; ii++ {
		if strings.Contains(kb_characters[ii], mychar) {
			return ii
		}
	}
	return -1
}

func is_in_same_kb_row(word string) bool {
	tmp_dict := make(map[string]int)

	for _, mychar := range word {
		tmp_dict[string(mychar)] = get_row_for_char(string(mychar))
	}
	row_list := []int{}
	for _, v := range tmp_dict {
		row_list = append(row_list, v)
	}
	desired_row_number := row_list[0]
	for ii := 1; ii < len(row_list); ii++ {
		if row_list[ii] != desired_row_number {
			return false
		} else {
			desired_row_number = row_list[ii]
		}
	}
	return true
}

func process_str_list(words []string) []string {
	out_list := []string{}

	for _, v := range words {
		if is_in_same_kb_row(v) {
			out_list = append(out_list, v)
		}
	}
	return out_list
}

func main() {
	words1 := []string{"two", "dad", "cat"}  // return ["two", "dad"]
	words2 := []string{"ufo", "xzy", "byte"} // return []

	fmt.Println(process_str_list(words1))
	fmt.Println(process_str_list(words2))
}
