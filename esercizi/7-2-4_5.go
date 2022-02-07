package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	linee, parole, lettere := StatisticheParole(LeggiTesto())
	Println("Linee ", linee)
	Println("Parole ", parole)
	Println("Media ", float64(lettere)/float64(parole))
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func StatisticheParole(s string) (nline int, nword int, wordlen int) {
	for _, linea := range strings.Split(s, "\n") {
		if len(linea) > 0 {
			nline++
		}
		for _, parola := range strings.Split(linea, " ") {
			if len(parola) > 0 {
				nword++
				for _, lettera := range parola {
					if unicode.IsLetter(lettera) {
						wordlen++
					}
				}
			}
		}
	}
	return
}