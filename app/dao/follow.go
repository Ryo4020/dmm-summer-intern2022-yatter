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
	// Implementation for repository.Follow
	follow struct {
		db *sqlx.DB
	}
)

// Create follow repository
func NewFollow(db *sqlx.DB) repository.Follow {
	return &follow{db: db}
}

// FindRelation : 2つのアカウントのフォロー関係を取得
func (r *follow) FindRelation(ctx context.Context, user object.Account, another object.Account) (*object.Relation, error) {
	relation := new(object.Relation)

	entity := new(object.Follow)
	err := r.db.QueryRowxContext(ctx, "select * from follow where follower_id = ? and followee_id = ?", user.ID, another.ID).StructScan(entity)
	if err == nil {
		relation.Following = true
	} else if errors.Is(err, sql.ErrNoRows) {
		relation.Following = false
	} else {
		return nil, fmt.Errorf("%w", err)
	}

	err = r.db.QueryRowxContext(ctx, "select * from follow where follower_id = ? and followee_id = ?", another.ID, user.ID).StructScan(entity)
	if err == nil {
		relation.FollowedBy = true
	} else if errors.Is(err, sql.ErrNoRows) {
		relation.FollowedBy = false
	} else {
		return nil, fmt.Errorf("%w", err)
	}

	return relation, nil
}

// AddFollow : followerとfollowee間のフォロー関係を作成
func (r *follow) AddFollow(ctx context.Context, follower object.Account, followee object.Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO follow (follower_id, followee_id) VALUES (?, ?)", follower.ID, followee.ID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
