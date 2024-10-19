package main

import (
	"encoding/json"
	"log"

	"github.com/goodelias/otusgo-basic/hw09_serialize/entity"
	"github.com/goodelias/otusgo-basic/hw09_serialize/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	book := entity.Book{
		ID:     1,
		Title:  "Go Basic",
		Author: "Alan Donovan",
		Year:   2015,
		Size:   315,
		Rate:   4.5,
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Println("Error marshaling Book to JSON:", err)
		return
	}
	log.Println("JSON output:", data)

	var book2 entity.Book
	err = json.Unmarshal(data, &book2)
	if err != nil {
		log.Println("Error unmarshaling JSON to Book:", err)
		return
	}
	log.Printf("Deserialized Book: %+v\n", book2)

	book3 := &pb.Book{
		Id:     2,
		Title:  "title",
		Author: "author",
		Year:   2021,
		Size:   300,
		Rate:   3.4,
	}

	{
		data, err := proto.Marshal(book3)
		if err != nil {
			return
		}
		log.Printf("Serialised pb book %+v\n", data)

		var book4 pb.Book

		err = proto.Unmarshal(data, &book4)
		if err != nil {
			log.Fatalf("unmarshalling pb book %v", err)
		}
		log.Printf("Unmarshaled pb book: %+v\n", &book4)
	}
}
