package main

func main(){
	funcinarios := map[string]string{
		"João": "Gerente",
		"Maria": "Analista",
		"Pedro": "Desenvolvedor",
	}

	for nome, cargo := range funcinarios {
		println(nome, "-", cargo)
	}
}