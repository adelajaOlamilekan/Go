package main

import (
	"books"
	"fmt"
	"os"
)

func main(){
	// program_name := os.Args[0] 

	if len(os.Args) != 2{
		fmt.Println("Usage: find <BOOK ID>")
		return
	}

	book_id := os.Args[1]
	catalog := books.GetCatalog()
	book, ok := books.GetBook(catalog, book_id)
	if !ok{
		fmt.Println("Sorry, I couldn't find that book in the catalog.")
		return
	}

	// fmt.Println(program_name)
	fmt.Println(books.BookToString(book))
}