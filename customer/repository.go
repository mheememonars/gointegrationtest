package customer

import (
	"go_integration_test/db"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
}

func (r *Repo) collection() *mongo.Collection {
	coll := db.CustomerCollection
	return coll
}

func (r *Repo) findOne(query map[string]interface{}) (CustomerData, error) {
	log.Println("findOne")
	var result CustomerData
	err := db.FindOne(query, r.collection(), &result)
	return result, err
}

func (r *Repo) save(customerData CustomerData) (CustomerData, error) {
	log.Println("save")
	_, err := db.Save(customerData, r.collection())
	return customerData, err
}
