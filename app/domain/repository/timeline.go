package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	// Fetch public timeline
	GetPublic(ctx context.Context, max_id object.StatusID, since_id object.StatusID, limit int) ([]*object.Status, error)
}
