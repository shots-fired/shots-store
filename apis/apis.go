package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/shots-fired/shots-common/models"
	"github.com/shots-fired/shots-store/streamers"
)

var streamersEngine streamers.Engine

// NewRouter returns a new router
func NewRouter(engine streamers.Engine) *httprouter.Router {
	streamersEngine = engine
	r := httprouter.New()
	r.GET("/streamers", getAllStreamers)
	r.GET("/streamers/:id", getStreamer)
	r.POST("/streamers/:id", setStreamer)
	r.PUT("/streamers/:id", setStreamer)
	r.DELETE("/streamers/:id", deleteStreamer)

	return r
}

// NewServer returns the web server for the router
func NewServer(r *httprouter.Router) *http.Server {
	return &http.Server{
		Handler:      r,
		Addr:         os.Getenv("SERVER_ADDRESS"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func getStreamer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	res, err := streamersEngine.GetStreamer(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	streamer, _ := json.Marshal(res)

	w.Write(streamer)
}

func getAllStreamers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func setStreamer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, _ := ioutil.ReadAll(r.Body)
	var streamer models.Streamer
	err := json.Unmarshal(body, &streamer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = streamersEngine.SetStreamer(ps.ByName("id"), streamer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func deleteStreamer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := streamersEngine.DeleteStreamer(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
