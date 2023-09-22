package main

import (
	"GoRealEstateManagement/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}
	a.Routes()
	a.Run()
}
