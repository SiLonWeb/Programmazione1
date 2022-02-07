package main

import(
  ."fmt"
  "bufio"
  "os"
)
func main() {
  Println("Inserisci testo:")
  Println("Testo letto:", LeggiTesto())
	
}

func LeggiTesto() (testo string) {
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    testo += scanner.Text() + "\n"
  }
  return
}