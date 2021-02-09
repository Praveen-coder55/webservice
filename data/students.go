package data

import (
	"encoding/json"
	"io"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Description string  `json:"description"`
	// Price       float32 `json:"price"`
	// SKU         string  `json:"sku"`
	// CreatedOn   string  `json:"-"`
	// UpdatedOn   string  `json:"-"`
	// DeletedOn   string  `json:"-"`
}

type Students []*Student

func (s *Students) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

func GetStudents() Students {
	return studentList
}

var studentList = []*Student{
	&Student{
		ID:   1,
		Name: "Varun",
		// Description: "Frotty milky coffee",
		// Price:       2.45,
		// SKU:         "abc323",
		// CreatedOn:   time.Now().UTC().String(),
		// UpdatedOn:   time.Now().UTC().String(),
		// DeletedOn:   "",
	},
	&Student{
		ID:   2,
		Name: "Praveen",
		// Description: "Short and strong coffee without milk",
		// Price:       1.99,
		// SKU:         "fjd34",
		// CreatedOn:   time.Now().UTC().String(),
		// UpdatedOn:   time.Now().UTC().String(),
		// DeletedOn:   "",
	},
}
