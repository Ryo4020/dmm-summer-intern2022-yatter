package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	// Fetch home timeline
	GetHome(ctx context.Context, accounts []*object.Account) ([]*object.Status, error)
	// Fetch public timeline
	GetPublic(ctx context.Context, max_id object.StatusID, since_id object.StatusID, limit int) ([]*object.Status, error)
}
