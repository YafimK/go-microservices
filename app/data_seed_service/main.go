package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Yafimk/go-microservices/document-service/common"
	"github.com/Yafimk/go-microservices/document-service/document_service/model"
	"log"
	"strconv"
)

func SeedDbData(db *common.DbDriver, bucketName string, amount int, fakeDataGeneratorFunc func(seed int) []byte) error {
	for i := 0; i < amount; i++ {
		key := strconv.Itoa(10000 + i)
		value := fakeDataGeneratorFunc(i)
		if err := db.AddValue(bucketName, []byte(key), value); err != nil {
			return err
		}
	}
	fmt.Printf("Seeded %v fake accounts...\n", amount)
	return nil
}

func main() {

	DbName := flag.String("db", "db", "database name")
	BucketName := flag.String("bucket", "documents", "table (bucket) name")
	flag.Parse()
	dbClient, err := common.NewDbDriver(*DbName)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = SeedDbData(dbClient, *BucketName, 1000, func(seed int) []byte {
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
}
