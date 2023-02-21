// This is an HTTP framework that adds optional hooks to a simple HTTP server, then exposes an
// "Application" interface for us. We must then create an object that implements the
// Application interface. See app.go

// This is a simple form of the pattern implemented by Twirp https://github.com/twitchtv/twirp
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Application interface {
	GetUser(u UserRequest) UserResponse
}

func main() {
	a := app{}
	h := Hook{
		RequestReceived: func(ctx context.Context) { fmt.Println("Request received") },
		ResponseSent:    func(ctx context.Context) { fmt.Println("Response sent") },
	}
	w := NewWrappedServer(a, h)
	if err := http.ListenAndServe(":8080", w); err != nil {
		log.Fatalln(err)
	}
}

type Hook struct {
	RequestReceived func(ctx context.Context)
	ResponseSent    func(ctx context.Context)
}

func NewWrappedServer(a Application, h Hook) http.Handler {
	return WrappedServer{a, h}
}

type WrappedServer struct {
	Application Application
	Hook        Hook
}

func (wrappedServer WrappedServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if wrappedServer.Hook.RequestReceived != nil { // Call RequestReceived hook if it exists
		wrappedServer.Hook.RequestReceived(ctx)
	}
	if r.URL.Path == "/GetUser" && r.Method == "POST" {
		var req UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res := wrappedServer.Application.GetUser(req)
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if wrappedServer.Hook.ResponseSent != nil {
			wrappedServer.Hook.ResponseSent(ctx) // Call ResponseSent hook if it exists
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
