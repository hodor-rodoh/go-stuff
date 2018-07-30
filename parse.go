package main

 // kubectl --context am562-kube0 get ing -n nozomi -o json | jq -r '.items[] | { app: .metadata.name, namespace: .metadata.namespace }' > output.txt
 // kubectl --context am562-kube0 get ingress -o=custom-columns=NAME:.metadata.name,NAMESPACE:.metadata.namespace -n nozomi | grep telematicsct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type output struct {
//   output []metadata
// }

type metadata struct {
	App       string //`json:"app"`
	Namespace string //`json:"namespace"`
}

func main() {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("opened dat file homes")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var info []metadata

	json.Unmarshal(([]byte(byteValue)), &info)
	// data should be printed but I don't think its Unmarshalling properly
	fmt.Println(info) // output is emtpy []

	// for i := 0; i < len(metadata.output); i++ {
	// 	fmt.Println(metadata.output[i].Namespace + " " + metadata.output[i].App)
	// }

}
