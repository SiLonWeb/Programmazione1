package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	input := LeggiUtenze()

	lista := InizializzaRegistro()

	for _, v := range input {
		lista = AggiungiUtenza(lista, v)
	}

  for num := range lista {
    codicesim := Identifica(lista, num)
		if strings.HasPrefix(num, "338") &&  codicesim != ""{
      Printf("Il numero %s è associato alla sim %s\n", num, Identifica(lista, num))
    }
	}

	/* for {
			op := ""
			Scan(&op)

			switch op {
			case "1":
				input := LeggiUtenze()


				for _, v := range input {
					lista = AggiungiUtenza(lista, v)
				}

				Println(lista)
			case "2":
	      num := ""
	      Scan(&num)

	      Println(Identifica(lista, num))
			}
		} */
}

// FUNZIONI

func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)

	input := ""
	for scanner.Scan() {
		if scanner.Text() != "" {
			input += scanner.Text() + "\n"
		}
	}

	righe := strings.Split(input, "\n")

	for _, riga := range righe {
		if len(riga) > 0 {
			dati := strings.Split(riga, ";")
			utenze = append(utenze, Utenza{dati[0], dati[1]})
		}
	}

	return
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(map[string][]string)
	return
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) RegistroTelefonico {
		registro[utenza.numero_tel] = append(registro[utenza.numero_tel], utenza.codice_sim)
	return registro
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	if registro[telefono] != nil {
		if len(registro[telefono]) > 1 {
			codiceSIM = registro[telefono][len(registro[telefono])-1]
		} else {
			codiceSIM = registro[telefono][0]
		}
	} else {
		codiceSIM = ""
	}
	return codiceSIM
}

// TIPI

type Utenza struct {
	numero_tel, codice_sim string
}

type RegistroTelefonico map[string][]string
