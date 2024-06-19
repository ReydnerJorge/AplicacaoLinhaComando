package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a app de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "globo.com.br",
	},
}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPS de endereço na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Buscar o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},

	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidores := range servidores {
		fmt.Println(servidores.Host)
	}
 }