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

	// help text
	createHelp = "create a todo: enter description. e.g. todo -create \"get milk\""
	listHelp   = "list all, pending or completed todos. e.g. todo -list <criteria> (criteria can be all, pending or completed"
	updateHelp = "update a todo: enter todo ID and new status e.g. todo -update <id>,<new status> e.g. todo -update 1,completed"
	deleteHelp = "delete a todo: enter todo ID e.g. todo -delete 42"
)

// connects to MongoDB
func connect() *(mongo.Client) {
	mongoDBConnectionString := os.Getenv(mongoDBConnectionStringEnvVarName)
	if mongoDBConnectionString == "" {
		log.Fatal("missing environment variable: ", mongoDBConnectionStringEnvVarName)
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

// creates a a voting vandidate
func (voteCandidate VoteCandiate) CreateIfDoesntExist() {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	filter := bson.D{{"id", voteCandidate.ID}}
	count, err := todoCollection.CountDocuments(ctx, filter)

	if err != nil {
		log.Fatalf("failed to count %v", err)
	} else if count <= 0 {
		r, err := todoCollection.InsertOne(ctx, voteCandidate)
		if err != nil {
			log.Fatalf("failed to add voting %v", err)
		}
		fmt.Println("added todo", r.InsertedID)
	} else {
		log.Println("Document already exisits")
	}

}

// creates a a voting vandidate
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

// updates a todo
func (voteCandidate VoteCandiate) Update() {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	oid, err := primitive.ObjectIDFromHex(voteCandidate.ID)
	if err != nil {
		log.Fatalf("failed to update todo %v", err)
	}
	filter := bson.D{{"_id", oid}}
	update := bson.D{{"$inc", bson.D{{"vote", voteCandidate.Votes}, {"name", voteCandidate.Name}, {"logo_url", voteCandidate.LogoUrl}}}}
	_, err = todoCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("failed to update todo %v", err)
	}
}

func VoteUp(id string) {
	c := connect()
	ctx := context.Background()
	defer c.Disconnect(ctx)

	todoCollection := c.Database(database).Collection(collection)
	filter := bson.D{{"id", id}}
	update := bson.D{{"$inc", bson.D{{"votes", 1}}}}
	_, err := todoCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("failed to update todo %v", err)
	}
}

// lists todos
func List() []VoteCandiate {

	var filter interface{}
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

		return []VoteCandiate{}
	}

	todoTable := [][]string{}

	for _, todo := range todos {
		todoTable = append(todoTable, []string{todo.ID, todo.Name, string(todo.Votes)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Votes"})

	for _, v := range todoTable {
		table.Append(v)
	}
	table.Render()

	return todos
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
	ID      string `bson:"id"`
	Name    string `bson:"name"`
	LogoUrl string `bson:"logo_url"`
	Votes   int    `bson:"votes"`
}
