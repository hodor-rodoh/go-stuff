package main

 // kubectl --context am562-kube0 get ing -n nozomi -o json | jq -r '.items[] | { app: .metadata.name, namespace: .metadata.namespace }' > output.txt
 // kubectl --context am562-kube0 get ingress -o=custom-columns=NAME:.metadata.name,NAMESPACE:.metadata.namespace -n nozomi | grep telematicsct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)
// go get k8s.io/client-go/*version#
// kubectl --context am562-kube0 get ing -n nozomi -o json

type metadata struct { //struct for app and namespace for outputs
	App       string //`json:"app"`
	Namespace string //`json:"namespace"`
}

type tinder struct { //struct for my regexp matches
	App 			string
	Namespace string
}

func main() {
	jsonFile, err := os.Open("output.json") 	// opens output.json file prints error with err
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("opened dat file homes!") 	// confirmation that file is opened

	defer jsonFile.Close() 	//close file

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var info []metadata

	json.Unmarshal(([]byte(byteValue)), &info) // Unmarshal jsonFile to parse and shit
	// fmt.Println(info)

	for i := 0; i < len(info); i++ { // loop through all the stuff
		// fmt.Println(info[i].App + " " + info[i].Namespace)
		// fmt.Println(info[i])

		match := (info[i].App + " " + info[i].Namespace)
		// fmt.Println(match)

		re := regexp.MustCompile("-telematicsct") // must contain -telematicsct
		fmt.Printf("%q\n", re.FindStringSubmatch(match)) // finds telematicsct in a substring and returns

	}
}
