package main

import (
	. "fmt"
	"math/rand"
)
func main() {
	s:=[]float64{0.2, 1.2, 2, 3, 4, 5}
	rand.Shuffle(len(s), func (i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	Println(s)
}