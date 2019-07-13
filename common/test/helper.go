package test

import (
	"github.com/Yafimk/go-microservices/common"
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
	log.Printf("Seeded %v fake accounts...\n", amount)
	return nil
}
