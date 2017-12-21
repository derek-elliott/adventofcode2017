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
			"Day6Pt1",
			"POST",
			"/2017/daysix/1",
			DaySixPartOneHandler,
		},
		Route{
			"Day6Pt2",
			"POST",
			"/2017/daysix/2",
			DaySixPartTwoHandler,
		}, Route{
			"Day7Pt1",
			"POST",
			"/2017/dayseven/1",
			DaySevenPartOneHandler,
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

// DaySixResponse holds the response from DaySixHandler
type DaySixResponse struct {
	Steps int `json:"steps"`
}

// DaySixPartOneHandler parses the request body for DaySix and serializes the response
func DaySixPartOneHandler(w http.ResponseWriter, r *http.Request) {
	var input InputDataIntList
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("Unable to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.WithError(err).Error("Unable to deserialize request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	steps, _ := handlers.DaySix(input.Data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(DaySixResponse{steps}); err != nil {
		log.WithError(err).Error("Unable to serialize response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// DaySixPartTwoHandler parses the request body for DaySix part 2 and serializes the response
func DaySixPartTwoHandler(w http.ResponseWriter, r *http.Request) {
	var input InputDataIntList
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("Unable to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &input); err != nil {
		log.WithError(err).Error("Unable to deserialize request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, steps := handlers.DaySix(input.Data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(DaySixResponse{steps}); err != nil {
		log.WithError(err).Error("Unable to serialize response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// DaySevenResponse holds the response for day seven
type DaySevenResponse struct {
	Root string `json:"root"`
}

// DaySevenPartOneHandler parses the request for day seven part one and returns the name of the root node
func DaySevenPartOneHandler(w http.ResponseWriter, r *http.Request) {
	var inputData handlers.DaySevenInputs
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).Error("Unable to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &inputData); err != nil {
		log.WithError(err).Error("Unable to deserialize request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	root := handlers.DaySevenPartOne(inputData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(DaySevenResponse{root}); err != nil {
		log.WithError(err).Error("Unable to serialize response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
