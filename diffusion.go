package linker

import (
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	// DiffusionCollectionName is the name of the collections in the DB
	DiffusionCollectionName = "Diffusions"
)

type Diffusion struct {
	ID          string `bson:"_id"`
	Broadcaster string `bson:"Broadcaster"`
	Receiver    string `bson:"Receiver"`
	Active      bool   `bson:"Active"`
	//TotalMessage is the number off message receiving by this link
	TotalMessage int    `bson:"TotalMessage"`
	CreatedAt    string `bson:"CreatedAt"`
	UpdatedAt    string `bson:"UpdatedAt"`
}

//CreateDiffusion create a new diffusion
func CreateDiffusion(broadcaster, receiver string) (*Diffusion, error) {

	_, err := GetGroupByID(broadcaster)
	if err != nil {
		return nil, err
	}
	_, err = GetGroupByID(receiver)
	if err != nil {
		return nil, err
	}

	t, _ := time.Now().UTC().MarshalText()
	lnk := &Diffusion{
		ID:           uniuri.NewLen(5),
		Broadcaster:  broadcaster,
		Receiver:     receiver,
		Active:       true,
		TotalMessage: 0,
		CreatedAt:    string(t),
		UpdatedAt:    string(t),
	}

	_, err = DB.Collection(DiffusionCollectionName).InsertOne(*MongoCtx, lnk)
	if err != nil {
		return nil, err
	}

	return lnk, nil
}

//GetDiffusionsByBroadcaster get all the diffusion where the broadcaster is the group given id
func GetDiffusionsByBroadcaster(id string) ([]Diffusion, error) {

	filter := bson.M{"Broadcaster": id}
	cur, err := DB.Collection(DiffusionCollectionName).Find(*MongoCtx, filter)
	if err != nil {
		return nil, err
	}

	links := []Diffusion{}
	err = cur.All(*MongoCtx, &links)
	if err != nil {
		return nil, err
	}

	return links, nil
}

//GetDiffusionsByReceiver get all the diffusion where the receiver is the group given id
func GetDiffusionsByReceiver(id string) ([]Diffusion, error) {

	filter := bson.M{"Receiver": id}
	cur, err := DB.Collection(DiffusionCollectionName).Find(*MongoCtx, filter)
	if err != nil {
		return nil, err
	}

	links := []Diffusion{}
	err = cur.All(*MongoCtx, &links)
	if err != nil {
		return nil, err
	}

	return links, nil
}

//GetDiffusionByBroadcasterAndReceiver get a diffusion where the broadcaster and receiver is the group given id
func GetDiffusionByBroadcasterAndReceiver(broadcaster, receiver string) (*Diffusion, error) {

	diff := &Diffusion{}
	filter := bson.M{"Broadcaster": broadcaster, "Receiver": receiver}
	err := DB.Collection(DiffusionCollectionName).FindOne(*MongoCtx, filter).Decode(diff)
	if err != nil {
		return nil, err
	}

	return diff, nil
}

// IncrementMessage increment the number of message diffused
func (dif *Diffusion) IncrementMessage() error {

	filter := bson.M{"_id": dif.ID}
	updates := bson.M{"$inc": bson.M{"TotalMessage": 1}}
	result := DB.Collection(DiffusionCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}
	return nil
}
