package linker

import (
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// LinkCollectionName is the name of the collections in the DB
	LinkCollectionName = "Links"
)

type Link struct {
	ID       string    `bson:"_id"`
	GroupsID [2]string `bson:"GroupsID"`
	Active   bool      `bson:"Active"`
	//TotalMessage is the number off message receiving by this link
	TotalMessage int    `bson:"TotalMessage"`
	CreatedAt    string `bson:"CreatedAt"`
	UpdatedAt    string `bson:"UpdatedAt"`
}

//CreateLink create a new Link
func CreateLink(GroupsID [2]string) (*Link, error) {

	for _, gID := range GroupsID {
		_, err := GetGroupByID(gID)
		if err != nil {
			return nil, err
		}
	}

	t, _ := time.Now().UTC().MarshalText()
	lnk := &Link{
		ID:           uniuri.NewLen(5),
		GroupsID:     GroupsID,
		Active:       true,
		TotalMessage: 0,
		CreatedAt:    string(t),
		UpdatedAt:    string(t),
	}

	_, err := DB.Collection(LinkCollectionName).InsertOne(*MongoCtx, lnk)
	if err != nil {
		return nil, err
	}

	return lnk, nil
}

//GetLinksByGroupID get all the links where a group is in
func GetLinksByGroupID(groupId string) ([]Link, error) {

	filter := bson.M{"GroupsID": groupId}
	cur, err := DB.Collection(LinkCollectionName).Find(*MongoCtx, filter)
	if err != nil {
		return nil, err
	}

	links := []Link{}
	err = cur.All(*MongoCtx, &links)
	if err != nil {
		return nil, err
	}

	return links, nil
}

//GetLinkByID get a link identify by the given id
func GetLinkByID(id string) (*Link, error) {

	lnk := &Link{}
	filter := bson.M{"_id": id}
	err := DB.Collection(LinkCollectionName).FindOne(*MongoCtx, filter).Decode(lnk)
	if err != nil {
		return nil, err
	}

	return lnk, nil
}

// IncrementMessage increment the number of message shared
func (lnk *Link) IncrementMessage() error {

	filter := bson.M{"_id": lnk.ID}
	updates := bson.M{"$inc": bson.M{"TotalMessage": 1}}
	result := DB.Collection(LinkCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	lnk.TotalMessage++
	return nil
}

//StopLink set the Active field to false
func (lnk *Link) StopLink() error {

	err := stopLink(lnk.ID)
	if err != nil {
		return err
	}
	lnk.Active = false
	return nil
}

func stopLink(id string) error {

	filter := bson.M{"_id": id}
	updates := bson.M{"$set": bson.M{"Active": false}}
	result := DB.Collection(LinkCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}

func stopLinks(groupId string) error {

	update := bson.M{"$set": bson.M{
		"Active": false,
	}}
	filter := bson.M{"GroupsID": groupId}
	_, err := DB.Collection(LinkCollectionName).UpdateMany(*MongoCtx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

//StartLink set the Active field to true
func (lnk *Link) StartLink() error {

	err := startLink(lnk.ID)
	if err != nil {
		return err
	}
	lnk.Active = true
	return nil
}

func startLink(id string) error {

	filter := bson.M{"_id": id}
	updates := bson.M{"$set": bson.M{"Active": true}}
	result := DB.Collection(LinkCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
