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
      "app": "infradocs-qa-goromasamu",
      "namespace": "nozomi"
    },
    {
      "app": "infradocs-qa-telematicsct",
      "namespace": "nozomi"
    },
    {
      "app": "infradocs-tsp-qa",
      "namespace": "nozomi"
    },
    {
      "app": "pghttpbin-qa-goromasamu",
      "namespace": "nozomi"
    },
    {
      "app": "pghttpbin-qa-telematicsct",
      "namespace": "nozomi"
    },
    {
      "app": "pghttpbin-tsp-qa",
      "namespace": "nozomi"
    },
    {
      "app": "pipelinedemo-qa-goromasamu",
      "namespace": "nozomi"
    },
    {
      "app": "pipelinedemo-qa-telematicsct",
      "namespace": "nozomi"
    },
    {
      "app": "pipelinedemo-tsp-qa",
      "namespace": "nozomi"
    }
  ]`)

	var tmpRecords []Record
	err := json.Unmarshal(input, &tmpRecords)
	if (err != nil) {
		log.Fatal(err)
	}

	allRecords = append(allRecords, tmpRecords...)

	fmt.Println(tmpRecords)
}
