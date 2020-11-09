package database

type User struct {
	ID string `bson:"_id"`

	Profile Profile
	Setting Setting

	JoinedGroupIds  []string
	InvitedGroupIds []string
	FriendRequests  []FriendRequest
	FriendIds       []string
	BlockedIds      []string
	DeletedIds      []string
}

type Profile struct {
	Name        string
	Bio         string `bson:"omitempty"`
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
