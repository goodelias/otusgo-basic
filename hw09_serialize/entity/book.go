package entity

import "encoding/json"

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
