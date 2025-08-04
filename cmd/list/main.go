package main

import (
	"fmt"
	"books"
)
// import "main_test"


func main(){
	fmt.Println("Books in stock: ")
	// var char strin
	// var title = "New School Mathematics"
	// var author = "Adelaja Qowiyyu"
	// var copies = 15
	// printBook(title, author, copies)
	// title = "New School Pyhsics"
	// author = "Asalu Paul"
	// printBook(title, author, copies)
	// title = 42
	// copies = "None"
	// var book = "'Master and Commander', by Patrick O'Brian"
	// fmt.Println(book)
	// book = "'A Morbid Taste for Bones', by Ellis Peters"
	// fmt.Println(book)
	// hello()
	book := books.Book{
		Title: "Engineering in Plain Sight",
		Author: "Grady Willhouse",
		Copies: 2,
	}
	// printBook(book)
	// TestBokToString_FormatsBookInfoAsString()
	fmt.Println(books.BookToString(book))

}

