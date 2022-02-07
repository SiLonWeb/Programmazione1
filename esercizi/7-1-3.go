package main

import (
	. "fmt"
	"os"
)

func main() {
	input := Leggi()
  for _,v := range input {
    if v < '0' || v > '9'{
      return
    } 
  }
  Println(Sottostringhe(input))
}

func Leggi() string {
	stringa := os.Args[1]
	return stringa
}

func Sottostringhe(s string) (ssq []string) {
  for i:=0; i<len(s); i++ {
    for j:=i+1; j<len(s); j++{
      if s[j] > s[j-1]{
        ssq = append(ssq, string(s[i:j+1]))
      }else{
        break
      }
    }
  }
  return ssq
}