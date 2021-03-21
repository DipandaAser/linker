package main

import (
	"fmt"
	"github.com/DipandaAser/linker/config"
	"github.com/DipandaAser/linker/message"
	"github.com/DipandaAser/linker/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"

	"os"
	"time"
)

func main() {

	_ = os.Setenv("APIKEY", "test")
	_ = os.Setenv("DB_NAME", "linker")

	_ = os.Setenv("PORT", "8080")
	_ = os.Setenv("WEB_URL", "")

	config.ProjectConfig.ServiceName = "whatsapp"
	config.ProjectConfig.ProjectName = "Linker Whatsapp"
	config.ProjectConfig.AuthKey = os.Getenv("APIKEY")
	config.ProjectConfig.DBName = os.Getenv("DB_NAME")
	config.ProjectConfig.MongodbURI = os.Getenv("MONGO_URI")
	config.ProjectConfig.HTTPPort = os.Getenv("PORT")
	config.ProjectConfig.WebUrl = os.Getenv("WEB_URL")

	// ─── MONGO ──────────────────────────────────────────────────────────────────────
	err := MongoConnect()
	if err != nil {
		log.Fatal("Can't setup mongodb")
	}

	// ─── WE REFRESH THE MONGO CONNECTION EACH 10MINS ──────────────────────────────────────
	ticker := time.NewTicker(time.Minute * 10)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			go MongoReconnectCheck()
		}
	}()

	err = service.SetService(config.ProjectConfig.ServiceName, config.ProjectConfig.WebUrl, config.ProjectConfig.AuthKey, service.StatusOnline)
	if err != nil {
		log.Fatal("Can't set service")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, config.ProjectConfig.ProjectName)
	})

	router.POST("/linker/:authKey/message", func(ctx *gin.Context) {

		autKey := ctx.Param("autKey")
		if autKey != config.ProjectConfig.AuthKey {
			ctx.String(http.StatusUnauthorized, "")
		}

		messageInfo := message.Info{}
		err = ctx.BindJSON(&messageInfo)
		switch messageInfo.MessageData.(type) {
		case message.TextMessage:
			break
		case message.AudioMessage:
			break
		case message.ImageMessage:
			break
		case message.VideoMessage:
			break
		case message.DocumentMessage:
			break
		}
		ctx.String(http.StatusAccepted, config.ProjectConfig.ProjectName)
	})

	log.Printf("%s Start successfully", config.ProjectConfig.ProjectName)
	_ = router.Run(fmt.Sprintf(":%s", config.ProjectConfig.HTTPPort))
}

// MongoConnect connects to mongoDB
func MongoConnect() error {

	clientOptions := options.Client().ApplyURI(config.ProjectConfig.MongodbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(config.MongoCtx, clientOptions)
	if err != nil {
		return err
	}

	// We make sure we have been connected
	err = client.Ping(config.MongoCtx, readpref.Primary())
	if err != nil {
		return err
	}

	config.DB = client.Database(config.ProjectConfig.MongodbURI)

	return nil
}

// MongoReconnectCheck reconnects to MongoDB
func MongoReconnectCheck() {

	// We make sure we are still connected
	err := config.DB.Client().Ping(config.MongoCtx, readpref.Primary())
	if err != nil {
		// We reconnect
		_ = MongoConnect()
	}
}
