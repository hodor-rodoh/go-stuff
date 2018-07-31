package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)
// go get k8s.io/client-go/*version#

type metadata struct { //struct for app and namespace for outputs
	App       string //`json:"app"`
	Namespace string //`json:"namespace"`
}

func main() {
	jsonFile, err := os.Open("output.json") 	// opens output.json file prints error with err
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close() 	//close file

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var info []metadata
	var results []metadata

	json.Unmarshal(([]byte(byteValue)), &info) // Unmarshal jsonFile to parse and shit

	for _, app := range info { // loop through all the stuff
		match, _ := regexp.MatchString("telematicsct", app.App)
		if match {
			results = append(results, app)
			fmt.Printf("kubectl --namespace %s delete ingress/%s \n", app.Namespace, app.App)
		}
	}
}
