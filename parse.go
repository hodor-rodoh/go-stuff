package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
  // "strconv"
)

type metadata struct {
	App       string `json:"app"`
	Namespace string `json:"namespace"`
}

type tag struct {
  tag []metadata
}

func main() {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var metadata tag

	json.Unmarshal(byteValue, &metadata)
  fmt.Println(metadata.tag)

	for i := 0; i < len(metadata.tag); i++ {
		fmt.Println(metadata.tag[i].Namespace + " " + metadata.tag[i].App)
    // fmt.Println(strconv.Itoa(metadata.tag[i].App))
	}

}
