package main

import(
  "fmt"
  "os"
  "strconv"
)

func main() {
  a,_ := strconv.Atoi(os.Args[1])
  b,_ := strconv.Atoi(os.Args[2])
  somma := 0
  for i := a+1; i < b; i++ {
    if i%2 == 1 {
      somma += i
    }
  }
  fmt.Println(somma)
}