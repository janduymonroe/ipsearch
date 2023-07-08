package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Search IP"
	app.Usage = "Search IP and server name on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPS on the internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servers",
			Usage:  "Find out servers name",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, errors := net.LookupIP(host)

	if errors != nil {
		log.Fatal(errors)
	}

	fmt.Println("Host Ips:\n")
	for _, ip := range ips {
		fmt.Println(ip)
	}
	fmt.Println()
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, errors := net.LookupNS(host)

	if errors != nil {
		log.Fatal(errors)
	}

	fmt.Println("Servers Name:\n")
	for _, server := range servers {
		fmt.Println(server)
	}
	fmt.Println()
}
