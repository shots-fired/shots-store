package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shots-fired/shots-store/streamers"
)

var streamersEngine streamers.Engine

func HostAPIs(engine streamers.Engine) {
	streamersEngine = engine
	r := mux.NewRouter()
	r.HandleFunc("/streamers/{id}", getStreamer).Methods("GET")
	r.HandleFunc("/streamers/{id}", setStreamer).Methods("POST")

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
	log.Printf("Res, Err, %+v, %v", res, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	streamer, marshalErr := json.Marshal(res)
	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(streamer)
}

func setStreamer(w http.ResponseWriter, r *http.Request) {
	err := streamersEngine.SetStreamer(mux.Vars(r)["id"], streamers.Streamer{
		Name: "jake",
	})

	log.Printf("Err, %v", err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
