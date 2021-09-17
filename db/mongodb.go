package db

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	DB		*mongo.Database
}
