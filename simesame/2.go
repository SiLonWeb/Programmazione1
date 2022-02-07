package main

import (
	. "fmt"
  "sort"
)

func main() {
	var stringa string
	Scan(&stringa)

	if !VerifyString(stringa) {
		return
	}

	StampaOrdinata(Sottostringhe(stringa))
}

func Sottostringhe(input string) map[string]int {
	lista := map[string]int{}
  
  for i,v := range input {
    stringa := string(v)

    for j:=i+1; j < len(input); j++{
      
      if input[j] > input[j-1] {
        stringa += string(input[j])
        lista[stringa]++
      } else {
        break
      }
    
    }
  
  }
  return lista
}

func StampaOrdinata(lista map[string]int) {
  ordinata := []string{}

  for indice,_ := range lista {
    ordinata = append(ordinata, indice)
  }
  
  sort.Strings(ordinata)

  for _,stringa := range ordinata {
    Println(stringa ,lista[stringa])
  }
}

func VerifyString(s string) bool {
	for _, v := range s {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}
