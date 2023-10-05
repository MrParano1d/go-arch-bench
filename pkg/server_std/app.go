package serverstd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrparano1d/archs-go/pkg/auth"
	"github.com/mrparano1d/archs-go/pkg/session"
)

type User struct {
	Identifier string `json:"id"`
	Name       string `json:"username"`
	Password   string `json:"-"`
}

func (u *User) ID() string {
	return u.Identifier
}

func (u *User) Username() string {
	return u.Name
}

func (u *User) SetID(id string) {
	u.Identifier = id
}

func Serve() error {
	sessStorage := session.NewInMemoryStorage()
	authStorage := auth.NewAuthMemoryStorage[*User]()

	server := New[*User](sessStorage, authStorage, func(r chi.Router) {})

	log.Println("Listening on :8080")
	return http.ListenAndServe(":8080", server)
}

// curl -X POST -d '{"username":"test","password":"test"}' http://localhost:8080/login
// curl -X POST -d '{"username":"test","password":"test"}' http://localhost:8080/register
