package mongodb

import (
	"context"
	"fmt"
	"log"

	"bodyplate.com/config"
	abtractRepository "bodyplate.com/internal/database/type"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	collectionName string
	collection *mongo.Collection
}


func (m *MongoRepository) Create(obj interface{}) (interface{}, error) {
	result, err := m.collection.InsertOne(context.TODO(), obj)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (m *MongoRepository) Find(con interface{}) ([]interface{}, error) {
	var results []interface{}
	cursor, err := m.collection.Find(context.TODO(), con)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), &results)
	return results, nil
}

func (m *MongoRepository) FindOne(con interface{}) (interface{}, error) {
	var result interface{}
	err := m.collection.FindOne(context.TODO(), con).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MongoRepository) Update(con interface{},obj interface{}) (error) {
	_, err := m.collection.UpdateOne(context.TODO(), con, obj)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepository) Delete(con interface{}) (error) {
	_, err := m.collection.DeleteOne(context.TODO(), con)
	if err != nil {
		return err
	}
	return nil
}

type MongoDB struct {
	conection *mongo.Client
	ResHttpRepo abtractRepository.Repository
}
var MongoDatabase MongoDB = MongoDB{}

func (m *MongoDB) Init() {
	clientOptions := options.Client().ApplyURI(config.GetConfigAvaiable().Uri_mongo)
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    } 	
	fmt.Println("Connected to MongoDB!")
	m.conection = client
	m.ResHttpRepo = &MongoRepository{
		collectionName: "res_http",
		collection: client.Database(config.GetConfigAvaiable().DB_Mongo).Collection("res_http"),
	}

}