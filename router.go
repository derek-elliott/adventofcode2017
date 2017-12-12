package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route holds the route information
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds multiple routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TopicIndex",
		"GET",
		"/v1/topics",
		TopicIndex,
	},
	Route{
		"TopicShow",
		"GET",
		"/v1/topics/{topicID}",
		TopicShow,
	},
	Route{
		"TopicPut",
		"PUT",
		"/v1/topics/{topicID}",
		TopicPut,
	},
}

// NewRouter creates the routes held in the Routes struct and wrapps them with a logger.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
