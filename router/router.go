package router

import (
	"app/config"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() http.Handler {
	app := chi.NewRouter()

	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	app.Route("/", func(r chi.Router) {
		r.Handle("/account/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			backendURL, _ := url.Parse(config.GetUrlAccountService())
			proxy := httputil.NewSingleHostReverseProxy(backendURL)

			req.URL.Host = backendURL.Host
			req.URL.Scheme = backendURL.Scheme
			req.Header.Set("Host", backendURL.Host)

			proxy.ServeHTTP(w, req)
		}))
	})

	return app
}
