package main

import (
	"encoding/json"
	"fmt"
)
// Purpose of this code is to store and retrieve students marks
// linear search is used

// data structure
type Profile struct {
	Student_name string `json:"name"`
	Marks        int    `json:"marks"`
}
type Profiles struct {
	ProfileList []Profile
}

// Store marks for a given student
func (p *Profiles) set_marks(student string, marks int) int {
	entry := Profile{Student_name: student, Marks: marks}
	// find free slot
	for _, v := range p.ProfileList {
		if v.Student_name == student {
			break
		}
	}
	// We are at the end
	p.ProfileList = append(p.ProfileList, entry)
	return 0
}

// linear search
func (p *Profiles) get_marks(student string) int {
	for _, v := range p.ProfileList {
		if v.Student_name == student {
			return v.Marks
		}
	}
	return -1
}

func main() {
	myp := Profiles{}

        // build marks
	myp.set_marks("student1", 98)
	myp.set_marks("student2", 97)
	myp.set_marks("student3", 47)

        // report
	b, err := json.MarshalIndent(myp, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
