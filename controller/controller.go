package controller

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://mahipal:mongodb%40123123@cluster0.qskv2lv.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// db reference

var collection *mongo.Collection

// connecting with MongoDB

// init runs first and only one time
func init() {
	// client option

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb conncetion is successful ")

	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("collection instance is ready")
}
