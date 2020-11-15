package database

type User struct {
	ID string `bson:"_id"`

	Profile Profile `bson:"profile"`
	Setting Setting `bson:"setting"`

	JoinedGroupIds  []string                 `bson:"JoinedGroupIds"`  //map[uuid]timestamp
	InvitedGroupIds []string                 `bson:"InvitedGroupIds"` //map[guid]timestamp
	FriendRequests  map[string]FriendRequest `bson:"FriendRequests"`  //map[uuid]FriendReq
	FriendIds       []string                 `bson:"FriendIds"`       //map[uuid]timestamp
	BlockedIds      []string                 `bson:"BlockedIds"`      //map[uuid]timestamp
	DeletedIds      []string                 `bson:"DeletedIds"`      //map[uuid]timestamp
	Contacts        map[string]Contact       `bson:"contacts"`        //map[uuid]Contact
	Tags            map[string]Tag           `bson:"tags"`            //map[tagID]Tag
}

type Profile struct {
	Name        string `bson:"name"`
	Bio         string `bson:"bio"`
	IconPath    string `bson:"omitempty"`
	CoverPath   string `bson:"omitempty"`
	TwitterID   string `bson:"omitempty"`
	InstagramID string `bson:"omitempty"`
	GithubID    string `bson:"omitempty"`
}

type Setting struct {
	PrivateUserID              string `bson:"PUserID"`
	AllowSearchByPrivateUserID bool   `bson:"asbPUserID"`
	Email                      string `bson:"Email"`
	AllowSearchByEmail         bool   `bson:"asbEmail"`
	UserTicket                 string `bson:"UTicket"`
	AllowSearchByUserTicket    bool   `bson:"asbUTicket"`
}

type FriendRequest struct {
	FromID               string `bson:"omitempty"`
	ToID                 string `bson:"omitempty"`
	CreatedTime          int64  `bson:"omitempty"`
	Metadata             string `bson:"omitempty"`
	IsAccepted           bool   `bson:"omitempty"`
	IsRejected           bool   `bson:"omitempty"`
	AcceptedOrRejectedAt int64  `bson:"omitempty"`
}
type Contact struct {
	UUID            string   `bson:"uuid"`
	OverWrittenName string   `bson:"owname"`
	Status          int64    `bson:"status"`
	TagIds          []string `bson:"tags"`
}

type Tag struct {
	TagID       string `bson:"tagID"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Color       string `bson:"color"`
	CreatorUUID string `bson:"creator"`
	CreatedTime int64  `bson:"createdTime"`
}
