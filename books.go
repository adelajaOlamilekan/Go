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

func BookToString(book Book) string{
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func GetAllBooks(catalog map[string]Book)[]Book{
	return slices.Collect(maps.Values(catalog))
}

func GetBook(catalog map[string]Book, book_id string)(Book, bool){
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

func AddBook(catalog map[string]Book, book Book){
	catalog[book.Id] = book
}

func GetCatalog()map[string] Book{
	return map[string]Book{
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