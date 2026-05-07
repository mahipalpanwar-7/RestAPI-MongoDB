package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

// delete one record

func deleteOneRecord(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie got deleted with delete count:", deleteCount)
}

// delete all records from mongoDB

func deleteAllMovie() int64 {
	// filter := bson.D{{}}  // {} all will be selected
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("number of movies delete:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}


// get all movies from database 

func getAllMovies() []primitive.M{
	cur,err := collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err!= nil{
			log.Fatal(err)
		}

		movies = append(movies, primitive.M(movie))
	}

	defer cur.Close(context.Background())
    return movies
}

// actual controller for all movies to be exported

func GetMYAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-from-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}