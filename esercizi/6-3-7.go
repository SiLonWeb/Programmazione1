package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
  Println(Calcola(LeggiNumeri()))
}

func LeggiNumeri() []int{
  var listanumeri []int
  for _,v := range os.Args{
    numero, err := strconv.Atoi(v)
    if err == nil {
      listanumeri = append(listanumeri, numero)
    }
  }
  return listanumeri
}

func Calcola(sl []int) int{
  var somma int = 0
  for i,v := range sl {
    for j:=i+1; j<len(sl); j++{
      Println(v, "-" ,sl[j])
      prod := v * sl[j]
      if prod%2 == 0{
        somma += prod
      }
    }
  }
  return somma
}