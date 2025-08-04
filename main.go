package main

import "fmt"

func main(){
	fmt.Println("Books in stock: ")
	var title = "New School Mathematics"
	var author = "Adelaja Qowiyyu"
	printBook(title, author)
	title = "New School Pyhsics"
	author = "Asalu Paul"
	printBook(title, author)
	// var book = "'Master and Commander', by Patrick O'Brian"
	// fmt.Println(book)
	// book = "'A Morbid Taste for Bones', by Ellis Peters"
	// fmt.Println(book)
	// hello()
}

func printBook(title, author string){
	fmt.Println(title, "by", author)
}

func hello(){
	fmt.Println("Hello, World")
}