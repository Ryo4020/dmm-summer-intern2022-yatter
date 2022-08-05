package object

type (
	FollowID = int64

	// Follow follow
	Follow struct {
		// The internal ID of the follow
		ID FollowID

		// The account ID of the follower
		FollowerID AccountID `db:"follower_id"` // フォローする人

		// The account ID of the followee
		FolloweeID AccountID `db:"followee_id"` // フォローされる人

		// The time the follow was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}

	Relation struct {
		Following  bool `json:"following"`
		FollowedBy bool `json:"followed_by"`
	}
)
