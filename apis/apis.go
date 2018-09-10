package apis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shots-fired/shots-store/streamers"
)

var streamersEngine streamers.Engine

// New returns a new router
func New(engine streamers.Engine) *mux.Router {
	streamersEngine = engine
	r := mux.NewRouter()
	r.HandleFunc("/streamers", getAllStreamers).Methods("GET")
	r.HandleFunc("/streamers/{id}", getStreamer).Methods("GET")
	r.HandleFunc("/streamers/{id}", setStreamer).Methods("POST", "PUT")
	r.HandleFunc("/streamers/{id}", deleteStreamer).Methods("DELETE")

	return r
}

// Host hosts the web server for the router
func Host(r *mux.Router) {
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8888",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getStreamer(w http.ResponseWriter, r *http.Request) {
	res, err := streamersEngine.GetStreamer(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	streamer, _ := json.Marshal(res)

	w.Write(streamer)
}

func getAllStreamers(w http.ResponseWriter, r *http.Request) {
	res, err := streamersEngine.GetAllStreamers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(res) > 0 {
		streamers, _ := json.Marshal(res)

		w.Write(streamers)
	} else {
		w.Write([]byte("[]"))
	}
}

func setStreamer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var streamer streamers.Streamer
	err = json.Unmarshal(body, &streamer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = streamersEngine.SetStreamer(mux.Vars(r)["id"], streamer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func deleteStreamer(w http.ResponseWriter, r *http.Request) {
	err := streamersEngine.DeleteStreamer(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
