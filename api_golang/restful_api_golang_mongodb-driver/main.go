package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.k00qa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person", GetUserEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetOneUserEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", UpdateUserEndpoint).Methods("PUT")
	router.HandleFunc("/person/{id}", DeleteUserEndpoint).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

func CreatePersonEndpoint(response http.ResponseWriter, r *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	collection := client.Database("golang_mongodb_api").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}
func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var user []Person
	collection := client.Database("golang_mongodb_api").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{message":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close((ctx))
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		user = append(user, person)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetOneUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("golang_mongodb_api").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(person)
}
func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	opts := options.FindOneAndUpdate().SetUpsert(true)
	var personUpdated Person
	json.NewDecoder(r.Body).Decode(&personUpdated)
	collection := client.Database("golang_mongodb_api").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	update := bson.M{"$set": personUpdated}
	err := collection.FindOneAndUpdate(ctx, Person{ID: id}, update, opts).Decode(&personUpdated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(personUpdated)
}

func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("golang_mongodb_api").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.FindOneAndDelete().
		SetProjection(bson.D{{"firstname", 1}})
	var deleteDocument bson.M
	err := collection.FindOneAndDelete(
		ctx, Person{ID: id}, opts,
	).Decode(&deleteDocument)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(deleteDocument)

}
