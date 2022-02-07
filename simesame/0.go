package main

import(
  . "fmt"
  "os"
  "strconv"
)

func main() {
  Stampa(LeggiOS())
}

func LeggiOS() (output []int) {
  lettura := os.Args[1]
  for _,v := range lettura {
    num, err := strconv.Atoi(string(v))
    if err == nil {
      output = append(output, num)
    }
  }
  return
}

func Stampa(nums []int) {
  for i,v := range nums[:len(nums)-1] {
    if v < nums[i+1] {
      Print(v)
    }
  }
  Println()
  return
}