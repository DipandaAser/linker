package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProjectSettings type allow reading config.json file
type ProjectSettings struct {
	ServiceName string
	MongodbURI  string
	DBName      string
	HTTPPort    string
	WebUrl      string
	ProjectName string
	AuthKey     string
}

// ProjectConfig holds info of project settings
var ProjectConfig = &ProjectSettings{}

// MongoCtx is the mongo context
var MongoCtx = context.TODO()

// LocalDB is the mongo db
var DB *mongo.Database
