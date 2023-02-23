// https://google.github.io/styleguide/go/decisions#interfaces
// Go interfaces generally belong in the package that consumes
// values of the interface type, not a package that implements the interface type.
package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Our logger interface provides access only to the methods specified here, not all logrus methods
type Logger interface {
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

// Even though this Repo interface is comprised of all Repo methods, we still pull it into the implementer
// rather than exporting it from our repository.
type Repo interface {
	GetItem() Item
}

// Server's field types are interfaces, for two reasons:
//  1. We can replace their underlying implementations without changing our Server code
//  2. In the case of Logger, we are claiming our Logger object is comprised only of the logrus
//     methods in our Logger interface
type Server struct {
	Repository Repo
	Logger     Logger
}

func NewServer() Server {
	logrus.Error()
	r := NewRepository()
	l := logrus.New()
	return Server{
		Repository: r,
		Logger:     l,
	}
}

func main() {
	s := NewServer()
	http.Handle("/", s.GetItem())
	http.ListenAndServe(":8080", nil)
}

func (s Server) GetItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		i := s.Repository.GetItem()
		err := json.NewEncoder(w).Encode(&i)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
	}
}
