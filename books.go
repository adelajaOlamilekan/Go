package books

import "fmt"

type Book struct{
	Title string
	Author string
	Copies int
	Id string
}

func BookToString(book Book) string{
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

var catalog = []Book{
	{
		Title: "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		Id: "abcd",
	},
	{
		Title: "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		Id: "abc",
	},

}

func GetAllBooks()[]Book{
	return catalog
}

func GetBook(book_id string)(Book, bool){
	// var bookFound bool = false

	for _, book := range catalog{
		if book.Id == book_id {
			return book, true
		}
	}
	return Book{}, false

	// return Book{}
}