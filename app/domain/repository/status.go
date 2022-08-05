package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch status by id
	FindById(ctx context.Context, id object.StatusID) (*object.Status, error)
	// Create status with content and account_id
	AddStatus(ctx context.Context, status object.Status) error
	// Delete status
	DeleteStatus(ctx context.Context, id object.StatusID) error
}
