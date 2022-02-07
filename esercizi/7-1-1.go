package main

import (
	. "fmt"
	"os"
)

func main() {
	Println(Sottostringhe(LeggiStringa()))
}

func LeggiStringa() string {
	stringa := os.Args[1]
	return stringa
}

func Sottostringhe(stringa string) (output []string) {
	for i := range stringa {
		for j := i + 2; j <= len(stringa); j++ {
			sub := stringa[i:j]
			if isPalindroma(sub) {
				output = append(output, sub)
			}
		}
	}
	return output
}

func isPalindroma(parola string) bool {
  var reverse []string
  for i:=len(parola)-1; i>=0; i--{
    reverse = append(reverse, string(parola[i]))
  }
  for i,v := range reverse {
    if v != string(parola[i]){
      return false
    }
  }
  return true
}
