package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type Mongodb struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
	Context    context.Context
}

func (mdb *Mongodb) Connect(uri string) error {
	var err error
	mdb.Client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return err
	}
	mdb.Context, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = mdb.Client.Connect(mdb.Context)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = mdb.Client.Ping(mdb.Context, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (mdb *Mongodb) InitDatabase() {
	mdb.initHomeDB()
	mdb.initEnvironmentCollection()
}

func (mdb *Mongodb) InsertElement(i interface{}) error {
	_, err := mdb.Collection.InsertOne(mdb.Context, i)
	return err
}

func (mdb *Mongodb) GetLatest() (interface{}, error) {
	var i interface{}
	err := mdb.Collection.FindOne(context.TODO(), bson.D{}).Decode(&i)
	return i, err
}

func (mdb *Mongodb) initHomeDB() {
	mdb.Database = mdb.Client.Database("home")
}

func (mdb *Mongodb) initEnvironmentCollection() {
	mdb.Collection = mdb.Database.Collection("environment")
}
