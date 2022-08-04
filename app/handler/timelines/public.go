package timelines

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/timelines/public`
func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q PublicQuery
	q.setQuery(r.URL.Query())

	t := h.app.Dao.Timeline() // domain/repository の取得
	statuses, err := t.GetPublic(ctx)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}