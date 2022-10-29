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

type planted_seed struct { //seeeds the farm has planted
	Title     string
	Plant     string
	Water_lvl int
}
type Seeds struct { //seed that farmers can buy
	plants string
	amout  int
	cost   int
	price  int
}
type Store struct { // store with seed on hand for farmer to buy
	Iventory []Seeds `json:"Iventory"`
}
type Farm struct {
	Plot     []planted_seed `json:"fields"`
	SeedIvn  []Seeds        `json:"seedIvn"`
	plantIvn []planted_seed
	Stock    []Seeds `json:"myStockivn"`
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

	//new_user_str := &Store{Iventory: []Seeds{}}

	new_fld := planted_seed{Title: "dirt", Plant: "#", Water_lvl: 0}

	new_user_farm := &Farm{Plot: []planted_seed{new_fld}, SeedIvn: []Seeds{}, Stock: []Seeds{}, plantIvn: []planted_seed{}}
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

	//fmt.Println("add user to the database", new_user_str.InsertedID)
}

func show_farm() {

	//fmt.Println("this is the current user", username)
	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	//if err != nil {
	//	panic(err)
	//}

	//corn := planted_seed{Title: "corn", Plant: "c", Water_lvl: 0}
	//coll := client.Database("atna_db").Collection("user")

	//filter := bson.M{"name": bson.M{"$eq": "kim"}}
	//update := bson.M{"$set": bson.M{"name": "martell"}}

	// Call the driver's UpdateOne() method and pass filter and update to it
	//result, err := coll.UpdateOne(
	//	context.Background(),
	//	filter,
	//	update,
	//)
	//fmt.Println(result)
}

func start_app() {
	var username string
	fmt.Println("username: ")
	fmt.Scan(&username)
	var passwd string
	fmt.Println("password: ")
	fmt.Scan(&passwd)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	coll := client.Database("atna_frm_land").Collection("user")

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "Name", Value: username}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			fmt.Println("didnt find", username, "making new account..")
			new_fld := planted_seed{Title: "dirt", Plant: "#", Water_lvl: 0} //first plot
			new_user_farm := &Farm{Plot: []planted_seed{new_fld}, SeedIvn: []Seeds{}, Stock: []Seeds{}, plantIvn: []planted_seed{}}
			new_act := User{Name: username, Password: passwd, Myfarm: *new_user_farm}
			new_user_act, err := coll.InsertOne(context.TODO(), new_act)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("add user to the database", new_user_act.InsertedID)
		}
	}

	fmt.Println("starting app...")
	fmt.Println("connecting...")
	fmt.Println("welcome to the farm land", username)

	var opts string
	fmt.Println("show farm:1")
	fmt.Scan(&opts)
	switch opts {
	case "1":
		fmt.Println("displaying farm...")
	}

	//docs := []interface{} {
	//	bson.D{{"type", "English Breakfast"}, {"rating", 6}},
	//	bson.D{{"type", "Oolong"}, {"rating", 7}, {"vendor", bson.A{"C"}}},
	//	bson.D{{"type", "Assam"}, {"rating", 5}},
	//	bson.D{{"type", "Earl Grey"}, {"rating", 8}, {"vendor", bson.A{"A", "B"}}},
	//}

	//result, err := coll.InsertMany(context.TODO(), docs)

	//corn := planted_seed{Title: "corn", Plant: "c", Water_lvl: 0}

}

func main() {
	start_app()

}
