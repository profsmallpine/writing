package web

import (
	"net/http"

	"github.com/profsmallpine/mid/http/routes"
	"github.com/xy-planning-network/trails/http/resp"
	"github.com/xy-planning-network/trails/http/session"
	"github.com/xy-planning-network/trails/logger"
)

var badCredsFlash = session.Flash{
	Type: session.FlashError,
	Msg:  "That's no good mate.",
}

func (h *Handler) createLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.Redirect(w, r, resp.Url(routes.LoginURL))
		return
	}

	password := r.PostForm.Get("password")
	if password != h.writerKey {
		h.Redirect(w, r, resp.Flash(badCredsFlash), resp.Url(routes.LoginURL))
		return
	}

	s, err := h.session(r.Context())
	if err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Url(routes.LoginURL))
		return
	}

	if err := s.RegisterUser(w, r, 1); err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Url(routes.LoginURL))
		return
	}

	h.Redirect(w, r, resp.Url(routes.NewArticleURL))
}

func (h *Handler) getLogin(w http.ResponseWriter, r *http.Request) {
	if err := h.Html(w, r, resp.Tmpls("tmpl/base.tmpl", "tmpl/login.tmpl")); err != nil {
		h.Logger.Info(err.Error(), &logger.LogContext{Error: err})
	}
}

func (h *Handler) logoff(w http.ResponseWriter, r *http.Request) {
	s, err := h.session(r.Context())
	if err != nil {
		h.Redirect(w, r, resp.Err(err), resp.Url(routes.RootURL))
		return
	}

	if err := s.Delete(w, r); err != nil {
		h.Redirect(w, r, resp.Err(err), resp.Url(routes.RootURL))
		return
	}

	h.Redirect(w, r, resp.Url(routes.LoginURL))
}
