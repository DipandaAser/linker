package link

import (
	"github.com/DipandaAser/linker/config"
	"github.com/DipandaAser/linker/group"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// CollectionName
	CollectionName = "Links"
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

//CreateLink
func CreateLink(GroupsID [2]string) (*Link, error) {

	for _, gID := range GroupsID {
		_, err := group.GetGroupByID(gID)
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

	_, err := config.DB.Collection(CollectionName).InsertOne(config.MongoCtx, lnk)
	if err != nil {
		return nil, err
	}

	return lnk, nil
}

func GetLinksByGroupID(id string) ([]Link, error) {

	filter := bson.M{"GroupsID": id}
	cur, err := config.DB.Collection(CollectionName).Find(config.MongoCtx, filter)
	if err != nil {
		return nil, err
	}

	links := []Link{}
	err = cur.All(config.MongoCtx, &links)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// IncrementMessage
func (lnk *Link) IncrementMessage() error {

	filter := bson.M{"_id": lnk.ID}
	updates := bson.M{"$inc": bson.M{"TotalMessage": 1}}
	result := config.DB.Collection(CollectionName).FindOneAndUpdate(config.MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
