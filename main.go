package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}

/*
comando para busca de host:
go run main.go ip --host mercadolivre.com

comando para buscar os servidores:
go run main.go servidores --host mercadolivre.com.br
*/