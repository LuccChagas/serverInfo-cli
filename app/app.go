package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

var hostValue string = "www.google.com.br"

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Buscar IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: hostValue,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:    "servidores",
			Usage:   "Busca o nome dos servidores na internet",
			Aliases: []string{"-s"},
			Flags:   flags,
			Action:  buscarServidores,
		},
	}

	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	if c.String("host") == hostValue {
		fmt.Println("Host inalterado, buscando IP padrão de:", c.String("host"))
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
