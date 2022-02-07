package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	Println("Inserisci testo:")
	Println("Testo inverso:",invertiStringa(LeggiTesto()))
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			testo += scanner.Text() + "\n"
		} else {
			return
		}
	}
	return
}

func invertiStringa(s string) (inversa string){
  for i := len(s)-1; i >= 0; i--{
    inversa += string(s[i])
  }
  return
}
