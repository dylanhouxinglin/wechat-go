package user

import (
	"context"
	"fmt"
	"github.com/siddontang/go/bson"
	"strconv"
	"wechat-go/preload"
	"github.com/bwmarrin/snowflake"
)

type User struct {
	Uid			string		`json:"uid"`
	NickName	string		`json:"nickname"`
}

func CheckUser(uname string) (interface{}) {
	db, err := preload.GetDBIns()
	if err != nil {
		panic(err)
		return nil
	}
	collection := db.Mongo.Collection("user")
	var newUser User
	ctx := context.TODO()
	filter := bson.M{"nickname":uname}
	err = collection.FindOne(ctx, filter).Decode(&newUser)
	// 没找到，新增
	if err != nil {
		// https://github.com/bwmarrin/snowflake
		// snowflake: A very simple Twitter snowflake generator
		node, err := snowflake.NewNode(1)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println("No Record Matched...Insert New Record...")
		newUser = User{
			NickName: uname,
			Uid: strconv.FormatInt(int64(node.Generate()), 10),
		}
		_, err = collection.InsertOne(ctx, newUser)
		if err != nil {
			panic(err)
			return nil
		}
		return newUser.Uid
	}
	// 找到了，直接返回
	return newUser.Uid
}
