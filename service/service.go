package service

import (
	"errors"
	"github.com/DipandaAser/linker/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// StatusOnline
	StatusOnline = "online"
	// StatusOffline
	StatusOffline = "offline"
	// CollectionName
	CollectionName = "Services"
)

var (
	// ErrServiceOffline
	ErrServiceOffline = errors.New("service is offline")
)

// Service represent a linker service, Ex: linker-whatsapp, linker-telegram
type Service struct {
	ID   string `json:"_id" bson:"_id"`
	Name string `json:"Name" bson:"Name"` // whatsapp, telegram....
	//Url is the web url where is hosted the current service, it is use by other service to locate where they gonna transfer message
	Url string `json:"Url" bson:"Url"`
	//AuthKey is the key using by other service to transfer message to the current service
	AuthKey string `json:"AuthKey" bson:"AuthKey"`
	Status  string `json:"Status" bson:"Status"` // online, offline
}

// SetService
func SetService(name, url, authKey, status string) error {

	// we check service existence, to update or create a new service
	isServiceAlreadyExist := true
	filter := bson.M{"Name": name}
	result := config.DB.Collection(CollectionName).FindOne(config.MongoCtx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			isServiceAlreadyExist = false
		} else {
			return err
		}
	}

	if isServiceAlreadyExist {

		updates := bson.M{
			"$set": bson.M{
				"Url":     url,
				"AuthKey": authKey,
				"Status":  status,
			},
		}

		_, err := config.DB.Collection(CollectionName).UpdateOne(config.MongoCtx, filter, updates)
		if err != nil {
			return err
		}

	} else {

		myService := Service{
			ID:      uuid.Must(uuid.NewRandom()).String(),
			Name:    name,
			Url:     url,
			AuthKey: authKey,
			Status:  status,
		}

		_, err := config.DB.Collection(CollectionName).InsertOne(config.MongoCtx, myService)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetService
func GetService(serviceName string) (*Service, error) {

	service := &Service{}
	filter := bson.M{"Name": serviceName}
	err := config.DB.Collection(CollectionName).FindOne(config.MongoCtx, filter).Decode(service)
	if err != nil {
		return nil, err
	}

	return service, nil
}
