package router

import (
	"app/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func Router() http.Handler {
	app := chi.NewRouter()

	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	app.Route("/", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, map[string]interface{}{
				"text": "pong",
			})
		})
		r.Handle("/account/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			backendURL, _ := url.Parse(config.GetUrlAccountService())
			proxy := httputil.NewSingleHostReverseProxy(backendURL)

			req.URL.Host = backendURL.Host
			req.URL.Scheme = backendURL.Scheme
			req.Header.Set("Host", backendURL.Host)

			proxy.ServeHTTP(w, req)
		}))

		r.Handle("/shop/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			backendURL, _ := url.Parse(config.GetUrlShopService())
			proxy := httputil.NewSingleHostReverseProxy(backendURL)

			req.URL.Host = backendURL.Host
			req.URL.Scheme = backendURL.Scheme
			req.Header.Set("Host", backendURL.Host)

			proxy.ServeHTTP(w, req)
		}))

		r.Handle("/product/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			backendURL, _ := url.Parse(config.GetUrlProductService())
			proxy := httputil.NewSingleHostReverseProxy(backendURL)

			req.URL.Host = backendURL.Host
			req.URL.Scheme = backendURL.Scheme
			req.Header.Set("Host", backendURL.Host)

			proxy.ServeHTTP(w, req)
		}))

		r.Handle("/order/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			backendURL, _ := url.Parse(config.GetUrlOrderService())
			proxy := httputil.NewSingleHostReverseProxy(backendURL)

			req.URL.Host = backendURL.Host
			req.URL.Scheme = backendURL.Scheme
			req.Header.Set("Host", backendURL.Host)

			proxy.ServeHTTP(w, req)
		}))
	})

	log.Println("Sevice h-shop-be-api-gateaway starting success!")

	return app
}
