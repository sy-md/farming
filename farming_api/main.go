package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://mdorsett:12345@cluster0.erkptpf.mongodb.net/?retryWrites=true&w=majority"
const add_user = "PrFGBAfshdio6TnQRfaWR0ENGqp3kqGYfiv0OaShnDhxhOiL3PFGyyz3kiOGG4xz"
const get_user = "7gHMdgjBkauShRwuyj3JXPCZesPdjI2hlFtvTFOl4PJ3tWK1LqjRPCpj926zKf5M"

type RDY_plants struct {
	RdyPlt     string
	RdyPlt_amt int
}
type planted_seed struct {
	Title     string
	Plant     string
	Water_lvl int
}
type GRW_plants struct {
	GrwPlt []planted_seed
}
type Farm struct {
	Plot    []planted_seed `json:"Iventory"`
	SeedIvn []GRW_plants   `json:"seedIvn"`
	Stock   Store          `json:"myStockivn"`
}

type Store struct {
	Iventory []RDY_plants `json:"Iventory"`
	SeedIvn  []GRW_plants `json:"seedIvn"`
}
type User struct {
	Name     string `json:"userName"`
	Password string `json:"PassWord"`
	Myfarm   Farm   `json:"Myfarm"`
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
	//collection_store := client.Database("atna_db").Collection("store")
	collection_user := client.Database("atna_db").Collection("user")

	var username string
	fmt.Println("username: ")
	fmt.Scan(&username)
	var password string
	fmt.Println("password: ")
	fmt.Scan(&password)

	new_user_str := &Store{Iventory: []RDY_plants{}, SeedIvn: []GRW_plants{}}

	new_fld := planted_seed{Title: "dirt", Plant: "#", Water_lvl: 0}
	new_user_farm := &Farm{Plot: []planted_seed{new_fld}, SeedIvn: []GRW_plants{}, Stock: *new_user_str}
	new_act := User{Name: username, Password: password, Myfarm: *new_user_farm}

	new_user, err := collection.InsertOne(context.TODO(), new_user_farm)
	if err != nil {
		log.Fatal(err)
	}
	new_user_act, err := collection_user.InsertOne(context.TODO(), new_act)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("add user to the database", new_user.InsertedID)
	fmt.Println("add user to the database", new_user_act.InsertedID)
	show_farm(username)
	//fmt.Println("add user to the database", new_user_str.InsertedID)
}

func login_user() {
	fmt.Println("asked fo username and password")
}

func show_farm(username string) {

	fmt.Println("this is the current user", username)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	//corn := planted_seed{Title: "corn", Plant: "c", Water_lvl: 0}
	coll := client.Database("atna_db").Collection("user")

	filter := bson.M{"name": bson.M{"$eq": "kim"}}
	update := bson.M{"$set": bson.M{"name": "martell"}}

	// Call the driver's UpdateOne() method and pass filter and update to it
	result, err := coll.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	fmt.Println(result)
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
	fmt.Println(("plant:1, water:2, sell:3 -> "))
	fmt.Scan(&options)
	switch options {
	case 1:
		fmt.Println("plant")
	case 2:
		fmt.Println("water")
	case 3:
		fmt.Println("sell")
	case 4:
		fmt.Println("display farm")
	}

}
