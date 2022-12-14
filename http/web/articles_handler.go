package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/profsmallpine/mid/domain"
	"github.com/profsmallpine/mid/http/routes"
	"github.com/xy-planning-network/trails/http/resp"
	"github.com/xy-planning-network/trails/logger"
)

type createArticleReq struct {
	Body    string `schema:"body,required"`
	Slug    string `schema:"slug,required"`
	Summary string `schema:"summary,required"`
	Title   string `schema:"title,required"`
}

func (h *Handler) getArticles(w http.ResponseWriter, r *http.Request) {
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
	if err := h.Html(w, r, resp.Data(data), resp.Tmpls("tmpl/articles/_list.wrapper.tmpl", "tmpl/articles/_list.tmpl")); err != nil {
		h.Logger.Info(err.Error(), &logger.LogContext{Error: err})
	}
}

func (h *Handler) createArticle(w http.ResponseWriter, r *http.Request) {
	// Parse + decode form
	var req createArticleReq
	if err := h.parseForm(r, &req); err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Url(routes.NewArticleURL))
		return
	}

	// Setup article for db insert
	goal := &domain.Article{
		Body:    req.Body,
		Summary: req.Summary,
		Slug:    req.Slug,
		Title:   req.Title,
	}
	if err := h.DB.Create(goal).Error; err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Url(routes.NewArticleURL), resp.Code(http.StatusInternalServerError))
		return
	}

	h.Redirect(w, r, resp.Success("Article created!"), resp.Url(routes.RootURL))
}

func (h *Handler) showArticle(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)[routes.MuxSlugParam]
	article := &domain.Article{}
	if err := h.First(article, "slug = ?", slug).Error; err != nil {
		h.Redirect(w, r, resp.GenericErr(err), resp.Url(routes.RootURL), resp.Code(http.StatusInternalServerError))
		return
	}

	data := map[string]interface{}{"article": article}
	if err := h.Html(w, r, resp.Data(data), resp.Tmpls("tmpl/base.tmpl", "tmpl/articles/show.tmpl")); err != nil {
		h.Logger.Info(err.Error(), &logger.LogContext{Error: err})
	}
}

func (h *Handler) newArticle(w http.ResponseWriter, r *http.Request) {
	if err := h.Html(w, r, resp.Tmpls("tmpl/base.tmpl", "tmpl/articles/new.tmpl")); err != nil {
		h.Logger.Info(err.Error(), &logger.LogContext{Error: err})
	}
}
