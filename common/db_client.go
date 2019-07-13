package common

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type DbDriver struct {
	db *bolt.DB
}

func NewDbDriver(dbName string) (*DbDriver, error) {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &DbDriver{db: db}, nil
}

func (dbClient *DbDriver) AddBucket(bucketName string) error {
	return dbClient.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketName)); err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (dbClient *DbDriver) IsAlive() bool {
	return dbClient.db != nil
}

func (dbClient *DbDriver) IsBucketExists(bucketName string) bool {
	result := dbClient.db.View(func(tx *bolt.Tx) error {
		if result := tx.Bucket([]byte(bucketName)); result == nil {
			return fmt.Errorf("bucket %v doesn't exist\n", bucketName)
		}
		return nil
	})
	if result == nil {
		return true
	}
	return false
}

func (dbClient *DbDriver) AddValue(bucketName string, key []byte, value []byte) error {
	return dbClient.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		return bucket.Put(key, value)
	})
}

func (dbClient *DbDriver) QueryValue(bucketName string, key string, valueUnMarshalFunc func(value []byte) (interface{}, error)) (interface{}, error) {
	var value []byte
	err := dbClient.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		value = bucket.Get([]byte(key))
		if value == nil {
			return fmt.Errorf("no value found for the key %v", key)
		}
		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	if err != nil {
		return nil, err
	}
	return valueUnMarshalFunc(value)
}
