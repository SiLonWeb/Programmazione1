package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	for _, v := range LeggiTesto() {
		Println(Formatta(v))
	}

}

func LeggiTesto() (testo []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			testo = append(testo, scanner.Text())
		} else {
			return
		}
	}
	return
}

func Formatta(s string) (formattata string) {
  elementi := strings.Split(s, "/")
  anno := string(elementi[0])
  mese := string(elementi[1])
  giorno := string(elementi[2])
  if len(mese) < 2{
    mese = "0"+mese
  }
  if len(giorno) < 2{
    giorno = "0"+giorno
  }
	formattata = anno +"-"+mese+"-"+giorno
	return
}