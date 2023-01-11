package people

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Person struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Age      int                `bson:"age,omitempty"`
	FullName string             `bson:"full_name,omitempty"`
}

type People []*Person

func GetPeople() []primitive.M {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// Initialising our router
	//r := gin.Default()

	usersCollection := client.Database("GoLang").Collection("users")

	cursor, err := usersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var people []primitive.M
	for cursor.Next(context.Background()) {
		var person bson.M
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}
	defer cursor.Close(context.Background())

	return people
}

func AddPerson(Person Person) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// Initialising our router
	//r := gin.Default()

	usersCollection := client.Database("GoLang").Collection("users")

	cursor, err := usersCollection.InsertOne(context.Background(), Person)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cursor.InsertedID)

}
func CreatePerson(Person Person) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// Initialising our router
	//r := gin.Default()

	usersCollection := client.Database("GoLang").Collection("users")

	cursor, err := usersCollection.InsertOne(context.Background(), Person)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(cursor.InsertedID)
}

func DelPerson(id int) int64 {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// Initialising our router
	//r := gin.Default()

	usersCollection := client.Database("GoLang").Collection("users")

	per := usersCollection.FindOne(context.Background(), id)

	cursor, err := usersCollection.DeleteOne(context.Background(), per)
	if err != nil {
		log.Fatal(err)
	}
	return cursor.DeletedCount
}
