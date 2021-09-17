package preload

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"wechat-go/db"
	"wechat-go/defines"
	"github.com/spf13/viper"
)

var database *db.MongoDB

func Init()  {
	initConf()
	initDB()
}

func initConf() {
	confFile := defines.ProjectPath + "/conf/config.json"
	viper.SetConfigFile(confFile)
	if err := viper.ReadInConfig(); err != nil {
		// add log
		panic(err)
	}
}

func initDB() {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration(`db.timeout`)*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI(viper.GetString(`db.uri`))
	opt.SetMaxPoolSize(viper.GetUint64(`db.maxPoolSize`))

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	database = &db.MongoDB{}
	database.DB = client.Database(viper.GetString(`db.database`))
}

func GetDBIns() (*db.MongoDB, error) {
	if database == nil && database.DB == nil {
		return nil, errors.New("Get DB Ins Error")
	}
	return database, nil
}