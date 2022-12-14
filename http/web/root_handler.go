package web

import (
	"net/http"
	"strconv"

	"github.com/profsmallpine/mid/domain"
	"github.com/xy-planning-network/trails/http/resp"
	"github.com/xy-planning-network/trails/logger"
)

func (h *Handler) root(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	if page == 0 {
		page = 1
	}

	order := "created_at DESC"
	articles := []*domain.Article{}
	pd, err := h.EmitDB().PagedByQuery(&articles, "", nil, order, page, 1)
	if err != nil {
		h.Logger.Error(err.Error(), &logger.LogContext{Error: err}) // NOTE: not returning as there are no other routes to send folks
	}

	data := map[string]interface{}{"articles": pd}
	if err := h.Html(w, r, resp.Data(data), resp.Tmpls("tmpl/base.tmpl", "tmpl/root.tmpl", "tmpl/articles/_list.tmpl")); err != nil {
		h.Logger.Info(err.Error(), &logger.LogContext{Error: err})
	}
}