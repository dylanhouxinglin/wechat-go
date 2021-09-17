package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"wechat-go/preload"
)

func TestMongoDB(t *testing.T)  {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to MongoDB")
	db := client.Database("admin")
	fmt.Println(db.Name())
	//err = client.Disconnect(context.TODO())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}

func TestMongoDBPool(t *testing.T)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	o := options.Client().ApplyURI("mongodb://localhost:27017")
	o.SetMaxPoolSize(3)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := client.Database("admin")
	fmt.Println(db.Name())
	//return db
}

type stu struct {
	name	string
}

func TestDBInit(t *testing.T)  {
	preload.Init()
	db, err := preload.GetDBIns()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(db.DB.Name())
	collection := db.DB.Collection("testColl")
	s := stu{"hxl"}
	res, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(res)
}