package main

import (
	"fmt"
	"home/integrais"
)

const CaseLNX = 1

// Function that receives the input and returns the polynomial
func input(ordem int, quant int) ([]integrais.Term, int) {
	var termoArray []integrais.Term
	var termo integrais.Term

	// loop que recebe o input
	for i := 0; i < quant; i++ {
		fmt.Printf("Digite o coeficiente do %d° x: \n", i+1)
		fmt.Scan(&termo.Numerador)
		fmt.Printf("Digite o expoente do %d° x: \n", i+1)
		fmt.Scan(&termo.Expoente)
		termo.Denominador = 1
		termoArray = append(termoArray, termo)
	}

	var c int
	fmt.Print("Digite a constante: \n")
	fmt.Scan(&c)
	return termoArray, c
}

// Function that calls the function Integral once and prints the result
func printResult1Grau(terms []integrais.Term, c int) {
	// print of the original function
	var original string

	for _, term := range terms {
		if term.Numerador == 0 {
			original = original + ""
		} else {
			original = original + fmt.Sprintf(" %dx^(%d) +", term.Numerador, term.Expoente)
		}
	}

	original = original + fmt.Sprintf("%d\n", c)

	fmt.Print("y' = ", original)

	// print of the integrated function
	var resultado string
	var integral integrais.Term
	for _, term := range terms {
		integral = integrais.Integral(term)
		if integral.Numerador == 0 {
			resultado = resultado + ""
		} else {
			if integral.Expoente == 0 {
				resultado = resultado + fmt.Sprintf("( %dln(x) ) + ", integral.Numerador)
			} else {
				resultado = resultado + fmt.Sprintf(" (%d/%d)x^(%d) +", integral.Numerador, integral.Denominador, integral.Expoente)
			}
		}
	}
	if c != 0 {
		resultado = resultado + fmt.Sprintf("(%dx) + ", c)
	}

	resultado = resultado + fmt.Sprintf(" C\n")

	// prints the final result
	fmt.Print("y = ", resultado)
}

// Function that calls the function Integral two times and prints the result
func printResult2Grau(terms []integrais.Term, c int) {
	// print of the original function
	var original string

	for _, term := range terms {
		if term.Numerador == 0 {
			original = original + ""
		} else {
			original = original + fmt.Sprintf(" %dx^(%d) +", term.Numerador, term.Expoente)
		}
	}

	original = original + fmt.Sprintf(" %d\n", c)

	fmt.Print("y'' = ", original)

	// first integral
	var integral1 []integrais.Term
	for _, term := range terms {
		integral1 = append(integral1, integrais.Integral(term))
	}

	var termo_c integrais.Term
	termo_c.Numerador = c
	termo_c.Denominador = 1
	termo_c.Expoente = 1

	integral1 = append(integral1, termo_c)

	// second integral
	var resultado string
	var novo_termo integrais.Term
	for _, term := range integral1 {
		novo_termo = integrais.Integral(term)
		if novo_termo.Numerador == 0 {
			resultado = resultado + ""
		} else {
			if novo_termo.Expoente == CaseLNX {
				resultado = resultado + fmt.Sprintf(" %d(xln(x) - x) +", novo_termo.Numerador)
			} else if novo_termo.Expoente == 0 {
				resultado = resultado + fmt.Sprintf(" (%d/%d)ln(x) +", novo_termo.Numerador, novo_termo.Denominador)
			} else {
				resultado = resultado + fmt.Sprintf(" (%d/%d)x^(%d) +", novo_termo.Numerador, novo_termo.Denominador, novo_termo.Expoente)
			}
		}
	}

	fmt.Print("y = ", resultado, " C\n")
}

// Main function
func main() {
	var ordem int
	var c int
	var qntd_x int
	var term []integrais.Term
	fmt.Print("Digite a ordem da Equação: \n")
	fmt.Scan(&ordem)
	fmt.Print("Digite a quantidade de termos x: \n")
	fmt.Scan(&qntd_x)
	term, c = input(ordem, qntd_x)
	if ordem == 1 {
		printResult1Grau(term, c)
	} else if ordem == 2 {
		printResult2Grau(term, c)
	}
}
