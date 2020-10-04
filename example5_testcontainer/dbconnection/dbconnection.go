package dbconnection

import (
	"context"
	"fmt"
	"log"
	_ "net/http/pprof"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

const testTableName = "keyvalue"

type DBItem struct {
	Key   string
	Value string
}

type MyCollection struct {
	collection *mongo.Collection
}

//regular connection to a DB
func NewDBConnection(configName string) (*MyCollection, error) {
	dbHost, dbPort, dbName, tableName, err := readDbConfig(configName)
	if err != nil {
		log.Fatalf("Failed to read configuration file; %s", err)
	}
	return connectToDB(dbHost, dbPort, dbName, tableName)
}

//create, start and connect to mongo DB in a docker container
func NewTestDBConnection(t *testing.T) (*MyCollection, testcontainers.Container) {
	//create and start mongo DB container
	ctx := context.Background()
	mdb, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo:latest",
			ExposedPorts: []string{"27017/tcp"},
		},
		Started: true,
	})
	if err != nil {
		t.Error(err)
	}
	port, err := mdb.MappedPort(ctx, "27017")
	host, err := mdb.Host(ctx)
	if err != nil {
		t.Fatal("Container failed to start as expected", host, port)
	}
	t.Logf("MongoDB test container started on %s:%v", host, port)

	collection, err := connectToDB(host, port.Int(), "local", testTableName)
	if err != nil {
		t.Fatal("Failed to connect to DB and create new table", err)
	}
	return collection, mdb
}

func connectToDB(dbHost string, dbPort int, dbName string, tableName string) (*MyCollection, error) {
	dbURI := fmt.Sprintf("mongodb://%s:%d", dbHost, dbPort)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatalf("Failed connect to db; %s", err)
	}
	collection := client.Database(dbName).Collection(tableName)
	return &MyCollection{collection}, nil
}

func (coll *MyCollection) Insert(key string, value string) (interface{}, error) {
	res, err := coll.collection.InsertOne(Ctx, bson.M{"key": key, "value": value})
	if res == nil {
		return nil, fmt.Errorf("failed to insert [%s, %s] to db", key, value)
	}
	return res.InsertedID, err
}

func (coll *MyCollection) Delete(key string) (bool, error) {
	res, err := coll.collection.DeleteOne(Ctx, bson.M{"key": key})
	if res == nil || res.DeletedCount == 0 {
		return false, err
	}
	return true, err
}

type Result struct {
	ID    [12]byte
	Key   string
	Value string
}

func (coll *MyCollection) GetValue(key string) (interface{}, error) {
	result := Result{}
	err := coll.collection.FindOne(Ctx, bson.M{"key": key}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}

func readDbConfig(configName string) (dbHost string, dbPort int, dbName string, tableName string, err error) {
	viper.AddConfigPath("db/config")
	viper.SetConfigName(configName)
	viper.SetConfigType("properties")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Printf("Using config: %s\n", viper.ConfigFileUsed())

	dbHost = viper.GetString("db.Host")
	dbPort = viper.GetInt("db.Port")
	dbName = viper.GetString("db.Name")
	tableName = viper.GetString("db.tableName")
	return
}
