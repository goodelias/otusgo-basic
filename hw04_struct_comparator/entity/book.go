package entity

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func newBook(id int, title string, author string, year int, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

type BookBuilder struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBookBuilder() *BookBuilder {
	return &BookBuilder{}
}

func (b *BookBuilder) Build() *Book {
	return newBook(b.id, b.title, b.author, b.year, b.size, b.rate)
}

// GETTERS

func (b *Book) ID() int {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) Rate() float64 {
	return b.rate
}

// SETTERS

func (b *BookBuilder) SetID(id int) *BookBuilder {
	b.id = id
	return b
}

func (b *BookBuilder) SetTitle(title string) *BookBuilder {
	b.title = title
	return b
}

func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.author = author
	return b
}

func (b *BookBuilder) SetYear(year int) *BookBuilder {
	b.year = year
	return b
}

func (b *BookBuilder) SetSize(size int) *BookBuilder {
	b.size = size
	return b
}

func (b *BookBuilder) SetRate(rate float64) *BookBuilder {
	b.rate = rate
	return b
}

func (b *Book) SetRate(rate float64) *Book {
	b.rate = rate
	return b
}
