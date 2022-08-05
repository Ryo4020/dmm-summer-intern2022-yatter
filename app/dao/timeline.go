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

// GetHome : home timelineを取得
func (r *timeline) GetHome(ctx context.Context, accounts []*object.Account) ([]*object.Status, error) {
	ids := make([]*object.AccountID, len(accounts), cap(accounts))
	// フォローしているアカウントのIDのslice生成
	for i, v := range accounts {
		ids[i] = &v.ID
	}

	s, params, err := sqlx.In(
		`SELECT
		s.create_at AS 'create_s_at', a.create_at AS 'create_a_at', s.*, a.*
		FROM status AS s
		JOIN account AS a ON s.account_id = a.id
		WHERE a.id IN (?)`,
		ids)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	rows, err := r.db.QueryContext(ctx, s, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	entity := make([]*e, 0)
	err = sqlx.StructScan(rows, &entity)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	timeline := make([]*object.Status, len(entity), cap(entity))
	for i, v := range entity {
		timeline[i] = &object.Status{
			AccountID: v.AccountID,
			Content:   v.Content,
			CreateAt:  v.CreateStatusAt,
		}
		timeline[i].Account = &object.Account{
			Username:     v.Username,
			PasswordHash: v.PasswordHash,
			DisplayName:  v.DisplayName,
			Avatar:       v.Avatar,
			Header:       v.Header,
			Note:         v.Note,
			CreateAt:     v.CreateAccountAt,
		}
	}

	return timeline, nil
}

// GetPublic : public timelineを取得
func (r *timeline) GetPublic(ctx context.Context, max_id object.StatusID, since_id object.StatusID, limit int) ([]*object.Status, error) {
	var rows *sql.Rows
	var err error
	if max_id != 0 {
		rows, err = r.db.QueryContext(
			ctx,
			`SELECT
			s.create_at AS 'create_s_at', a.create_at AS 'create_a_at', s.*, a.*
			FROM status AS s
			JOIN account AS a ON s.account_id = a.id
			WHERE s.id < ?
			AND s.id > ?
			LIMIT ?`,
			max_id,
			since_id,
			limit)
	} else {
		rows, err = r.db.QueryContext(
			ctx,
			`SELECT
			s.create_at AS 'create_s_at', a.create_at AS 'create_a_at', s.*, a.*
			FROM status AS s
			JOIN account AS a ON s.account_id = a.id
			WHERE s.id > ?
			LIMIT ?`,
			since_id,
			limit)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	entity := make([]*e, 0)
	err = sqlx.StructScan(rows, &entity)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	timeline := make([]*object.Status, len(entity), cap(entity))
	for i, v := range entity {
		timeline[i] = &object.Status{
			AccountID: v.AccountID,
			Content:   v.Content,
			CreateAt:  v.CreateStatusAt,
		}
		timeline[i].Account = &object.Account{
			Username:     v.Username,
			PasswordHash: v.PasswordHash,
			DisplayName:  v.DisplayName,
			Avatar:       v.Avatar,
			Header:       v.Header,
			Note:         v.Note,
			CreateAt:     v.CreateAccountAt,
		}
	}

	return timeline, nil
}
