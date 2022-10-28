package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI -- ask the user to log in instead
// this is the database
const uri = "mongodb+srv://@cluster0.erkptpf.mongodb.net/?retryWrites=true&w=majority"

type plot struct {
	field map[string]int
}
type RDY_plants struct {
	RdyPlt     string
	RdyPlt_amt int
}
type GRW_plants struct {
	GrwPlt     string
	GrwPlt_amt int
}
type user struct {
	Farm     [10]plot     `json:"farm"`
	Iventory []RDY_plants `json:"Iventory"`
	SeedIvn  []GRW_plants `json:"seedIvn"`
}

func main() {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Successfully connected and pinged.")

	collection := client.Database("atna_db").Collection("atna_frm_land")

	p1 := &user{Farm: [10]plot{}, Iventory: []RDY_plants{}, SeedIvn: []GRW_plants{}}

	insertentry, err := collection.InsertOne(context.TODO(), p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertentry.InsertedID)

}