package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "command line application"
	app.Usage = "search IPS and server names on the internet"

	flags :=  []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br",
				},
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "IPS lookup of internet addresses",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name: "servers",
			Usage: "search the name servers of internet",
			Flags: flags,
			Action: searchServers,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")


	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}