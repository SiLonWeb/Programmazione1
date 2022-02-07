package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	x, y      float64
}

type Segmento struct {
	punto1, punto2 Punto
}

func main() {
	soglia, _ := strconv.ParseFloat(os.Args[1], 64)
	punti := Leggi()
	segmenti := GeneraSegmenti(punti)

	for _, seg := range segmenti {
		valida := ValidaSegmento(seg, soglia)
		if valida {
			Println(StringSegmento(seg))
		}
	}
}

func StringPunto(p Punto) string {
	return p.etichetta + " = (" + strconv.FormatFloat(p.x, 'f', 2, 64) + ", " + strconv.FormatFloat(p.y, 'f', 2, 64) + ")"
}

func StringSegmento(s Segmento) string {
	return "Segmento con estremi " + StringPunto(s.punto1) + " e " + StringPunto(s.punto2) + "."
}

func ValidaSegmento(seg Segmento, soglia float64) bool {
	valido := true
	dist := Distanza(seg.punto1, seg.punto2)
	if dist >= soglia {
		valido = false
	}
	//Println(valido, seg, dist)
	if seg.punto1.x == seg.punto2.x {
		// parallelo y
		valido = false
	}
	if seg.punto1.y == seg.punto2.y {
		// parallelo x
		valido = false
	}
	if (seg.punto1.x * seg.punto2.x) < 0 {
		// segno x diverso --> quadrante diverso
		valido = false
	}
	if (seg.punto1.y * seg.punto2.y) < 0 {
		// segno y diverso --> quadrante diverso
		valido = false
	}
	return valido
}

func Leggi() (punti []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ";")
		if len(riga) > 0 {
			etichetta := riga[0]
			x, _ := strconv.ParseFloat(riga[1], 64)
			y, _ := strconv.ParseFloat(riga[2], 64)
			punti = append(punti, Punto{etichetta, x, y})
		}
	}
	return
}

func GeneraSegmenti(punti []Punto) (segmenti []Segmento) {
	for i, p1 := range punti {
		for _, p2 := range punti[i+1:] {
			segmenti = append(segmenti, Segmento{p1, p2})
		}
	}
	return
}

func Distanza(p1, p2 Punto) (euclidea float64) {
	euclidea = math.Sqrt((math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2)))
	return
}
