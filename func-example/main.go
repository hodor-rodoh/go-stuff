package main

import (
  // "encoding/json"
  // "io/ioutil"
  // "regexp"
  "os"
  "os/exec"
  "log"
  // "fmt"
  // "strings"
)

// func getData() {
//   hodor := fmt.Sprintf("kubectl --context am560-kube0 get ingress --all-namespaces -o json")
//   // rodoh := fmt.Sprintf(" | jq -r '[.items[] | { app: .metadata.app, namespace: .metadata.namespace }]'")
//   // butts := hodor + rodoh
//   fmt.Println(hodor)
//   tokens := strings.Fields(hodor)
//   executable := tokens[0]
//   args := tokens[1:len(tokens)]
//
//   cmd := exec.Command(executable, args...)
//   cmd.Stdout = os.Stdout
//   cmd.Stderr = os.Stderr
//   err := cmd.Run()
//   if err != nil {
//     log.Fatal(err)
//   }
// }

func helloTest() {
  cmd := exec.Command("sh", "-c", "kubectl --context am560-kube0 get ingress -n nozomi -o json")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  // getData()
  helloTest()
}
