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

// kubectl --context am360-kube0 get ing --all-namespaces -o json | jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
// kubectl --context am360-kube0 get secrets --all-namespaces -o json | jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json

type metadata struct {
	App       string
	Namespace string
}

func main() {
	jsonFile, err := os.Open("output.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var info []metadata
	var results []metadata

	json.Unmarshal(([]byte(byteValue)), &info)

	for _, app := range info {
		match, _ := regexp.MatchString("telematicsct", app.App)
		if match {
			results = append(results, app)
			// hodor := fmt.Sprintf("kubectl --context am360-kube0 get ingress %s --namespace %s", app.App, app.Namespace)
			// hodor := fmt.Sprintf("kubectl --context am360-kube0 --namespace %s delete ingress/%s \n", app.Namespace, app.App)
			// hodor := fmt.Sprintf("kubectl --context am360-kube0 get secrets %s --namespace %s", app.App, app.Namespace)
			hodor := fmt.Sprintf("kubectl --context am360-kube0 --namespace %s delete secret/%s \n", app.Namespace, app.App)
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

