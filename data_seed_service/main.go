package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Yafimk/go-microservices/common/model"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type FakeDataGenerator func(seed int) []byte
type Requester func(dbAddress string, data []byte) error

func sendRequest(dbAddress string, data []byte) error {
	resp, err := http.Post(fmt.Sprintf("%v/documents/Add", dbAddress), "application/json", bytes.NewReader(data))
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed ")
	}
	return nil
}

func seedDbData(dbAddress string, amount int, fakeDataGenerator FakeDataGenerator, requester Requester) error {
	for i := 0; i < amount; i++ {
		value := fakeDataGenerator(i)
		err := requester(dbAddress, value)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Printf("Seeded %v fake accounts...\n", amount)
	return nil
}

func main() {
	dbAddress := flag.String("dbHost", "http://localhost:8082", "database address")
	flag.Parse()

	bindAddress, err := url.Parse(*dbAddress)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if bindAddress.Scheme == "" {
		log.Fatalln("Missing protocol in bind address. host address should be in the following format <protocol://host:port>")
	}

	err = seedDbData(bindAddress.String(), 1000, func(seed int) []byte {
		doc := model.Document{
			Id:   strconv.Itoa(seed),
			Name: "Person_" + strconv.Itoa(seed),
			Type: "regular",
		}
		jsonBytes, _ := json.Marshal(doc)
		return jsonBytes
	}, sendRequest)
	if err != nil {
		log.Fatalln(err)
	}
}
