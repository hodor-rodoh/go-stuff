package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Record struct {
	App        string `json:"app"`
	Namespace  string `json:"namespace"`
}

func main() {
	var allRecords []Record

	input := []byte(`[
		{
    	"app": "name1",
    	"namespace": "ns"
    },
  	{
    	"app": "name2",
    	"namespace": "ns"
  	},
  	{
    	"app": "infradocs-tsp-qa",
    	"namespace": "nozomi"
  	},
	]`)

	var tmpRecords []Record
	err := json.Unmarshal(input, &tmpRecords)
	if (err != nil) {
		log.Fatal(err)
	}

	allRecords = append(allRecords, tmpRecords...)

	fmt.Println(tmpRecords)
}
