package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
  "sort"
)

type Comando struct {
	direzione string
	passi     int
}

func main() {
	listaComandi := LeggiComandi()
  comandiAnalizzati := AnalizzaComandi(listaComandi)
  percorsoInverso := PercorsoInverso(listaComandi)

	Println("Movimenti totali:")
	for dir, num := range comandiAnalizzati {
		Println(dir, num)
	}
	Println("La direzione in cui vengono compiuti più passi è:")
	Println(SpostamentoMaggiore(comandiAnalizzati))
  Println("Lista comandi inversi:")
  for _,com := range percorsoInverso {
    Println(com.direzione, com.passi)
  }
}

func LeggiComandi() (comandi []Comando) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		movimento := strings.Split(scanner.Text(), " ")
		if len(movimento) > 0 {
			d := movimento[0]
			p, _ := strconv.Atoi(movimento[1])
			comandi = append(comandi, Comando{d, p})
		}
	}
	return
}

func AnalizzaComandi(comandi []Comando) map[string]int {
	numeroComandi := map[string]int{}
	for _, comando := range comandi {
		numeroComandi[comando.direzione] += comando.passi
	}
	return numeroComandi
}

func PercorsoInverso(comandi []Comando) (comandiInversi []Comando) {
  inverso := ""
  for i := len(comandi)-1; i>=0; i--{
    switch comandi[i].direzione {
      case "NORD":
        inverso = "SUD"
      case "SUD":
        inverso = "NORD"
      case "OVEST":
        inverso = "EST"
      case "EST":
        inverso = "OVEST"
    }
    comandiInversi = append(comandiInversi, Comando{inverso, comandi[i].passi})
  }
  return
}

func SpostamentoMaggiore(analisiComandi map[string]int) string {
	passi := []int{}
  for _, v := range analisiComandi {
    passi = append(passi, v)
	}
  sort.Ints(passi)
  max := passi[len(passi)-1]
  direzioneMax := []string{}
  for dir,val := range analisiComandi{
    if val == max {
      direzioneMax = append(direzioneMax, dir)
    }
  }
  sort.Strings(direzioneMax)
  return direzioneMax[0]
}