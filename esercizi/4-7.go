package main

import (
	"fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
	for i:=1; i<=n; i++ {
    for j:=0; j<n; j++ {
      if i%2 == 1 {
        fmt.Print("*")
      }else{
        fmt.Print("+")
      }
    }
    fmt.Println()
  }
}