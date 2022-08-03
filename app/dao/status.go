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
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// FindById : StatusIdからstatusを取得
func (r *status) FindById(ctx context.Context, id object.StatusID) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// AddStatus : statusを作成
func (r *status) AddStatus(ctx context.Context, s object.Status) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO status (account_id, content) VALUES (?, ?)", s.AccountID, s.Content)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
