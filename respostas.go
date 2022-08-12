package respostas

type Resposta struct {
	nome  string
	idade int
}

func retorna() string {
	print("Ola mundo")
	return "Ola mundo no return"
}
