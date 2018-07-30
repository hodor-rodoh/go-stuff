package main

import (
  "encoding/json"
  "fmt"

)

type Bird struct {
  App string
  Namespace string
}

func main() {
  BJson := `[
  {
    "app":"butts1",
    "namespace":"nozomi"
  },
  {
    "app":"butts2",
    "namespace":"nozomi"
  }]`
  var birds []Bird
  json.Unmarshal([]byte(BJson), &birds)
  // fmt.Printf("%+v", birds)
  fmt.Println(birds[0])
  fmt.Println(birds[1])
}
