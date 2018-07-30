package main

import (
  "fmt"
  "regexp"
)

func main() {
  butts := "infradocs-qa-telematicsct"
  matched, err := regexp.MatchString(`telematicsct`, butts)
  fmt.Println(matched, butts) // true
  fmt.Println(err)     // nil (regexp is valid)
}
