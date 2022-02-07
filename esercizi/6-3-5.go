package main

import (
	. "fmt"
  "os"
  "strconv"
)

func main() {
  lista := leggiNumero()
  Println(Minimo(lista))
  Println(Massimo(lista))
  Println(Media(lista))

}

func leggiNumero() (numeri []int){
  for _,v := range os.Args{
    if num, err := strconv.Atoi(v); err == nil{
      numeri = append(numeri, num)
    }
  }
  return numeri
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