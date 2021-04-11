package linker

import (
	"github.com/google/uuid"
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
		ID:           uuid.Must(uuid.NewRandom()).String(),
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

//GetLinksByGroupID get all the links where a group is linked by his given id
func GetLinksByGroupID(id string) ([]Link, error) {

	filter := bson.M{"GroupsID": id}
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

// IncrementMessage increment the number of message shared
func (lnk *Link) IncrementMessage() error {

	filter := bson.M{"_id": lnk.ID}
	updates := bson.M{"$inc": bson.M{"TotalMessage": 1}}
	result := DB.Collection(LinkCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
