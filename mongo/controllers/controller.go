package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hydraveer/workingwithmongo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const MONGO_URL = "mongodb://localhost:27017/"
const DBName = "netflix"
const colName = "watchlist"

// collection for interacting with mongo (Most important)
var collections *mongo.Collection

// connecting with mongodb
func init() {
	//creating client options, option is a set of instruction that
	// tells your mongoDb that where we should look for database and how to connect
	clientOption := options.Client().ApplyURI(MONGO_URL)
	//connect to the database
	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal("Opps faild to connect with databases.")
	}
	fmt.Println("database connected")
	//collection instance
	collections = client.Database(DBName).Collection(colName)

	fmt.Println("collection instance ready for me")
}

//mongoDB helpers - file

// insert 1 document in mongodb
func insertOneMovie(movie models.Netflix) {
	inserted, err := collections.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The inserted value is ", inserted.InsertedID)
}
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "watched", Value: true}}}}
	result, err := collections.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", result.ModifiedCount)
}
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := collections.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the delecount is ", result.DeletedCount)
}
func deleteAllMovie() {

	result, err := collections.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The deleted row count ", result.DeletedCount)
}
func getAllMovie() []primitive.M {
	cursor, err := collections.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie primitive.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

//actual controller -file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-form-urlencode")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}
func CreateMoview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
	json.NewEncoder(w).Encode("Movie Insert Successfully!")
}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	// Call the update function and handle any errors
	updateOneMovie(id)
	// Send a success response
	response := map[string]string{"message": "Marked as watched"}
	json.NewEncoder(w).Encode(response)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("One Movie deleted from watchlist")
}
func DeleteAllVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllMovie()
	json.NewEncoder(w).Encode("All Movie deleted from watchlist and count is")
}
