package main

import (
	."fmt"
)

func main() {
	var n, pos int = 0, 0
	Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == pos{
        Print("o ")
      }else if j > pos{
        Print("+ ")
      }else{
        Print("* ")
      }
		}
    pos++
		Println()
	}
}