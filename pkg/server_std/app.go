package serverstd

import (
	"log"
	"net/http"

	"github.com/mrparano1d/archs-go/pkg/auth"
	"github.com/mrparano1d/archs-go/pkg/books"
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

type Book struct {
	BookID   string `json:"id"`
	Title    string `json:"title"`
	AuthorID string `json:"author_id"`
}

func (b *Book) ID() string {
	return b.BookID
}

func Serve() error {
	sessStorage := session.NewInMemoryStorage()
	authStorage := auth.NewAuthMemoryStorage[*User]()
	booksStorage := books.NewBooksMemoryStorage[*Book]()

	server := New[*User](sessStorage, authStorage, booksStorage)

	log.Println("Listening on :8080")
	return http.ListenAndServe(":8080", server)
}

// curl -X POST -d '{"username":"test","password":"test"}' http://localhost:8080/login
// curl -X POST -d '{"username":"test","password":"test"}' http://localhost:8080/register
