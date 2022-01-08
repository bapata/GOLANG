package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Node struct {
	Name          string    `json:"name"`
	CpuCount      int	`json:"cpu_count"`
	RamSizeGB     int	`json:"ram_size_gb"`
       CpuNames      []string  `json:"cpu_names"`
}

func main() {

	u := &Node{
		Name:      "myhost1.example.com",
		CpuCount:  4,
                RamSizeGB: 200,
		CpuNames: []string{"cpu0","cpu1","cpu2","cpu3"},
	}
        v := &Node{
		Name:      "myhost2.example.com",
		CpuCount:  2,
                RamSizeGB: 100,
		CpuNames: []string{"cpu0","cpu1"},
	}
	w := []*Node{u,v}
	out, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
