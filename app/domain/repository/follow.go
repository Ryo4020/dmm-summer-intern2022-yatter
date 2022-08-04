package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Follow interface {
	// Fetch follow relationship
	FindRelation(ctx context.Context, user object.Account, another object.Account) (*object.Relation, error)
	// Create follow relationship
	AddFollow(ctx context.Context, follower object.Account, followee object.Account) error
}
