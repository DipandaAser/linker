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
var MongoCtx *context.Context

// DB is the mongo db
var DB *mongo.Database

// Init init the linker package
func Init(config *ProjectSettings, ctx *context.Context, db *mongo.Database) {
	Config = config
	MongoCtx = ctx
	DB = db
}
