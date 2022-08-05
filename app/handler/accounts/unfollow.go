package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `POST /v1/accounts/{username}/unfollow`
func (h *handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	follower := auth.AccountOf(r)

	username := chi.URLParam(r, "username")

	a := h.app.Dao.Account() // domain/repository の取得
	followee, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	f := h.app.Dao.Follow() // domain/repository の取得
	// フォローの削除
	if err := f.DeleteFollow(ctx, *follower, *followee); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// 相互関係の取得
	relation, err := f.FindRelation(ctx, *follower, *followee)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relation); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
