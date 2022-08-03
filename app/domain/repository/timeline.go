package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	// Fetch public timeline
	GetPublic(ctx context.Context) ([]*object.Status, error)
}
