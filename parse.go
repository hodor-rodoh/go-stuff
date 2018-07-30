package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type output struct {
  output []metadata
}

type metadata struct {
	App       string `json:"app"`
	Namespace string `json:"namespace"`
}

func main() {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("opened dat file homes")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var metadata output

	json.Unmarshal(byteValue, &metadata)
	// data should be printed but I don't think its Unmarshalling properly
	fmt.Println(metadata.output) // output is emtpy []

	for i := 0; i < len(metadata.output); i++ {
		fmt.Println(metadata.output[i].Namespace + " " + metadata.output[i].App)
	}

}
