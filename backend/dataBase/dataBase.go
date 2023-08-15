package database

import (
	"backend/internal/models"
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DataBase struct {
	client         *mongo.Client
	userCollection *mongo.Collection
}

type rmUser struct {
	Id string
}

var Db DataBase

func (Db *DataBase) Initialization() (*mongo.Client, error) {

	var err error
	Db.client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&AppName=mongosh+1.8.2"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Db.client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// read from command line
	if err := Db.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	// BELOW THIS IS USED FOR DATABASE NAME AND COLLECTION

	// this is to store signup part
	Db.userCollection = Db.client.Database("signup").Collection("userData")

	return &mongo.Client{}, nil
}

// BELOW CODE IS USED TO ADD NEW USER TO THE DATABASE
func (Db *DataBase) AddUser(user *models.User) (*mongo.Collection, error) {

	addUser, err := Db.userCollection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println("add User not perfect")
	}
	fmt.Println(addUser.InsertedID)

	return &mongo.Collection{}, nil
}

// BELOW CODE IS USED TO REMOVE EXISITING USER FROM THE DATABASE
func (Db *DataBase) RemoveUser(Id *rmUser) {
	removeUser, _ := Db.userCollection.DeleteOne(context.Background(), Id)
	fmt.Println(removeUser.DeletedCount)
}

// THIS HANDLES LOGIN CREDENTIALS
func (Db *DataBase) ValidateUser(validUser *models.LoginCredentials) (*mongo.Collection, error) {
	valiDate, _ := Db.userCollection.Find(context.TODO(), bson.M{"objectId_": validUser})
	if valiDate != nil {
		fmt.Printf("user exists")
	} else {
		fmt.Println("user not exists")
	}
	return &mongo.Collection{}, nil
}