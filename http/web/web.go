package web

import (
	"context"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/profsmallpine/writing/domain"
	"github.com/xy-planning-network/trails/http/session"
	"github.com/xy-planning-network/trails/logger"
	"github.com/xy-planning-network/trails/ranger"
	"gorm.io/gorm"
)

type Handler struct {
	*gorm.DB
	*ranger.Ranger
	writerKey    string
	whitelistIPs []string
}

func NewHandler(db *gorm.DB, rng *ranger.Ranger, key string, ips []string) *Handler {
	return &Handler{db, rng, key, ips}
}

// session helps by retrieving the session.TrailsSessionable from the provided context
func (h *Handler) session(ctx context.Context) (session.TrailsSessionable, error) {
	s, ok := ctx.Value(h.EmitKeyring().SessionKey()).(session.TrailsSessionable)
	if !ok {
		return nil, domain.ErrNoSession
	}
	return s, nil
}

// Set a Decoder instance as a package global, because it caches meta-data about structs, and an
// instance can be shared safely.
var decoder = schema.NewDecoder()

func (h *Handler) parseForm(r *http.Request, reqStructPtr interface{}) error {
	if err := r.ParseForm(); err != nil {
		h.Logger.Error(err.Error(), &logger.LogContext{Error: err})
		return err
	}

	if err := decoder.Decode(reqStructPtr, r.PostForm); err != nil {
		h.Logger.Error(err.Error(), &logger.LogContext{Error: err})
		return err
	}

	return nil
}
