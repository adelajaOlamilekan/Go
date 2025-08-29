package books_test

import (
	"testing"	
	"books"
	"slices"
	"cmp"
	// "fmt"
)

func TestBokToString_FormatsBookInfoAsString(t *testing.T){
	input := books.Book{
		Title: "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want:= "Sea Room by Adam Nicolson (copies: 2)"
	got:= input.String()

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

	catalog := getTestCatalog()
	got:= catalog.GetAllBooks()

	// fmt.Println(got)
	slices.SortFunc(got, func(a, b books.Book) int{
		return cmp.Compare(a.Author, b.Author)
	})

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

	catalog := getTestCatalog()
	got, ok := catalog.GetBook("abc")

	if !ok{
		t.Fatalf("got %v", got)
	}

	if want != got {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestBook_ReturnFalseWhenBookNotFound(t *testing.T){
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := catalog.GetBook("nonexistent id")
	if ok{
		t.Fatalf("want false for nonexistent ID got true")
	}
}

func TestBook_AddBook(t* testing.T){
	t.Parallel()
	want := books.Book{
		Id: "345",
		Title: "The Prize of all the Oceams",
		Author: "Giyo Williams",
		Copies: 2,
	}

	catalog := getTestCatalog()
	_, book_existed := catalog.GetBook(want.Id)

	if book_existed{
		t.Fatalf("Book with id %v exists already", want.Id)
	}

	catalog.AddBook(want)

	// if !ok{
	// 	t.Fatalf("Book %v with id %v exists already", got, got.Id)
	// }

	_, book_added := catalog.GetBook("345")

	if !book_added{
		t.Fatalf("Added book not found")
	}
}

func TestSetCopies_SetsNumberOfCopies(t *testing.T){
	t.Parallel()

	book := books.Book{
		Copies: 5,
	}

	err := book.SetCopies(12)

	if err != nil{
		t.Fatal(err)
	}

	if book.Copies != 12{
		t.Errorf("want 12 copies, got %d", book.Copies)
	}
}

func TestSetCopies_ReturnErrorForNegativeCopies(t *testing.T){
	t.Parallel()

	book := books.Book{}

	err := book.SetCopies(-1)
	if err == nil{
		t.Error("Error setting copies of book")
	}
}

func getTestCatalog() books.Catalog{
	return books.Catalog{
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