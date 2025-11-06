package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {
	// Determine Elasticsearch URL (allow overriding with ELASTIC_URL for local runs)
	elasticURL := os.Getenv("ELASTIC_URL")
	if elasticURL == "" {
		elasticURL = "http://elasticsearch:9200"
	}
	log.Printf("Using Elasticsearch URL: %s", elasticURL)

	// Create Elasticsearch client
	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Create context
	ctx := context.Background()

	// Create an index
	_, err = client.CreateIndex("people").Do(ctx)
	if err != nil && !elastic.IsStatusCode(err, 400) {
		log.Fatalf("Error creating index: %s", err)
	}

	// Create a person document
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}

	// Index the document
	_, err = client.Index().
		Index("people").
		Id("1").
		BodyJson(person).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}

	fmt.Println("Successfully indexed document")

	// Get the document
	result, err := client.Get().
		Index("people").
		Id("1").
		Do(ctx)
	if err != nil {
		log.Fatalf("Error getting document: %s", err)
	}

	// Parse the document
	var foundPerson Person
	err = json.Unmarshal(result.Source, &foundPerson)
	if err != nil {
		log.Fatalf("Error parsing document: %s", err)
	}

	fmt.Printf("Found person: %+v\n", foundPerson)
}
