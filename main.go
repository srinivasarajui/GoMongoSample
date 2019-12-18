package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	//Employee for testing
	Employee struct {
		Name string
		Age  int
	}
)

func sampleBson() {

}

func connect() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://user1:password1@localhost:27017/testing"))
	if err != nil {
		return client, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return client, err
	}
	err = client.Ping(ctx, readpref.Primary())
	return client, err
}
func sampleOperationsWithJSON(client *mongo.Client) {
	//DROP a old collection
	collection := client.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection.Drop(ctx)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		fmt.Println(err)
		return
	}
	id := res.InsertedID.(primitive.ObjectID)

	fmt.Println("JSON insert:", id)
	cur, err := collection.Find(ctx, bson.D{})
	defer cur.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("bson Item:", result)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}

func sampleOperationsWithStruct(client *mongo.Client) {
	//DROP a old collection
	collection := client.Database("testing").Collection("employees")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	collection.Drop(ctx)

	emp1 := &Employee{
		Name: "Emp1",
		Age:  100,
	}
	res, err := collection.InsertOne(ctx, emp1)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := res.InsertedID.(primitive.ObjectID)

	fmt.Println("Employee insert Id", emp1.Name, id)

	emp2 := &Employee{
		Name: "Emp2",
		Age:  99,
	}
	res, err = collection.InsertOne(ctx, emp2)
	if err != nil {
		fmt.Println(err)
		return
	}
	id = res.InsertedID.(primitive.ObjectID)

	fmt.Println("Employee insert Id", emp2.Name, id)

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Employee
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Name, result.Age)
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	var emp Employee
	err = collection.FindOne(ctx, bson.D{{Key: "name", Value: "Emp2"}}).Decode(&emp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("find one", emp.Name, emp.Age)
}
func main() {
	client, err := connect()
	if err != nil {
		fmt.Println("Unable to connect", err)
		return
	}
	sampleOperationsWithJSON(client)
	sampleOperationsWithStruct(client)
}
