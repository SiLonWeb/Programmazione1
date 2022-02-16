package main

import (
	."fmt"
)


func Funzione(f func(string) string, s string) string {
	risultato := f(s) + f(s)
	return risultato 
}

func main() {
	f := func(str string) string{
		str += "A"
		return str
	}
	Println(Funzione(f, "Buongiorno"))
}