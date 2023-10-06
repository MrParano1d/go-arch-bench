package books

type mockBook struct {
	id    string
	title string
}

func (b *mockBook) ID() string {
	return b.id
}

type IdentifiableBook interface {
	ID() string
}

type BooksMemoryStorage[B IdentifiableBook] struct {
	books map[string]B
}

var _ BooksStorage[*mockBook] = &BooksMemoryStorage[*mockBook]{}

func NewBooksMemoryStorage[B IdentifiableBook]() *BooksMemoryStorage[B] {
	return &BooksMemoryStorage[B]{
		books: make(map[string]B),
	}
}

func (s *BooksMemoryStorage[B]) Create(book B) (B, error) {
	s.books[book.ID()] = book
	return book, nil
}

func (s *BooksMemoryStorage[B]) FetchAll() ([]B, error) {
	books := make([]B, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books, nil
}

func (s *BooksMemoryStorage[B]) Fetch(id string) (B, error) {
	return s.books[id], nil
}

func (s *BooksMemoryStorage[B]) Update(id string, book B) (B, error) {
	s.books[id] = book
	return book, nil
}

func (s *BooksMemoryStorage[B]) Delete(id string) error {
	delete(s.books, id)
	return nil
}
