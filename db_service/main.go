package main

import (
	"flag"
	"fmt"
	"github.com/Yafimk/go-microservices/db_service/service"
	"log"
	"net/url"
)

const appName = "DB_SERVICE"

var bucketList = []string{
	"documents",
}

func main() {
	address := flag.String("bind", "http://localhost:8082", "database service bind address")
	dbName := flag.String("db", "db", "database name")
	flag.Parse()

	fmt.Printf("Starting %v on %v\n", appName, *address)
	bindAddresss, err := url.Parse(*address)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if bindAddresss.Scheme == "" {
		log.Fatalln("Missing protocol in bind address. host address should be in the following format <protocol://host:port>")
	}

	webService := service.NewService(bindAddresss.Host, *dbName, "documents")
	webService.Start()

}
