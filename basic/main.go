package main

import (
	"fmt"
	"github.com/Yafimk/go-microservices/basic/service"
)

const appName = "simple_service"

func main() {

	fmt.Printf("Starting %v\n", appName)
	webService := service.Service{}
	webService.Start(":8080")

}
