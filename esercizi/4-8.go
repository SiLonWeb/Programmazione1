package main

import (
	"fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
	for i:=1; i<=n; i++ {
    for j:=1; j<n*2; j++ {
      if i%3 == 2 {
        if j%2 == 1{
        fmt.Print("*")
        }else{
          fmt.Print(" ")
        }
      }else if i%3 == 1{
        if j%2 == 1{
        fmt.Print("+")
        }else{
          fmt.Print(" ")
        }
      }else{
        if j%2 == 1{
        fmt.Print("o")
        }else{
          fmt.Print(" ")
        }
      }
    }
    fmt.Println()
  }
}