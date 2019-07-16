package main

import (
	"flag"
	"fmt"
	"github.com/Yafimk/go-microservices/document-service/service"
)

const appName = "DOCUMENT_SERVICE"

func main() {
	port := flag.Int("port", 8083, "port number")
	host := flag.String("host", "", "host address (including protocol)")
	flag.Parse()

	fmt.Printf("Starting %v\n", appName)

	webService := service.NewService(fmt.Sprintf("%v:%v", *host, *port))
	webService.Start()

}
