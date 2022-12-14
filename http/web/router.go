package web

import (
	"net/http"

	"github.com/profsmallpine/mid/http/routes"
	"github.com/xy-planning-network/trails/http/router"
)

func (h *Handler) SetupRoutes() {
	authedRoutes := []router.Route{
		{Path: routes.ArticlesURL, Method: http.MethodPost, Handler: h.createArticle},
		{Path: routes.NewArticleURL, Method: http.MethodGet, Handler: h.newArticle},
		{Path: routes.LogoffURL, Method: http.MethodGet, Handler: h.logoff},
	}
	h.AuthedRoutes(h.EmitKeyring().CurrentUserKey(), routes.LoginURL, routes.LogoffURL, authedRoutes)

	siteRoutes := []router.Route{
		{Path: routes.ArticlesURL, Method: http.MethodGet, Handler: h.getArticles},
		{Path: routes.ArticleURL, Method: http.MethodGet, Handler: h.showArticle},
		{Path: routes.LegacyArticleURL, Method: http.MethodGet, Handler: h.showArticle},
		{Path: routes.RootURL, Method: http.MethodGet, Handler: h.root},
	}
	h.HandleRoutes(siteRoutes)

	unauthedRoutes := []router.Route{
		{Path: routes.LoginURL, Method: http.MethodPost, Handler: h.createLogin},
		{Path: routes.LoginURL, Method: http.MethodGet, Handler: h.getLogin},
	}
	h.UnauthedRoutes(h.EmitKeyring().CurrentUserKey(), unauthedRoutes, h.allowIPWhitelist())
}
