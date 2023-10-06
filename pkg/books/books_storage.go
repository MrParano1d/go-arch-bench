package books

type BooksStorage[B any] interface {
	Create(book B) (B, error)
	FetchAll() ([]B, error)
	Fetch(id string) (B, error)
	Update(id string, book B) (B, error)
	Delete(id string) error
}
