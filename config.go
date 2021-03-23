package linker

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

var HeaderAuthKey = "authKey"

// Config holds info of project settings
var Config = &ProjectSettings{}

// MongoCtx is the mongo context
var MongoCtx = context.TODO()

// LocalDB is the mongo db
var DB *mongo.Database
