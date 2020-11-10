package database

type User struct {
	ID string `bson:"_id"`

	Profile Profile `bson:"profile"`
	Setting Setting `bson:"setting"`

	JoinedGroupIds  []string        `bson:"JoinedGroupIds"`
	InvitedGroupIds []string        `bson:"InvitedGroupIds"`
	FriendRequests  []FriendRequest `bson:"FriendRequests"`
	FriendIds       []string        `bson:"FriendIds"`
	BlockedIds      []string        `bson:"BlockedIds"`
	DeletedIds      []string        `bson:"DeletedIds"`
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
	PrivateUserID              string `bson:"omitempty"`
	AllowSearchByPrivateUserID bool
	Email                      string `bson:"omitempty"`
	AllowSearchByEmail         bool
	UserTicket                 string `bson:"omitempty"`
	AllowSearchByUserTicket    bool
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
