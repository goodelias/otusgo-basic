package entity

import (
	"encoding/json"
	"testing"

	"github.com/goodelias/otusgo-basic/hw09_serialize/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestBook_MarshalJSON(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Go Basic",
		Author: "Alan Donovan",
		Year:   2015,
		Size:   315,
		Rate:   4.5,
	}

	want := []byte(`{"id":1,"title":"Go Basic","author":"Alan Donovan","year":2015,"size":315,"rate":4.5}`)
	data, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Marshaling book %v", err)
	}

	require.Equal(t, want, data)
}

func TestBook_UnmarshalJSON(t *testing.T) {
	data := `{"id":1,"title":"Go Basic","author":"Alan Donovan","year":2015,"size":315,"rate":4.5}`
	want := Book{
		ID:     1,
		Title:  "Go Basic",
		Author: "Alan Donovan",
		Year:   2015,
		Size:   315,
		Rate:   4.5,
	}

	var book Book
	err := json.Unmarshal([]byte(data), &book)
	if err != nil {
		t.Fatalf("Unmarshaling json data %v", err)
	}

	require.Equal(t, want, book)
}

func TestSerializeBooksToJSON(t *testing.T) {
	books := []Book{
		{
			ID:     1,
			Title:  "Go Basic",
			Author: "Alan Donovan",
			Year:   2015,
			Size:   315,
			Rate:   4.5,
		},
	}

	want := []byte(`[{"id":1,"title":"Go Basic","author":"Alan Donovan","year":2015,"size":315,"rate":4.5}]`)

	data, err := SerializeBooksToJSON(books)
	if err != nil {
		t.Fatalf("serializing books to JSON: %v", err)
	}

	assert.Equal(t, want, data)
}

func TestDeserializeBooksFromJSON(t *testing.T) {
	data := []byte(`[{"id":1,"title":"Go Basic","author":"Alan Donovan","year":2015,"size":315,"rate":4.5}]`)

	want := []Book{
		{
			ID:     1,
			Title:  "Go Basic",
			Author: "Alan Donovan",
			Year:   2015,
			Size:   315,
			Rate:   4.5,
		},
	}

	books, err := DeserializeBooksFromJSON(data)
	if err != nil {
		t.Fatalf("unmarshaling books from JSON: %+v", err)
	}

	assert.Equal(t, want, books)
}

func TestSerializeBooksToPb(t *testing.T) {
	books := []*pb.Book{
		{
			Id:     1,
			Title:  "Go Basic",
			Author: "Alan Donovan",
			Year:   2015,
			Size:   315,
			Rate:   4.5,
		},
	}

	data, err := SerializeBooksToPb(books)
	require.NoError(t, err, "SerializeBooksToPb should not return an error")

	var gotBookList pb.BookList
	err = proto.Unmarshal(data, &gotBookList)
	require.NoError(t, err, "Unmarshal should not return an error")

	expectedBookList := &pb.BookList{Books: books}
	require.True(t, proto.Equal(expectedBookList, &gotBookList), "Serialized and deserialized data should match original")
}

func TestDeserializeBooksFromPb(t *testing.T) {
	expectedBooks := []*pb.Book{
		{
			Id:     1,
			Title:  "Go Basic",
			Author: "Alan Donovan",
			Year:   2015,
			Size:   315,
			Rate:   4.5,
		},
	}
	expectedBookList := &pb.BookList{Books: expectedBooks}

	data, err := proto.Marshal(expectedBookList)
	require.NoError(t, err, "Failed to marshal test data")

	gotBookList, err := DeserializeBooksFromPb(data)
	require.NoError(t, err, "DeserializeBooksFromPb failed")

	require.True(t, proto.Equal(expectedBookList, gotBookList), "Deserialized data does not match expected")
}
