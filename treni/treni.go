package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mr := make(map[string][]int)
	docs := make([]string, 0)

	l := 0
	for scanner.Scan() {
		doc := scanner.Text()
		docs = append(docs, doc)
		a := strings.Split(strings.ToLower(doc), " ")
		for _, s := range a {
			if len(s) == 0 {
				continue
			}
			if lmr := len(mr[s]); lmr == 0 || mr[s][lmr-1] != l {
				mr[s] = append(mr[s], l)
			}
		}
		l++
	}
	keyword := strings.ToLower(os.Args[1])

	for _, l := range mr[keyword] {
		doc:=docs[l]
		pos :=strings.Index(strings.ToLower(doc), keyword)
		Print(l, "\t->\t", doc[:pos], "\033[31;1;4m", doc[pos:pos+len(keyword)], "\033[0m", doc[pos+len(keyword):], "\n");
	}
}
