package main

import (
	."fmt"
)

func main() {
	sl := []int{1,2,3,2,3,5,6,2,3,2}
	Println(Funzione(sl))
}

func Funzione(sl []int) (lunghezza int){
	count := 0
	max := 0
	for i := 1; i < len(sl); i++ {
		if sl[i-1] < sl[i] {
			count++
		} else {
			max = count
			count = 0
		}
	}
	return max+1
	
}