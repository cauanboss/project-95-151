package controllers

import (
	"api-test/core/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	ctx := context.TODO()
	filter := bson.M{"Name": "Cauan"}

	u := models.User{}

	uc.client.Database("Udemy").Collection("users").FindOne(ctx, filter).Decode(&u)

	if err := uc.client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("Erroooooo")
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	uc.client.Database("Udemy").Collection("users").InsertOne(context.TODO(), u)

	s, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	fmt.Println(r.URL.Query())

	// uc.client.Database("Udemy").Collection("users").DeleteOne(context.TODO(), filter)

	// s, err := json.Marshal(u)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(s))

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
// 	var cursor *mongo.Cursor

// 	filter := bson.M{"Name": "Cauan"}

// 	u := models.User{}

// 	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
// 	uc.client.Database("Udemy").Collection("users").Find(ctx, filter)

// 	if err := uc.client.Ping(ctx, readpref.Primary()); err != nil {
// 		fmt.Println("Erroooooo")
// 	}

// 	defer cursor.Close(context.Background())
// 	for cursor.Next(context.Background()) {
// 		cursor.Decode(&u)
// 		fmt.Println(u)
// 	}

// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(filter)
// }
