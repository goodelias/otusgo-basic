package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
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

	want := `{"id":1,"title":"Go Basic","author":"Alan Donovan","year":2015,"size":315,"rate":4.5}`
	data, err := json.Marshal(book)
	if err != nil {
		t.Fatalf("Marshaling book %v", err)
	}

	require.Equal(t, want, string(data))
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
