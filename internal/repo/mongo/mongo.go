package DatabaseMongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"bodyplate.com/config"
	DatabaseAbtract "bodyplate.com/internal/repo/type"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	CollectionName string
	Collection     *mongo.Collection
	Schema         interface{}
}

func (m *MongoRepository) Create(obj interface{}) (interface{}, error) {
	if !IsInstanceOfStruct(obj, m.Schema) {
		fmt.Printf("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
		return nil, errors.New("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
	}
	result, err := m.Collection.InsertOne(context.TODO(), obj)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
	// return obj, nil
}

func (m *MongoRepository) Find(con interface{}) ([]interface{}, error) {
	if !IsInstanceOfStruct(con, m.Schema) {
		fmt.Printf("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
		return nil, errors.New("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
	}
	var results []interface{}
	cursor, err := m.Collection.Find(context.TODO(), con)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &results)
	return results, nil
}

func (m *MongoRepository) FindOne(con interface{}) (interface{}, error) {
	if !IsInstanceOfStruct(con, m.Schema) {
		fmt.Printf("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
		return nil, errors.New("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
	}
	var result interface{}
	err := m.Collection.FindOne(context.TODO(), con).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MongoRepository) Update(con interface{}, obj interface{}) error {
	if !IsInstanceOfStruct(obj, m.Schema) {
		fmt.Printf("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
		return errors.New("CONDITION_NOT_IS_INSTANCE_OF_THIS_SCHEMA_REPO")
	}
	_, err := m.Collection.UpdateOne(context.TODO(), con, obj)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepository) Delete(con interface{}) error {
	_, err := m.Collection.DeleteOne(context.TODO(), con)
	if err != nil {
		return err
	}
	return nil
}

func IsInstanceOfStruct(value interface{}, sampleStruct interface{}) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(sampleStruct)
}

func (m *MongoRepository) CreateIndex(option []mongo.IndexModel) ([]string, error) {
	result, err := m.Collection.Indexes().CreateMany(context.TODO(), option)
	return result, err
}

type MongoDB struct {
	conection   *mongo.Client
	ResHttpRepo DatabaseAbtract.Repository
	DomRepo     DatabaseAbtract.Repository
}

var MongoDatabase MongoDB = MongoDB{}

func (m *MongoDB) Init() {
	clientOptions := options.Client().ApplyURI(config.GetConfigAvaiable().Uri_mongo)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	m.conection = client

	// * ResRepository * //
	resHTTPRepo := &MongoRepository{
		CollectionName: "res_http",
		Collection:     client.Database(config.GetConfigAvaiable().DB_Mongo).Collection("res_http"),
		Schema:         ResponseHttpSchema{},
	}
	resHTTPRepo.CreateIndex([]mongo.IndexModel{
		{
			Keys: bson.D{
				{"created_at", 1},
			},
			Options: options.Index(),
		},
	})
	m.ResHttpRepo = resHTTPRepo
	// * ResRepository * //

	// * DomRepository * //
	domRepo := &MongoRepository{
		CollectionName: "res_dom",
		Collection:     client.Database(config.GetConfigAvaiable().DB_Mongo).Collection("res_dom"),
		Schema:         DomSchema{},
	}
	domRepo.CreateIndex([]mongo.IndexModel{
		{
			Keys: bson.D{
				{"created_at", 1},
			},
			Options: options.Index(),
		},
	})
	m.DomRepo = domRepo
	// * DomRepository * //
	fmt.Println("Connected to MongoDB!")
}
