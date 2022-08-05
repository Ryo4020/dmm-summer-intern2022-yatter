package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Follow interface {
	// Fetch follow relationship
	FindRelation(ctx context.Context, user object.Account, another object.Account) (*object.Relation, error)
	// Fetch following of the account
	GetFollowing(ctx context.Context, account object.Account, limit int) ([]*object.Account, error)
	// Fetch followers of the account
	GetFollowers(ctx context.Context, account object.Account, max_id object.AccountID, since_id object.AccountID, limit int) ([]*object.Account, error)
	// Create follow relationship
	AddFollow(ctx context.Context, follower object.Account, followee object.Account) error
	// Delete follow relationship
	DeleteFollow(ctx context.Context, unfollower object.Account, unfollowee object.Account) error
}
