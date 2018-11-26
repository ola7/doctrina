package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"../model"
	"github.com/boltdb/bolt"
)

// IBoltClient acts as an interface that enables mocking
type IBoltClient interface {
	OpenBoltDb()
	QueryUser(userID string) (model.User, error)
	SeedFakeUsers(n int)
	CheckStatus() bool
}

// BoltClient is the real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDb opens or creates
func (bc *BoltClient) OpenBoltDb() {
	var err error
	// Open() will create if it doesn't exist
	bc.boltDB, err = bolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// SeedFakeUsers seeds a bunch of fake users into DB
func (bc *BoltClient) SeedFakeUsers(n int) {
	createdNew := bc.initializeBucket("UserBucket")
	if createdNew {
		bc.seedUsers(n)
	}
}

// CheckStatus is a simple status check to be used by health checks
func (bc *BoltClient) CheckStatus() bool {
	return bc.boltDB != nil
}

// initializeBucket creates (and override) a bucket in our DB
func (bc *BoltClient) initializeBucket(bucketName string) bool {
	createdNew := false
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		if tx.Bucket([]byte(bucketName)) != nil {
			log.Println("Using existing bucket:", bucketName)
			createdNew = false
		} else {
			_, err := tx.CreateBucket([]byte(bucketName))
			createdNew = true
			if err != nil {
				log.Printf("Create bucket '%v' failed:%v\n", bucketName, err)
			} else {
				log.Println("Created new bucket:", bucketName)
			}
		}
		return nil
	})
	return createdNew
}

// Seed (n) make-believe user objects into the UserBucket bucket.
func (bc *BoltClient) seedUsers(n int) {

	total := n
	for i := 0; i < total; i++ {

		// generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// create an instance of our User struct
		acc := model.User{
			Id:   key,
			Name: "Name_" + strconv.Itoa(i),
		}

		// serialize the struct to json
		jsonBytes, _ := json.Marshal(acc)

		// write the data to the UserBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("UserBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	log.Printf("Seeded fake users: %v users\n", total)
}

// QueryUser finds a User from its id
func (bc *BoltClient) QueryUser(userID string) (model.User, error) {

	// empty User instance to be populated
	user := model.User{}

	// read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("UserBucket"))
		userBytes := b.Get([]byte(userID))
		if userBytes == nil {
			return fmt.Errorf("No user with ID '%v' found", userID)
		}
		json.Unmarshal(userBytes, &user)
		return nil
	})

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
