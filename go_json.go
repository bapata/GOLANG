package main

// testing json package, why we don't get json output

import (
	"encoding/json"
	"fmt"
)

type host_type struct {
	host_fqdn       string   `json:"fqdn"`
	host_attributes []string `json:"attributes"`
}

type host_type2 struct {
	host_fqdn string `json:"fqdn"`
	ha1       string `json:"attribute1"`
	ha2       string `json:"attribute2"`
}

func main() {
	h1 := host_type{host_fqdn: "h1.example.com", host_attributes: []string{"disabled", "up"}}
	h2 := host_type2{host_fqdn: "h1.example.com", ha1: "enabled", ha2: "down"}

	fmt.Println(h1)

	json_raw1, e1 := json.Marshal(h1)
	if e1 == nil {
		fmt.Println(string(json_raw1))
	}

	fmt.Println(h2)
	json_raw2, e2 := json.Marshal(h2)
	if e2 == nil {
		fmt.Println(string(json_raw2))
	}
}
