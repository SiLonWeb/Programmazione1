package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	StampaIstogramma(Occorrenze(LeggiTesto()))
}

func StampaIstogramma(occorrenze map[rune]int) {
	var valori []int
	for i, _ := range occorrenze {
		valori = append(valori, int(i))
	}
	sort.Ints(valori)
	for _, v := range valori {
		Print(string(rune(v)), ": ")
		asterischi := occorrenze[rune(v)]
		for i := 0; i < asterischi; i++ {
			Print("*")
		}
    Println()
	}
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return
}

func Occorrenze(s string) map[rune]int {
	occorrenze := map[rune]int{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			occorrenze[v]++
		}
	}
	return occorrenze
}
