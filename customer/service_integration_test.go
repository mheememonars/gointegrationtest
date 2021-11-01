package customer_test

import (
	"bytes"
	"context"
	"encoding/json"
	"go_integration_test/config"
	"go_integration_test/customer"
	"go_integration_test/db"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var (
	client *mongo.Client
	dbName string
)

func TestMain(m *testing.M) {
	setup()
	cleanDatabase()
	code := m.Run()
	cleanDatabase()
	os.Exit(code)
}

func setup() {
	config.Setup("integrationtest", "../env")
	dbName = viper.GetString("DB.Name")
	ConnectDBTest()
}

func ConnectDBTest() {
	var err error
	client, err = db.ConnectionWithAuth()
	if err != nil {
		log.Panic(err)
	}
	db.SetupCollections(client)
}

func cleanDatabase() {
	colls, _ := client.Database(dbName).ListCollectionNames(context.TODO(), primitive.D{})
	for _, c := range colls {
		client.Database(dbName).Collection(c).DeleteMany(context.TODO(), bson.M{})
	}
}

func GetRequestBody(body string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBuffer([]byte(body)))
}

func GetJsonToMap(body string) map[string]interface{} {
	var response map[string]interface{}
	json.Unmarshal([]byte(body), &response)
	return response
}

func Test_CustomerAdd(t *testing.T) {
	req, _ := http.NewRequest("POST", "/customer/add", GetRequestBody(`{
		"customer_id":    "001",
		"name":           "ชินจัง",
		"lastname": "โนฮาระ"
	}`))
	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(customer.CustomerAddHandler)
	handler.ServeHTTP(resp, req)

	expectedResponseCode := 201
	expectedResponseBody := GetJsonToMap(`{
		"status":  "Pass"
	}`)

	var actualStatusCode = resp.Code
	var actualResponseBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &actualResponseBody)

	assert.Equal(t, expectedResponseCode, actualStatusCode)
	assert.Equal(t, expectedResponseBody, actualResponseBody)
}

func Test_CustomerFind(t *testing.T) {
	req, _ := http.NewRequest("POST", "/customer/find", GetRequestBody(`{
		"customer_id":    "001"
	}`))
	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(customer.CustomerFindHandler)
	handler.ServeHTTP(resp, req)

	expectedResponseCode := 200
	expectedResponseBody := GetJsonToMap(`{
		"customer_id":    "001",
		"name":           "ชินจัง",
		"lastname": "โนฮาระ"
	}`)

	var actualStatusCode = resp.Code
	var actualResponseBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &actualResponseBody)

	assert.Equal(t, expectedResponseCode, actualStatusCode)
	assert.Equal(t, expectedResponseBody, actualResponseBody)
}
