package main

import (
	"ipsearch/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	if errors := application.Run(os.Args); errors != nil {
		log.Fatal(errors)
	}

}
