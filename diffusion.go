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

//GetDiffusionsByGroupID get all the diffusion where the Broadcaster or the Receiver is the group given id
func GetDiffusionsByGroupID(id string) ([]Diffusion, error) {

	filter := bson.M{"$or": bson.A{
		bson.M{"Broadcaster": id},
		bson.M{"Receiver": id},
	}}
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

//GetDiffusionsByBroadcasterAndReceiver get all the diffusion where the broadcaster and receiver are the group givens id
func GetDiffusionsByBroadcasterAndReceiver(broadcasterID string, receiverID string) (*Diffusion, error) {

	filter := bson.M{"$or": bson.A{
		bson.M{"Broadcaster": broadcasterID},
		bson.M{"Receiver": receiverID},
	}}

	diff := &Diffusion{}
	err := DB.Collection(DiffusionCollectionName).FindOne(*MongoCtx, filter).Decode(diff)
	if err != nil {
		return nil, err
	}

	return diff, nil
}

//GetDiffusionById get a diffusion identify by the  given id
func GetDiffusionById(id string) (*Diffusion, error) {

	diff := &Diffusion{}
	filter := bson.M{"_id": id}
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

//StopDiffusion set the Active field to false
func (dif *Diffusion) StopDiffusion() error {

	err := stopDiffusion(dif.ID)
	if err != nil {
		return err
	}
	dif.Active = false
	return nil
}

//StartDiffusion set the Active field to true
func (dif *Diffusion) StartDiffusion() error {

	err := startDiffusion(dif.ID)
	if err != nil {
		return err
	}
	dif.Active = true
	return nil
}

func startDiffusion(id string) error {

	filter := bson.M{"_id": id}
	updates := bson.M{"$set": bson.M{"Active": true}}
	result := DB.Collection(DiffusionCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func stopDiffusion(id string) error {

	filter := bson.M{"_id": id}
	updates := bson.M{"$set": bson.M{"Active": false}}
	result := DB.Collection(DiffusionCollectionName).FindOneAndUpdate(*MongoCtx, filter, updates)
	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func stopDiffusions(groupId string) error {

	update := bson.M{"$set": bson.M{
		"Active": false,
	}}

	diffusionFilter := bson.M{"$or": bson.A{
		bson.M{"Broadcaster": groupId},
		bson.M{"Receiver": groupId},
	}}
	_, err := DB.Collection(DiffusionCollectionName).UpdateMany(*MongoCtx, diffusionFilter, update)
	if err != nil {
		return err
	}

	return nil
}
