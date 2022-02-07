package main

import (
	. "fmt"
  "os"
  "strconv"
)

func main() {
  var valori []int
  for _,v := range os.Args{
    if numero, err := strconv.Atoi(v); err == nil{
      valori = append(valori, numero)
    }
  }
  Println(Minimo(valori))
  Println(Massimo(valori))
  Println(Media(valori))

}

func Minimo(sl []int) int {
  var min int
  for i,v := range sl {
    if i == 0 {
      min = v
    }
    if v < min {
      min = v
    }
  }
  return min
}
func Massimo(sl []int) int {
var max int
  for i,v := range sl {
    if i == 0 {
      max = v
    }
    if v > max {
      max = v
    }
  }
  return max
}
func Media(sl []int) float64 {
  var somma float64
  for _,v := range sl{
    somma += float64(v)
  }
  return somma/float64(len(sl))
}