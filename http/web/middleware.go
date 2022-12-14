package web

import (
	"net/http"

	"github.com/profsmallpine/mid/http/routes"
	"github.com/xy-planning-network/trails/http/middleware"
)

func (h *Handler) allowIPWhitelist() middleware.Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIP := r.Context().Value(middleware.IpAddrCtxKey).(string)

			for _, ip := range h.whitelistIPs {
				if ip == userIP {
					handler.ServeHTTP(w, r)
					return
				}
			}

			http.Redirect(w, r, routes.RootURL, http.StatusTemporaryRedirect)
		})
	}
}
