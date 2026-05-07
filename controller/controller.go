package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/mahipalpanwar-7/RestAPI-MongoDB/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
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

// Helper function , insert one record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("one movie inserted in database with id:", inserted.InsertedID)
}

// update one movie

func updateOneMovie(movieID string){
	id,_ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id":id}
	update := bson.M{"$set": bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(),filter,update)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println("modified count:",result.ModifiedCount)
}

