package group

import (
	"github.com/DipandaAser/linker/config"
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// CollectionName
	CollectionName = "Groups"
)

// Group
type Group struct {

	//ID represent the group ID according to the Service,
	//in telegram the ID should be the chat_id
	ID string `bson:"_id"`

	//Service represent the service.Service{} where the group is stored
	//Example: whatsapp, telegram,
	Service string `bson:"Service"`

	//ShortCode is an internal Linker identifier of the group
	//Is used when someone want to create a new link
	//instead of using the ID who is bit long on some platform, they use the ShortCode
	ShortCode string `bson:"ShortCode"`

	//TotalMessage is the number off message receiving by the bot in this group
	TotalMessage int `bson:"TotalMessage"`

	//CreateGroup is the date when this group join linker
	CreatedAt string `bson:"CreatedAt"`
}

// CreateGroup
func CreateGroup(id, service string) (*Group, error) {

	// We generate a ShortCode an make sure that we dont have duplicate ShortCode
	code := uniuri.NewLen(5)
	_, err := GetGroupByShortCode(code)
	for err == nil {
		code = uniuri.NewLen(5)
		_, err = GetGroupByShortCode(code)
	}

	t, _ := time.Now().UTC().MarshalText()
	theGroup := Group{
		ID:           id,
		Service:      service,
		ShortCode:    code,
		TotalMessage: 0,
		CreatedAt:    string(t),
	}

	_, err = config.DB.Collection(CollectionName).InsertOne(config.MongoCtx, theGroup)
	if err != nil {
		return nil, err
	}

	return &theGroup, nil
}

// GetGroup
func GetGroup(id string) (*Group, error) {

	theGroup := &Group{}
	filter := bson.M{"_id": id}
	err := config.DB.Collection(CollectionName).FindOne(config.MongoCtx, filter).Decode(theGroup)
	if err != nil {
		return nil, err
	}

	return theGroup, nil
}

// GetGroupByShortCode
func GetGroupByShortCode(shortCode string) (*Group, error) {

	theGroup := &Group{}
	filter := bson.M{"ShortCode": shortCode}
	err := config.DB.Collection(CollectionName).FindOne(config.MongoCtx, filter).Decode(theGroup)
	if err != nil {
		return nil, err
	}

	return theGroup, nil
}

// IncrementMessage
func IncrementMessage(id string) error {

	g, err := GetGroup(id)
	if err != nil {
		return err
	}

	return g.IncrementMessage()
}

// IncrementMessage
func (g *Group) IncrementMessage() error {

	filter := bson.M{"_id": g.ID}
	updates := bson.M{"$inc": bson.M{"TotalMessage": 1}}
	result := config.DB.Collection(CollectionName).FindOneAndUpdate(config.MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
