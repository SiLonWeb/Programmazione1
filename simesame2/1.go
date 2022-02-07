package main

import(
  ."fmt"
  "unicode"
)

func main() {
  password := ""
  Scan(&password)
  errori_pw := VerificaPassword(password) 
  if len(errori_pw) > 0 {
    Println("La pw non è definita correttamente:")
    for _,nota := range errori_pw {
      Println(nota)
    }
  } else {
    Println("La pw è ben definita!")
  }

}

func VerificaPassword(pw string) []string{
  errori := []string{}
  maiusc, minusc, numeri, altro := 0, 0, 0, 0

  arrayRune := []rune(pw)
  if len(arrayRune) < 12 {
    errori = append(errori, "- La pw deve avere una lunghezza minima di 12 caratteri")
  }

  for _,lettera := range pw {
    if unicode.IsLetter(lettera) && unicode.IsLower(lettera) {
      minusc++
    } else if unicode.IsLetter(lettera) && unicode.IsUpper(lettera) {
      maiusc++
    } else if unicode.IsDigit(lettera) {
      numeri++
    } else {
      altro++
    }
  }

  if minusc < 2 {
    errori = append(errori, "- Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole")
  }
  if maiusc < 2 {
    errori = append(errori, "- Almeno 2 caratteri della pw devono rappresentare delle lettere maiuscole")
  }
  if numeri < 3 {
    errori = append(errori, "- Almeno 3 caratteri della pw devono rappresentare delle cifre decimali")
  }
  if altro < 4 {
    errori = append(errori, "- Almeno 4 caratteri della pw non devono rappresentare lettere o cifre decimali")
  }
  return errori
}