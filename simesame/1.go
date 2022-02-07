package main

import (
  . "fmt"
  "os"
  "strings"
  "strconv"
  "math"
)

func main() {
  eq := os.Args[1]
  a_divisa := strings.Split(eq, "x^2")
  a, _ := strconv.ParseFloat(a_divisa[0], 64)
  b_divisa := strings.Split(a_divisa[1], "x")
  b,_ := strconv.ParseFloat(b_divisa[0], 64)
  c_divisa := strings.Split(b_divisa[1], "=0")
  c,_ := strconv.ParseFloat(c_divisa[0], 64)
  var delta float64 = (b*b) - (4*a*c)
  if delta > 0 {
    sol1 := (-b - math.Sqrt(delta))/(2*a)
    sol2 := (-b + math.Sqrt(delta))/(2*a)
    Printf("Ha due soluzioni reali: %.3f e %.3f \n", sol1, sol2)

    if Soglia(sol1) {
      Printf("La soluzione %.3f è maggiore della soglia\n", sol1)
    }
    if Soglia(sol2) {
      Printf("La soluzione %.3f è maggiore della soglia\n", sol2)
    }
  } else if delta == 0 {
    sol1 := (-b)/(2*a)
    Printf("Ha una soluzione reale: %.3f \n", sol1)
    if Soglia(sol1) {
      Printf("La soluzione %.3f è maggiore della soglia\n", sol1)
    }
  } else {
    Printf("Non ha soluzioni reali.\n")
  }
}

func Soglia(sol float64) bool {
  soglia,_ := strconv.ParseFloat(os.Args[2], 64)
  epsilon,_ := strconv.ParseFloat(os.Args[3], 64)

  if sol > (soglia+epsilon){
    return true
  } else {
    return false
  }
}