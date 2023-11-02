package integrais

// Term Polynomial struct
type Term struct {
	Numerador   int
	Denominador int
	Expoente    int
}

// Function that receives a Term (polynomial) and returns the integrated Term
func Integral(term Term) Term {
	var Novo_denominador int
	var Novo_expoente int

	if term.Expoente != -1 {
		Novo_denominador = term.Denominador * (term.Expoente + 1)
		Novo_expoente = term.Expoente + 1
	} else {
		Novo_denominador = term.Denominador
		Novo_expoente = term.Expoente + 1
	}
	return Term{term.Numerador, Novo_denominador, Novo_expoente}
}
