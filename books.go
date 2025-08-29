package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct{
	Title string
	Author string
	Copies int
	Id string
}

type Catalog map[string] Book

func (book Book) String() string{
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func (catalog Catalog) GetAllBooks() []Book{
	return slices.Collect(maps.Values(catalog))
}

func (catalog Catalog) GetBook(book_id string)(Book, bool){
	// var bookFound bool = false

	// for _, book := range catalog{
	// 	if book.Id == book_id {
	// 		return book, true
	// 	}
	// }

	book,ok := catalog[book_id]

	return book , ok

	// return Book{}
}

func (catalog Catalog) AddBook(book Book){
	catalog[book.Id] = book
}

func GetCatalog() Catalog{
	return Catalog{
		"abcd": {
			Title: "In the Company of Cheerful Ladies",
			Author: "Alexander McCall Smith",
			Copies: 1,
			Id: "abcd",
		},
		"abc": {
			Title: "White Heat",
			Author: "Dominic Sandbrook",
			Copies: 2,
			Id: "abc",
		},
	
	}
}

func (book *Book) SetCopies(copies int) error {
	// fmt.Println("before update, book.Copies = ", book.Copies)

	if copies < 0{
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	return nil
	// fmt.Println("after update, book.Copies = at address", book.Copies)
	// fmt.Println(*book)
}