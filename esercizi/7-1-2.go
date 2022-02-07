package main

import (
	. "fmt"
	"os"
	"strings"
)

func main() {
	s := Leggi()
	if isValid(s) {
		if isBalanced(s) {
			Println("Bilanciata")
		} else {
			Println("Non bilanciata")
		}
	} else {
		Println("Sequenza Errata")
	}
	Println("---")
	Println("Sottosequenze Bilanciate")
  n := 0
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if isBalanced(s[i:j]) {
        n++
				Print(n,") ", s[i:j], "\n")
			}
		}
	}
  if n == 0{
    Println("Nessuna.")
  }
  }

func Leggi() string {
	stringa := os.Args[1]
	strings.Trim(stringa, "\"")
	return stringa
}

func isValid(stringa string) bool {
	for _, r := range stringa {
		if r != '(' && r != ')' {
			return false
		}
	}
	return true
}

func isBalanced(sequence string) bool {
	var status int = 0
	for _, v := range sequence {
		if v == '(' {
			status++
		}
		if v == ')' {
			status--
		}
		if status < 0 {
			return false
		}
	}
	if status != 0 {
		return false
	}
	return true
}
