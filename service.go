package linker

import (
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// StatusOnline
	StatusOnline = "online"
	// StatusOffline
	StatusOffline = "offline"
	// ServiceCollectionName
	ServiceCollectionName = "Services"
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
func SetService(name, url, authKey, status string) (*Service, error) {

	theService := &Service{}

	// we check service existence, to update or create a new service
	isServiceAlreadyExist := true
	filter := bson.M{"Name": name}
	err := DB.Collection(ServiceCollectionName).FindOne(*MongoCtx, filter).Decode(theService)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			isServiceAlreadyExist = false
		} else {
			return nil, err
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
		theService.Url = url
		theService.AuthKey = authKey
		theService.Status = status

		_, err = DB.Collection(ServiceCollectionName).UpdateOne(*MongoCtx, filter, updates)
		if err != nil {
			return nil, err
		}

	} else {

		myService := Service{
			ID:      uuid.Must(uuid.NewRandom()).String(),
			Name:    name,
			Url:     url,
			AuthKey: authKey,
			Status:  status,
		}

		_, err := DB.Collection(ServiceCollectionName).InsertOne(*MongoCtx, myService)
		if err != nil {
			return nil, err
		}
		theService = &myService
	}

	return theService, nil
}

// GetService
func GetService(serviceName string) (*Service, error) {

	service := &Service{}
	filter := bson.M{"Name": serviceName}
	err := DB.Collection(ServiceCollectionName).FindOne(*MongoCtx, filter).Decode(service)
	if err != nil {
		return nil, err
	}

	return service, nil
}
