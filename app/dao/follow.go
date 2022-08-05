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
	err := r.db.QueryRowxContext(ctx, "SELECT * FROM follow WHERE follower_id = ? AND followee_id = ?", user.ID, another.ID).StructScan(entity)
	if err == nil {
		relation.Following = true
	} else if errors.Is(err, sql.ErrNoRows) {
		relation.Following = false
	} else {
		return nil, fmt.Errorf("%w", err)
	}

	err = r.db.QueryRowxContext(ctx, "SELECT * FROM follow WHERE follower_id = ? AND followee_id = ?", another.ID, user.ID).StructScan(entity)
	if err == nil {
		relation.FollowedBy = true
	} else if errors.Is(err, sql.ErrNoRows) {
		relation.FollowedBy = false
	} else {
		return nil, fmt.Errorf("%w", err)
	}

	return relation, nil
}

// GetFollowing : アカウントがフォローしているアカウントリストを取得
func (r *follow) GetFollowing(ctx context.Context, account object.Account) ([]*object.Account, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT
		a.*
		FROM account AS a
		JOIN follow AS f ON f.followee_id = a.id
		WHERE f.follower_id = ?`,
		account.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	entity := make([]*object.Account, 0)
	err = sqlx.StructScan(rows, &entity)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// AddFollow : followerとfollowee間のフォロー関係を作成
func (r *follow) AddFollow(ctx context.Context, follower object.Account, followee object.Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO follow (follower_id, followee_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE follower_id = follower_id", follower.ID, followee.ID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// DeleteFollow : followerとfollowee間のフォロー関係を削除
func (r *follow) DeleteFollow(ctx context.Context, unfollower object.Account, unfollowee object.Account) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM follow WHERE follower_id = ? AND followee_id = ?", unfollower.ID, unfollowee.ID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
