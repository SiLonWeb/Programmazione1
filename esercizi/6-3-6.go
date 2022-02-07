package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
  numero, err := strconv.Atoi(os.Args[1])
  if err==nil {
    Println(Fattoriali(numero))
  }
}

func Fattoriali(n int) (f []int){
  f = make([]int, n)
  for i:=1; i<=n; i++{
    fact := 1
    for j:=1; j<=i; j++{
      fact *= j
    }
    f[i-1] = fact
  }
  return f
}