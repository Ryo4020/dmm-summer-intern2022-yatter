package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

// Create timeline repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

// GetPublic : public timelineを取得
func (r *timeline) GetPublic(ctx context.Context) ([]*object.Status, error) {
	rows, err := r.db.QueryContext(ctx, "select * from status")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	entity := make([]*object.Status, 0)
	err = sqlx.StructScan(rows, &entity)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
