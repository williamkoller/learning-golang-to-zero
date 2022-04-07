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

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "IPS lookup of internet addresses",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
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
