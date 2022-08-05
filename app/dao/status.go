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
	entity := new(e)
	err := r.db.QueryRowxContext(
		ctx,
		`SELECT
		s.create_at AS 'create_s_at', a.create_at AS 'create_a_at', s.*, a.*
		FROM status AS s
		JOIN account AS a ON s.account_id = a.id
		WHERE s.id = ?`,
		id,
	).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	var status = &object.Status{
		AccountID: entity.AccountID,
		Content:   entity.Content,
		CreateAt:  entity.CreateStatusAt,
	}
	status.Account = &object.Account{
		Username:     entity.Username,
		PasswordHash: entity.PasswordHash,
		DisplayName:  entity.DisplayName,
		Avatar:       entity.Avatar,
		Header:       entity.Header,
		Note:         entity.Note,
		CreateAt:     entity.CreateAccountAt,
	}

	return status, nil
}

// AddStatus : statusを作成
func (r *status) AddStatus(ctx context.Context, s object.Status) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO status (account_id, content) VALUES (?, ?)", s.AccountID, s.Content)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// DeleteStatus : statusを削除
func (r *status) DeleteStatus(ctx context.Context, id object.StatusID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM status WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// joinテーブルを取得するための、埋め込み構造体
type e struct {
	object.DBStatus
	object.DBAccount
	ID       int64
	CreateAt object.DateTime `db:"create_at"`
}
