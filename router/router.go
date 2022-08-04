package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github/go-backend/actions"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		especificRoute := router.Methods(route.Method).Path(route.Pattern).Subrouter()

		especificRoute.Name(route.Name).Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Healthy",
		"GET",
		"/api/healthy",
		actions.Healthy,
	},
	Route{
		"Get all books",
		"GET",
		"/api/heroes",
		actions.GetHeroes,
	},
	Route{
		"Get hero by id",
		"GET",
		"/api/heroes/{id}",
		actions.GetHeroByID,
	},
	// Route{
	// 	"Create hero",
	// 	"POST",
	// 	"/api/heroes",
	// 	actions.AddHero,
	// },
	// Route{
	// 	"Update book",
	// 	"PUT",
	// 	"/api/heroes/{id}",
	// 	actions.UpdateHero,
	// },
}
