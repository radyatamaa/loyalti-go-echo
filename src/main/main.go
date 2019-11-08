package main

import (
	"../router"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/api/host"
)

func main() {
	fmt.Println("Welcome to the webserver")
	e := router.New()

	e.Start(":8000")

	host.StartKafka()
}
