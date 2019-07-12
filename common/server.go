package common

import (
	"log"
	"net/http"
)

type WebServer struct {
	*http.Server
	*Router
	host string
}

func NewWebServer(hostAddress string) *WebServer {
	router := NewRouter()
	server := &http.Server{Addr: hostAddress, Handler: router.Mux()}
	return &WebServer{
		Server: server, Router: router, host: hostAddress,
	}
}

func (server *WebServer) Start() {
	log.Printf("Starting HTTP service at %v\n", server.host)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
