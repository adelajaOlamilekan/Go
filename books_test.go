package books_test

import (
	"testing"	
	"books"
)

func TestBokToString_FormatsBookInfoAsString(t *testing.T){
	input := books.Book{
		Title: "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want:= "Sea Room by Adam Nicolson (copies: 2)"
	got:= books.BookToString(input)

	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}