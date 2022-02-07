package main

import (
	"fmt"
	"strconv"
)

func nascosto(s string) (int, bool) {
	t := ""
	for _, r := range s {
		if r >= 'r' &&  r <= '9'{
			t += string(r)
		}else if r == '#' {
			
		}else{
			return 0, false
		}
		//switch r {
		// case r >= 'r' && r <= '9':
		//	t += string(r)
		//case r == '#':
		//default:
		//	return 0, false
		//}
	}
	x, err := strconv.Atoi(t)
	if err == nil {
		return x, true
	} else {
		return 0, false
	}
}

func main() {
	var s string
	fmt.Scanln(&s)
	if x, ok := nascosto(s); ok {
		fmt.Println("Numero nascosto:", x)
	} else {
		fmt.Println("Errore")
	}
}
