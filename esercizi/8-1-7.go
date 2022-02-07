package main

import (
	"bufio"
	. "fmt"
	"os"
	"unicode"
)

func main() {
	Stampa(ContaRipetizioni(SeparaParole(LeggiTesto())))
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func SeparaParole(s string) (parole []string) {
	var parola string
	for _, lettera := range s {
		if unicode.IsLetter(lettera) {
			parola += string(lettera)
		} else if parola != ""{
				parole = append(parole, parola)
				parola = ""
		}
	}
	return
}

func ContaRipetizioni(sl []string) map[string]int {
  ripetizioni := map[string]int{}
  for _,parola := range sl {
    ripetizioni[parola] += 1
  }
  return ripetizioni
}

func Stampa(rip map[string]int) {
  for parola, ripetizione := range rip {
    Printf("%s: %d\n", parola, ripetizione)
  }
}