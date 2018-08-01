package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"log"
	"regexp"
	"strings"
)
// go get k8s.io/client-go/*version#
// kubectl --context am562-kube0 get ing -n nozomi -o json | jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
// kubectl --context am562-kube0 get secrets -n nozomi -o json | jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json

type metadata struct { //struct for app and namespace for outputs
	App       string //`json:"app"`
	Namespace string //`json:"namespace"`
}

// todo break this shit out into different functions
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
			// hodor := fmt.Sprintf("kubectl --context am562-kube0 get ingress %s --namespace %s", app.App, app.Namespace) // to get ingress for testing
			// hodor := fmt.Sprintf("kubectl --context am562-kube0 --namespace %s delete ingress/%s \n", app.Namespace, app.App) // deleting ingress
			hodor := fmt.Sprintf("kubectl --context am562-kube0 get secrets %s --namespace %s", app.App, app.Namespace) // to get secrets for testing
			// hodor := fmt.Sprintf("kubectl --context am562-kube0 --namespace %s delete secret/%s \n", app.Namespace, app.App) // deleting ingress
			fmt.Println(hodor)
			tokens := strings.Fields(hodor)
			executable := tokens[0]
			args := tokens[1:len(tokens)]
			cmd := exec.Command(executable, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

