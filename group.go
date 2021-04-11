package linker

import (
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// GroupCollectionName is the name of the collections in the DB
	GroupCollectionName = "Groups"
)

// Group represent a group or a channel
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

// CreateGroup create a ne group
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

	_, err = DB.Collection(GroupCollectionName).InsertOne(*MongoCtx, theGroup)
	if err != nil {
		return nil, err
	}

	return &theGroup, nil
}

// GetGroupByID get a group by his id
func GetGroupByID(id string) (*Group, error) {

	theGroup := &Group{}
	filter := bson.M{"_id": id}
	err := DB.Collection(GroupCollectionName).FindOne(*MongoCtx, filter).Decode(theGroup)
	if err != nil {
		return nil, err
	}

	return theGroup, nil
}

// GetGroupByShortCode get a group by his shortcode
func GetGroupByShortCode(shortCode string) (*Group, error) {

	theGroup := &Group{}
	filter := bson.M{"ShortCode": shortCode}
	err := DB.Collection(GroupCollectionName).FindOne(*MongoCtx, filter).Decode(theGroup)
	if err != nil {
		return nil, err
	}

	return theGroup, nil
}

// IncrementMessage increment the number of message received
func (g *Group) IncrementMessage() error {

	filter := bson.M{"_id": g.ID}
	updates := bson.M{"$inc": bson.M{"TotalMessage": 1}}
	result := DB.Collection(GroupCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
