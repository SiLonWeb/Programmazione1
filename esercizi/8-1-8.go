package main

import (
  . "fmt"
  "os"
)

func main() {
  Stampa(TrovaSequenze())
}

func TrovaSequenze() map[string]int {
  occorrenze := map[string]int{}
  var seq string
  for i :=1; i < len(os.Args); i++ {
    seq += os.Args[i]
    for j:=i+1; j < len(os.Args); j++ {
      seq += os.Args[j]
      if IsSequenzaValida(seq) {
        occorrenze[seq]++
      }
    }
    seq = ""
  }
  return occorrenze
}

func IsSequenzaValida(seq string) bool {
  if len(seq) >= 3 && seq[0] == seq[len(seq)-1]{
    return true
  }else{
    return false
  }
}

func Stampa(occ map[string]int){
  for indice, numero := range occ {
    sequenza := ""
    for i:=0; i<len(indice); i++ {
      sequenza += string(indice[i]) + " "
    }
    Printf("%s -> Occorrenze: %d\n", sequenza, numero)
  }
}