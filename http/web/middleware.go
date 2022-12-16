package web

import (
	"net/http"
	"strings"

	"github.com/profsmallpine/writing/http/routes"
	"github.com/xy-planning-network/trails/http/middleware"
)

func (h *Handler) allowIPWhitelist() middleware.Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIP := getIpAddress(r.Header.Get("X-Forwarded-For"))

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

func getIpAddress(xForwardedFor string) string {
	ips := strings.Split(xForwardedFor, ",")
	if len(ips[0]) > 0 {
		return ips[0]
	}
	return "0.0.0.0"
}
