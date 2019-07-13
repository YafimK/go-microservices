package main

import (
	"encoding/json"
	"fmt"
	"github.com/Yafimk/go-microservices/common"
	"github.com/Yafimk/go-microservices/common/test"
	"github.com/Yafimk/go-microservices/document_service/model"
	"github.com/Yafimk/go-microservices/document_service/service"
	"log"
	"strconv"
)

const appName = "simple_service"

func main() {

	fmt.Printf("Starting %v\n", appName)
	dbClient, err := common.NewDbDriver("simple_service")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = dbClient.AddBucket("documents")
	if err != nil {
		log.Fatalln(err)
	}
	err = test.SeedDbData(dbClient, "documents", 1000, func(seed int) []byte {
		doc := model.Document{
			Id:   strconv.Itoa(seed),
			Name: "Person_" + strconv.Itoa(seed),
			Type: "regular",
		}
		jsonBytes, _ := json.Marshal(doc)
		return jsonBytes
	})
	if err != nil {
		log.Fatalln(err)
	}

	webService := service.NewService(":8080")
	webService.Start()

}
