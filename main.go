package main

import (
	"GoLangProject/handlers"
	"fmt"

	"context"

	"github.com/gin-gonic/gin"
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

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// Initialising our router
	r := gin.Default()

	// Routes with paths and their corresponding handler functions
	r.GET("/", handlers.GetPeople)
	r.GET("/person/:id", handlers.GetPersonByID)
	r.POST("/addPerson", handlers.AddPerson)
	r.DELETE("/person/:id", handlers.DelPerson)
	r.PUT("/person/:id", handlers.PutPerson)

	// Attaches our routers to http.Server to listen at port: 9090.
	r.Run("localhost:9000")

	usersCollection := client.Database("GoLang").Collection("users")

	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
}

// for Person, entry := range usersCollection {
// 	fmt.Println(entry.FullName)
// 	fmt.Println(entry.Age)
// 	fmt.Println(entry.ID)

//  }

// // retrieve single and multiple documents with a specified filter using FindOne() and Find()
// // create a search filer
// filter := bson.D{
// 	{"$and",
// 			bson.A{
// 					bson.D{
// 							{"age", bson.D{{"$gt", 25}}},
// 					},
// 			},
// 	},
// }

// // retrieve all the documents that match the filter
// cursor, err := usersCollection.Find(context.TODO(), filter)
// // check for errors in the finding
// if err != nil {
// 	panic(err)
// }

// // convert the cursor result to bson
// results []bson.M
// // check for errors in the conversion
// if err = cursor.All(context.TODO(), &results); err != nil {
// 	panic(err)
// }

// // display the documents retrieved
// fmt.Println("displaying all results from the search query")
// for _, result := range results {
// 	fmt.Println(result)
// }

// // retrieving the first document that matches the filter
// var result bson.M
// // check for errors in the finding
// if err = usersCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
// 	panic(err)
// }

// // display the document retrieved
// fmt.Println("displaying the first result from the search filter")
// fmt.Println(result)

// 	// retrieve all the documents in a collection
// cursor, err := usersCollection.Find(context.TODO(), bson.D{})
// // check for errors in the finding
// if err != nil {
//         panic(err)
// }

// // convert the cursor result to bson
// var results []bson.M
// // check for errors in the conversion
// if err = cursor.All(context.TODO(), &results); err != nil {
//         panic(err)
// }

// // display the documents retrieved
// fmt.Println("displaying all results in a collection")
// for _, result := range results {
//         fmt.Println(result)
// }

// // update a single document with a specified ObjectID using UpdateByID()
// // insert a new document to the collection
// user := bson.D{{"fullName", "User 5"}, {"age", 22}}
// insertResult, err := usersCollection.InsertOne(context.TODO(), user)
// if err != nil {
//         panic(err)
// }

// // create the update query for the client
// update := bson.D{
//         {"$set",
//                 bson.D{
//                         {"fullName", "User V"},
//                 },
//         },
//         {"$inc",
//                 bson.D{
//                         {"age", 1},
//                 },
//         },
// }

// // execute the UpdateByID() function with the filter and update query
// result, err := usersCollection.UpdateByID(context.TODO(), insertResult.InsertedID, update)
// // check for errors in the updating
// if err != nil {
//         panic(err)
// }
// // display the number of documents updated
// fmt.Println("Number of documents updated:", result.ModifiedCount)

// // update single and multiple documents with a specified filter using UpdateOne() and UpdateMany()
// // create a search filer
// filter := bson.D{
// 	{"$and",
// 			bson.A{
// 					bson.D{
// 							{"age", bson.D{{"$gt", 25}}},
// 					},
// 			},
// 	},
// }

// // create the update query
// update := bson.D{
// 	{"$set",
// 			bson.D{
// 					{"age", 40},
// 			},
// 	},
// }

// // execute the UpdateOne() function to update the first matching document
// result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
// // check for errors in the updating
// if err != nil {
// 	panic(err)
// }
// // display the number of documents updated
// fmt.Println("Number of documents updated:", result.ModifiedCount)

// // execute the UpdateMany() function to update all matching first document
// results, err := usersCollection.UpdateMany(context.TODO(), filter, update)
// // check for errors in the updating
// if err != nil {
// 	panic(err)
// }
// // display the number of documents updated
// fmt.Println("Number of documents updated:", results.ModifiedCount)

// // replace the fields of a single document with ReplaceOne()
// // create a search filer
// filter := bson.D{{"fullName", "User 1"}}

// // create the replacement data
// replacement := bson.D{
//         {"firstName", "John"},
//         {"lastName", "Doe"},
//         {"age", 30},
//         {"emailAddress", "johndoe@email.com"},
// }

// // execute the ReplaceOne() function to replace the fields
// result, err := usersCollection.ReplaceOne(context.TODO(), filter, replacement)
// // check for errors in the replacing
// if err != nil {
//         panic(err)
// }
// // display the number of documents updated
// fmt.Println("Number of documents updated:", result.ModifiedCount)

// // delete single and multiple documents with a specified filter using DeleteOne() and DeleteMany()
// // create a search filter
// filter := bson.D{
// 	{"$and",
// 			bson.A{
// 					bson.D{
// 							{"age", bson.D{{"$gt", 25}}},
// 					},
// 			},
// 	},
// }

// // delete the first document that match the filter
// result, err := usersCollection.DeleteOne(context.TODO(), filter)
// // check for errors in the deleting
// if err != nil {
// 	panic(err)
// }
// // display the number of documents deleted
// fmt.Println("deleting the first result from the search filter")
// fmt.Println("Number of documents deleted:", result.DeletedCount)

// // delete every document that match the filter
// results, err := usersCollection.DeleteMany(context.TODO(), filter)
// // check for errors in the deleting
// if err != nil {
// 	panic(err)
// }
// // display the number of documents deleted
// fmt.Println("deleting every result from the search filter")
// fmt.Println("Number of documents deleted:", results.DeletedCount)

//}
