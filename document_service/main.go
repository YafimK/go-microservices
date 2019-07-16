package main

import (
	"flag"
	"fmt"
	"github.com/Yafimk/go-microservices/document-service/service"
	"log"
	"net/url"
)

const appName = "DOCUMENT_SERVICE"

func main() {
	host := flag.String("host", "http://localhost:8083", "bind address <protocol://host:port>")
	flag.Parse()
	fmt.Printf("Starting %v on %v\n", appName, *host)
	bindAddresss, err := url.Parse(*host)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if bindAddresss.Scheme == "" {
		log.Fatalln("Missing protocol in bind address. host address should be in the following format <protocol://host:port>")
	}
	webService := service.NewService(bindAddresss.Host)
	webService.Start()

}
