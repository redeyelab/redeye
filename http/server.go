package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*httprouter.Router
}

// GetServer gets the server
func NewRouter() *Server {
	router := httprouter.New()

	// This should take care of any CORS problems
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	server := &Server{
		Router: router,
	}
	server.AddHandler("/", Index)
	router.GET("/", Index)
	router.GET("/health/:param", Health)
	return server
}

// StartServer starts the server
func (srv *Server) StartServer(stop chan interface{}) {
	err := http.ListenAndServe(":2222", srv)
	if err != nil {
		log.Error().Msg("The Listener has failed")
	}
	<-stop
}

// AddHandler allows new handlers to be added
func (srv *Server) AddHandler(path string, f httprouter.Handle) {
	srv.Router.GET(path, f)
}

// Quick health check
func Health(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a quick temporary variable for now
	resp := map[string]bool{
		"ok": true,
	}
	if ps == nil || len(ps) == 0 {
		// Add something specic according to the parameter
	}
	json.NewEncoder(w).Encode(resp)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func PlayHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "playing")
}
