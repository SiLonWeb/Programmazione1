package main

import (
	"bufio"
	. "fmt"
	"os"
  "strings"
)

func main() {
  Println(TraduciTestoInFarfallino(LeggiTesto()))
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func TraduciCarattereInFarfallino(c rune) (letteratradotta string){
  maiusc := "AEIOU"
  minusc := "aeiou"
  if strings.ContainsRune(maiusc, c) {
    letteratradotta = string(c)+"F"+string(c)
  } else if strings.ContainsRune(minusc, c) {
    letteratradotta = string(c)+"f"+string(c)
  } else {
    letteratradotta = string(c)
  }
  return
}

func TraduciTestoInFarfallino(s string) (testotradotto string){
  for _,v := range s{
    testotradotto += TraduciCarattereInFarfallino(v)
  }
  return
}