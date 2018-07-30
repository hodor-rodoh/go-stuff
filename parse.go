package main

 // kubectl --context am562-kube0 get ing -n nozomi -o json | jq -r '.items[] | { app: .metadata.name, namespace: .metadata.namespace }' > output.txt
 // kubectl --context am562-kube0 get ingress -o=custom-columns=NAME:.metadata.name,NAMESPACE:.metadata.namespace -n nozomi | grep telematicsct

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"regexp"
)
// go get go get k8s.io/client-go/*version#
// kubectl --context am562-kube0 get ing -n nozomi -o json

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

	fmt.Println("opened dat file homes!")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var info []metadata

	json.Unmarshal(([]byte(byteValue)), &info)
	// fmt.Println(info)

	for i := 0; i < len(info); i++ {
		// fmt.Println(info[i].App + " " + info[i].Namespace)
		// add a regexp.MatchString for .App (`telematicsct`,info)
		fmt.Println(info[i])
	}

}
