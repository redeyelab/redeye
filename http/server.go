package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*httprouter.Router
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func Health(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// GetServer gets the server
func NewRouter() *Server {
	router := httprouter.New()

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

	router.GET("/", Index)
	router.GET("/health/:param", Health)
	server := &Server{
		Router: router,
	}
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

// AddHandler
func (srv *Server) AddHandler(path string, f httprouter.Handle) {
	srv.Router.GET(path, f)
}

func PlayHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "playing")
}
