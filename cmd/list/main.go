package main

import (
	"fmt"
	"books"
)
// import "main_test"


func main(){
	// fmt.Println("Books in stock: ")
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
	// book := books.Book{
	// 	Title: "Engineering in Plain Sight",
	// 	Author: "Grady Willhouse",
	// 	Copies: 2,
	// }
	// // printBook(book)
	// // TestBokToString_FormatsBookInfoAsString()
	// fmt.Println(books.BookToString(book))

	// var all_book = []books.Book{
	// 	{
	// 		Title: "Chasing the Thrill",
	// 		Author: "Daniel Adelabu",
	// 	},
	// 	{
	// 		Title: "Birds and People",
	// 		Author: "Mark Cocker",
	// 	},
	// }

	// var ages = []int{}
	// ages = []int{1,2,3,4}
	// first_book := all_book[0]
	// fmt.Println(all_book, ages, first_book.Title)
	// fmt.Println(len(all_book))

	// all_book[0] = books.Book{
	// 	Title: "The Power Broker",
	// 	Author: "RObert A. Caro",
	// }

	// all_book[0].Title = "The Power isn't Broker"

	// fmt.Println(all_book)
	// new_book := books.Book{
	// 	Title: "The bird is good",
	// 	Author: "Richard C. Mammon",
	// }

	// all_book = append(all_book, new_book, new_book, new_book)

	// fmt.Println(all_book)
	catalog := books.GetCatalog()
	for index, book := range books.GetAllBooks(catalog){
		fmt.Println(index, books.BookToString(book))
	}
}

