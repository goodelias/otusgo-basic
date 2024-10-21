package entity

import (
	"encoding/json"

	"github.com/goodelias/otusgo-basic/hw09_serialize/pb"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(*b)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	temp := &Alias{}
	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}
	*b = Book(*temp)
	return nil
}

func SerializeBooksToJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeBooksFromJSON(data []byte) ([]Book, error) {
	var books []Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, err
	}
	return books, nil
}

func SerializeBooksToPb(books []*pb.Book) ([]byte, error) {
	bookList := pb.BookList{Books: books}
	return proto.Marshal(&bookList)
}

func DeserializeBooksFromPb(data []byte) (*pb.BookList, error) {
	var bookList pb.BookList
	if err := proto.Unmarshal(data, &bookList); err != nil {
		return nil, err
	}
	return &bookList, nil
}
