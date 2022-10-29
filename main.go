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
const uri = "mongodb+srv://mdorsett:12345@cluster0.erkptpf.mongodb.net/?retryWrites=true&w=majority"
const add_user = "PrFGBAfshdio6TnQRfaWR0ENGqp3kqGYfiv0OaShnDhxhOiL3PFGyyz3kiOGG4xz"

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
	Farm     []plot       `json:"farm"`
	Iventory []RDY_plants `json:"Iventory"`
	SeedIvn  []GRW_plants `json:"seedIvn"`
}

func new_user() { // Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//users connects
	fmt.Println("Successfully connected and pinged.")

	collection := client.Database("atna_db").Collection("atna_frm_land")

	new_p1 := &user{Farm: []plot{}, Iventory: []RDY_plants{}, SeedIvn: []GRW_plants{}}

	insertentry, err := collection.InsertOne(context.TODO(), new_p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("add user to the database", insertentry.InsertedID)

}

func main() {}
