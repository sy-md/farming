package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://mdorsett:12345@cluster0.erkptpf.mongodb.net/?retryWrites=true&w=majority"
const add_user = "PrFGBAfshdio6TnQRfaWR0ENGqp3kqGYfiv0OaShnDhxhOiL3PFGyyz3kiOGG4xz"
const get_user = "7gHMdgjBkauShRwuyj3JXPCZesPdjI2hlFtvTFOl4PJ3tWK1LqjRPCpj926zKf5M"

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
	Name     string       `json:"userName"`
	Farm     []plot       `json:"farm"`
	Iventory []RDY_plants `json:"Iventory"`
	SeedIvn  []GRW_plants `json:"seedIvn"`
	pswd     string       `json:"Password"`
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

	fmt.Println("Welcome to the farm clodhopper.")

	collection := client.Database("atna_db").Collection("atna_frm_land")

	var username string
	fmt.Println("username: ")
	fmt.Scan(&username)

	new_p1 := &user{Name: username, Farm: []plot{}, Iventory: []RDY_plants{}, SeedIvn: []GRW_plants{}}

	insertentry, err := collection.InsertOne(context.TODO(), new_p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("add user to the database", insertentry.InsertedID)

}

func login_user() {
	var username string
	fmt.Println("username: ")
	var password string
	fmt.Println("password: ")
	fmt.Scan(&username)
	fmt.Scan(&password)

	credential := options.Credential{
		AuthSource: uri,
		Username:   username,
		Password:   password,
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func main() {
	fmt.Println("starting app...")

	var has_act string
	fmt.Println("do u have a account y/n: ")
	fmt.Scan(&has_act)

	if has_act != "y" {
		new_user()
	} else {
		login_user()
	}

	var options int
	fmt.Println(("plant:1, water:2, sell:3 ->"))
	fmt.Scan(&options)
	switch options {
	case 1:
		fmt.Println("plant")
	case 2:
		fmt.Println("water")
	case 3:
		fmt.Println("sell")
	}

}
