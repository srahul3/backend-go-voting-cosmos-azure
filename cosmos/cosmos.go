
package cosmos

import (
	"context"
//	"flag"
	"fmt"
	"log"
	"os"
//	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"github.com/srahul3/govoting/model"
)

var (
	database   string
	collection string
)

const (
	// environment variables
	mongoDBConnectionStringEnvVarName = "MONGODB_CONNECTION_STRING"
	mongoDBDatabaseEnvVarName         = "MONGODB_DATABASE"
	mongoDBCollectionEnvVarName       = "MONGODB_COLLECTION"

	// status
	statusPending   = "pending"
	statusCompleted = "completed"
	listAllCriteria = "all"
	statusAttribute = "status"

	// flags (commands)
	createFlag = "create"
	listFlag   = "list"
	updateFlag = "update"
	deleteFlag = "delete"

	// help text
	createHelp = "create a todo: enter description. e.g. todo -create \"get milk\""
	listHelp   = "list all, pending or completed todos. e.g. todo -list <criteria> (criteria can be all, pending or completed"
	updateHelp = "update a todo: enter todo ID and new status e.g. todo -update <id>,<new status> e.g. todo -update 1,completed"
	deleteHelp = "delete a todo: enter todo ID e.g. todo -delete 42"
)

// connects to MongoDB
func connect() *mongo.Client {
	mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		mongoDBConnectionString = "mongodb://srahul3-voting-db:grFE2DyiYrxzavKkTA2x5KSOrTUhPP2g7ldKGCUljfJ1Kse9NUGpst4Ada0VKwI3VP3IsZakDSNfxSMNezgtTQ==@srahul3-voting-db.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&replicaSet=globaldb&maxIdleTimeMS=120000&appName=@srahul3-voting-db@"
		//log.Fatal("missing environment variable: ", mongoDBConnectionStringEnvVarName)
	}

	database = os.Getenv(mongoDBDatabaseEnvVarName)
	if database == "" {
		// log.Fatal("missing environment variable: ", mongoDBDatabaseEnvVarName)
		database = "voting"
	}

	collection = os.Getenv(mongoDBCollectionEnvVarName)
	if collection == "" {
		// log.Fatal("missing environment variable: ", mongoDBCollectionEnvVarName)
		collection = "voting"
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoDBConnectionString).SetDirect(true)
	c, err := mongo.NewClient(clientOptions)

	err = c.Connect(ctx)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}
	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	
	
	
	return c
}

// creates a todo
func (voteCandidate VoteCandiate) Create() {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	r, err := todoCollection.InsertOne(ctx, voteCandidate)
	if err != nil {
		log.Fatalf("failed to add voting %v", err)
	}
	fmt.Println("added todo", r.InsertedID)
}

// lists todos
func List(status string) {

	var filter interface{}
	switch status {
	filter = bson.D{}

	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	rs, err := todoCollection.Find(ctx, filter)
	if err != nil {
		log.Fatalf("failed to list todo(s) %v", err)
	}
	var todos []VoteCandiate
	err = rs.All(ctx, &todos)
	if err != nil {
		log.Fatalf("failed to list todo(s) %v", err)
	}
	if len(todos) == 0 {
		fmt.Println("no todos found")
		return
	}

	todoTable := [][]string{}

	for _, todo := range todos {		
		todoTable = append(todoTable, []string{todo.ID, todo.Name, string(todo.Votes) })
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Votes"})

	for _, v := range todoTable {
		table.Append(v)
	}
	table.Render()
}

// updates a todo
func update(todoid, newStatus string) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	oid, err := primitive.ObjectIDFromHex(todoid)
	if err != nil {
		log.Fatalf("failed to update todo %v", err)
	}
	filter := bson.D{{"_id", oid}}
	update := bson.D{{"$set", bson.D{{statusAttribute, newStatus}}}}
	_, err = todoCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("failed to update todo %v", err)
	}
}

// deletes a todo
func delete(todoid string) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	oid, err := primitive.ObjectIDFromHex(todoid)
	if err != nil {
		log.Fatalf("invalid todo ID %v", err)
	}
	filter := bson.D{{"_id", oid}}
	_, err = todoCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("failed to delete todo %v", err)
	}
}

// Todo represents a todo
type VoteCandiate struct {
    ID      string  `bson:"id"`
    Name    string  `bson:"name"`
    LogoUrl string  `bson:"logo_url"`
    Votes   int     `bson:"votes"`
}