package main

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// The claim is that this logger is comprised only of these methods, not all zap.Logger methods
type Logger interface {
	Error(msg string, fields ...zap.Field)
}

// Even though this Repo interface is comprised of all Repo methods, we still pull it into the implementer
// rather than exporting it from our repository.
type Repo interface {
	GetItem() Item
}

// Server's field types are interfaces, for two reasons:
//  1. We can replace their underlying implementations without changing our Server code
//  2. In the case of Logger, we are claiming our Logger object is comprised only of the zap.Logger
//     methods in our Logger interface
type Server struct {
	Respository Repo
	Logger      Logger
}

func NewServer() Server {
	r := NewRepository()
	return Server{
		Respository: r,
		Logger:      zap.NewExample(),
	}
}

func main() {
	s := NewServer()
	http.Handle("/", s.GetItem())
	http.ListenAndServe(":8080", nil)
}

func (s Server) GetItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i := s.Respository.GetItem()
		err := json.NewEncoder(w).Encode(&i)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
	}
}
