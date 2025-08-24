package books_test

import (
	"testing"	
	"books"
	"slices"
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

func TestBook_GetAllBooks(t *testing.T){

	want:= []books.Book{
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

	got:= books.GetAllBooks()

	if !slices.Equal(want, got){
		t.Fatalf("want %#v, got %#v", want, got)
	}
	
}

func TestBook_FindBookInCatalogByID(t *testing.T){
	t.Parallel()
	want := books.Book{
		Title: "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		Id: "abc",
	}

	got, ok := books.GetBook("abc")

	if !ok{
		t.Fatalf("got %v", got)
	}

	if want != got {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestBook_ReturnFalseWhenBookNotFound(t *testing.T){
	t.Parallel()
	_, ok := books.GetBook("nonexistent id")
	if ok{
		t.Fatalf("want false for nonexistent ID got true")
	}
}