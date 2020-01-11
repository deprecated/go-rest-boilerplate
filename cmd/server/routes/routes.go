package routes

import (
	"time"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/deprecated/go-rest-boilerplate/api/user"
	"github.com/deprecated/go-rest-boilerplate/api/admin"
)

// InitRoutes initiales our router
func InitRoutes(db *sql.DB, jwtAdminSecret string, jwtUserSecret string) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RealIP,
		middleware.Timeout(30*time.Second),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
		middleware.WithValue("DBCONN", db),
		middleware.WithValue("JWTADMINSECRET", jwtAdminSecret),
		middleware.WithValue("JWTUSERSECRET", jwtUserSecret),
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/user", user.Routes())
		r.Mount("/api/admin", admin.Routes())
	})

	return router
}