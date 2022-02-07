package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

func main() {
  Println(Spazia(LeggiTesto()))
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func Spazia(s string) (spaziata string){
  for _,v := range s {
    spaziata += string(v)
    if !unicode.IsSpace(v){
      spaziata += " "
    }
  }
  return
}