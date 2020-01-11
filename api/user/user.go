package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Routes sets up all user routes
func Routes() *chi.Mux {
	router := chi.NewRouter()

	// Protected routes
	router.Group(func(router chi.Router) {
		/* router.Use(jwtauth.Verifier(tokenAuth))

		router.Use(jwtauth.Authenticator) */

		router.Delete("/delete", Delete)
		router.Post("/edit", Edit)
	})

	// Public routes
	router.Group(func(router chi.Router) {
		router.Post("/login", Login)
		router.Post("/logout", Logout)
		router.Post("/register", Register)
	})

	return router
}

// Login admin login function
func Login(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Successfully logged in"
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

// Logout admin login function
func Logout(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Successfully logged out"
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

// Delete admin login function
func Delete(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Successfully deleted"
	render.JSON(w, r, response)
}

// Edit admin login function
func Edit(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Successfully edited"
	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

// Register admin login function
func Register(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Successfully registered"
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}







