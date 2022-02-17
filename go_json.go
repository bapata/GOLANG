package main

// Notice that fields of struct are starting with upperCase letter

import (
	"encoding/json"
	"fmt"
)

type host_type struct {
	Host_fqdn       string            `json:"fqdn"`
	Host_attributes []string          `json:"attributes"`
	Host_extra      map[string]string `json:"host_extras"`
}

func main() {
	h1 := host_type{
		Host_fqdn:       "h1.example.com",
		Host_attributes: []string{"disabled", "up"},
		Host_extra:      map[string]string{"a": "1", "b": "2"},
	}

	fmt.Println(h1)

	json_raw1, e1 := json.MarshalIndent(h1, "", " ")
	if e1 == nil {
		fmt.Println(string(json_raw1))
	}
}
