package main

import (
  "encoding/json"
  "io/ioutil"
  "regexp"
  "os"
  "os/exec"
  "log"
  "fmt"
  "strings"
)

type metadata struct {
	App       string
	Namespace string
}

var info []metadata
var results []metadata

func openFile() {
  jsonFile, err := os.Open("output.json")
  if err != nil {
    fmt.Println(err)
  }

  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(([]byte(byteValue)), &info)
}

func searchData() {

  for _, app := range info {
    match, _ := regexp.MatchString("telematicsct", app.App)
    if match {
      results = append(results, app)
      hodor := fmt.Sprintf("kubectl --context am160-kube0 get ingress %s --namespace %s", app.App, app.Namespace)
      // hodor := fmt.Sprintf("kubectl --context am160-kube0 --namespace %s delete ingress/%s \n", app.Namespace, app.App)
      // hodor := fmt.Sprintf("kubectl --context am160-kube0 get secrets %s --namespace %s", app.App, app.Namespace)
      // hodor := fmt.Sprintf("kubectl --context am160-kube0 --namespace %s delete secret/%s \n", app.Namespace, app.App)
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

func main() {
  openFile()
  searchData()
}
