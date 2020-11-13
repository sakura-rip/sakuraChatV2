package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User b
type User struct {
	ID       string             `bson:"_id"`
	Contacts []Contact          `bson:"contacts"`
	BIDs     []string           `bson:"BIDs"`
	MapTest  map[string]Contact `bson:"MapTest"`
}

// Contact a
type Contact struct {
	Name string `bson:"name"`
	UUID string `bson:"uuid"`
}

var col *mongo.Collection

func insert() {
	doc := User{
		ID:       "ID1",
		Contacts: []Contact{{Name: "name1", UUID: "uuid1"}, {Name: "name2", UUID: "uuid2"}},
		BIDs:     []string{"aaa", "aaa", "bbbb"},
	}
	col.InsertOne(context.Background(), doc)
}

func maptest() {
	doc := User{
		ID:       "ID2",
		Contacts: []Contact{{Name: "name2", UUID: "uuid2"}, {Name: "name3", UUID: "uuid3"}},
		BIDs:     []string{"aaa", "aaa", "bbbb"},
		MapTest:  map[string]Contact{"testuuid": Contact{Name: "name", UUID: "testuuid"}},
	}
	col.InsertOne(context.Background(), doc)
}
func find() {
	rs := col.FindOne(
		context.Background(),
		bson.D{{"_id", "ID1"}},
		options.FindOne().SetProjection(bson.D{{"contacts.name", 1}, {"contacts.uuid", 1}}),
	)
	var usr *User
	if err := rs.Decode(&usr); err != nil {
		panic(err)
	}
	fmt.Println(usr)
}

func findArray() {
	fmt.Println("find arry")
	rs := col.FindOne(
		context.Background(),
		bson.D{{"_id", "ID1"}},
		options.FindOne().SetProjection(bson.M{"contacts": bson.M{"$elemMatch": bson.M{"uuid": "dfasdfa"}}}),
	)
	var usr *User
	if err := rs.Decode(&usr); err != nil {
		panic(err)
	}
	for _, con := range usr.Contacts {
		fmt.Println(con.Name)
	}
	fmt.Println(usr)
}

func pushArray() {
	rs, err := col.UpdateOne(context.Background(), bson.M{"_id": "ID1"}, bson.M{"$push": bson.M{"BIDs": "aaaafafsdaa"}})
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
func pushArrays() {
	rs, err := col.UpdateOne(context.Background(), bson.M{"_id": "ID2"}, bson.M{"$set": bson.M{"MapTest.tefasdfad": Contact{Name: "asdas", UUID: "fdsfa"}}})
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
func findMap(uuid string) {
	rs := col.FindOne(
		context.Background(),
		bson.D{{"_id", "ID2"}},
		options.FindOne().SetProjection(bson.D{{"MapTest", 1}}),
	)
	var usr *User
	if err := rs.Decode(&usr); err != nil {
		panic(err)
	}
	fmt.Println(usr)
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:password123@localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err = client.Connect(context.Background()); err != nil {
		panic(err)
	}
	col = client.Database("sakura").Collection("sakurachatTest")
	// findArray()
	// insert()
	// find()
	// pushArray()
	// maptest()
	// maptest()
	pushArrays()
	// findMap("testuuid")
}
