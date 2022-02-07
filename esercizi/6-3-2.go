package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main() {
	var posizione int
	var errore bool
	valori := make([]int, len(os.Args)-1)
	for i := 0; i < len(os.Args)-1; i++ {

		num, err := strconv.Atoi(os.Args[i+1])
		if err == nil {
			valori[i] = num
		} else {
			posizione = i
			errore = true
			break
		}
	}

	for i := 1; i < len(valori); i++ {
		var val int = valori[i]
		if i%2 == 0 {
			if val != 0 && val <= valori[i-1] {
				posizione = i
				errore = true
				break
			}
		} else {
			if val != len(valori) && val >= valori[i+1] {
				posizione = i
				errore = true
				break
			}
		}
	}
	if errore {
		Printf("Posizione %d non valida", posizione)
    Println()
	} else {
		Println("Sequenza valida")
	}

}
