package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/derek-elliott/advent-of-code-2017/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Route holds the route information
type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Pattern     string           `json:"pattern"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

// Routes holds multiple routes
type Routes []Route

// InputDataIntList represents a request that holds a list of ints
type InputDataIntList struct {
	Data []int `json:"data"`
}

// DaySixResponse holds the response from DaySixHandler
type DaySixResponse struct {
	Jumps int `json:"jumps"`
}

var routes Routes

func init() {
	routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
		Route{
			"Day6",
			"POST",
			"/2017/daysix",
			DaySixHandler,
		},
	}
}

// NewRouter creates the routes held in the Routes struct and wrapps them with a logger.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(routes)
}

// DaySixHandler parses the request body for DaySix and serializes the response
func DaySixHandler(w http.ResponseWriter, r *http.Request) {
	var input InputDataIntList
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("Unable to read request body")
	}
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.WithError(err).Error("Unable to deserialize request body")
	}
	jumps := handlers.DaySix(input.Data)

	if err = json.NewEncoder(w).Encode(DaySixResponse{jumps}); err != nil {
		log.WithError(err).Error("Unable to serialize response")
	}
}
