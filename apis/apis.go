package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/shots-fired/shots-store/streamers"
)

var streamersEngine streamers.Engine

// NewRouter returns a new router
func NewRouter(engine streamers.Engine) *mux.Router {
	streamersEngine = engine
	r := mux.NewRouter()
	r.HandleFunc("/streamers", getAllStreamers).Methods("GET")
	r.HandleFunc("/streamers/{id}", getStreamer).Methods("GET")
	r.HandleFunc("/streamers/{id}", setStreamer).Methods("POST", "PUT")
	r.HandleFunc("/streamers/{id}", deleteStreamer).Methods("DELETE")

	return r
}

// NewServer returns the web server for the router
func NewServer(r *mux.Router) *http.Server {
	return &http.Server{
		Handler:      r,
		Addr:         os.Getenv("SERVER_ADDRESS"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
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
	body, _ := ioutil.ReadAll(r.Body)
	var streamer streamers.Streamer
	err := json.Unmarshal(body, &streamer)
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
