package serverstd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrparano1d/archs-go/pkg/auth"
	"github.com/mrparano1d/archs-go/pkg/books"
	"github.com/mrparano1d/archs-go/pkg/session"
	"go.uber.org/zap"
	"moul.io/chizap"

	json "github.com/bytedance/sonic"
)

type RouterRegister func(r chi.Router)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New[U auth.UserIdentifiable, B any](sessionStorage session.SessionStorage, authStorage auth.AuthStorage[U], bookStorage books.BooksStorage[B]) chi.Router {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	sessionManager := session.NewSessionManager(sessionStorage)
	authManager := auth.NewAuthManager[U](authStorage, sessionManager)

	r := chi.NewRouter()
	r.Use(chizap.New(logger, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {

		var req loginRequest

		err := json.ConfigStd.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sess, err := authManager.Login(r.Context(), req.Username, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.ConfigStd.NewEncoder(w).Encode(sess)
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {

		var user U

		err := json.ConfigStd.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err = authManager.Register(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware(sessionManager))

		r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
			books, err := bookStorage.FetchAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.ConfigStd.NewEncoder(w).Encode(books)
		})

		r.Post("/books", func(w http.ResponseWriter, r *http.Request) {
			var book B

			err := json.ConfigStd.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			book, err = bookStorage.Create(book)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.ConfigStd.NewEncoder(w).Encode(book)
		})

		r.Route("/books/{id}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")

				book, err := bookStorage.Fetch(id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.ConfigStd.NewEncoder(w).Encode(book)
			})

			r.Put("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")

				var book B

				err = json.ConfigStd.NewDecoder(r.Body).Decode(&book)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				book, err = bookStorage.Update(id, book)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.ConfigStd.NewEncoder(w).Encode(book)
			})

			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")

				err = bookStorage.Delete(id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
			})
		})
	})

	return r
}
