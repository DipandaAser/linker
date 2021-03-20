package link

const (
	// CollectionName
	CollectionName = "Links"
)

type Link struct {
	ID       string
	GroupsID [2]string
	Active   bool
	//TotalMessage is the number off message receiving by this link
	TotalMessage int `bson:"TotalMessage"`
	CreatedAt    string
	UpdatedAt    string
}

func CreateLink(GroupsID [2]string) {

}
