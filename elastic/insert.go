package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/learn-golang/elastic/models"
	elastic "gopkg.in/olivere/elastic.v7"
)

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthCheck(false))

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("ES initialized...")

	return client, nil

}

func main() {
	ctx := context.Background()
	esclient, err := GetESClient()

	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	var mod = models.Student{"Gopher Doe", 10, 99.9}

	dataJson, err := json.Marshal(mod)
	js := string(dataJson)

	ind, err := esclient.Index().
		Index("students").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}
