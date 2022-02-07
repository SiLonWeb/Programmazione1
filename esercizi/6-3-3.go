package main

import (
	. "fmt"
  "os"
  "strconv"
)

func main() {
  nums := leggiStringa()
  Println(moltiplica(nums))
}

func moltiplica(lista []int) int{
  var risultato int = 1
  for _,v := range lista{
    risultato = risultato * v
  }
  return risultato
}

func leggiStringa() []int{
  var stringa []int
  for _,v := range os.Args {
    numero, err := strconv.Atoi(v)

    if err == nil{
      stringa = append(stringa, numero)
    }
  }
  return stringa
}
