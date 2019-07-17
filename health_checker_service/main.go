package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Yafimk/go-microservices/common/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	host := flag.String("dbHost", "http://localhost:8082", "host address (including protocol)")
	flag.Parse()
	bindAddresss, err := url.Parse(*host)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if bindAddresss.Scheme == "" {
		log.Fatalln("Missing protocol in bind address. host address should be in the following format <protocol://host:port>")
	}
	dbHealthCheckUrl := fmt.Sprintf("%v/HealthCheck", bindAddresss)
	response, err := http.Get(dbHealthCheckUrl)
	if err != nil || response.StatusCode != 200 {
		os.Exit(1)
	}
	if response != nil {
		healthCheck := model.HealthCheck{}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(body, &healthCheck)
		fmt.Println(healthCheck)
	}

	os.Exit(0)
}
