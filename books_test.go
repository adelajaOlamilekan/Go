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

	catalog := getTestCatalog()
	got:= books.GetAllBooks(catalog)

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
	got, ok := books.GetBook(catalog, "abc")

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
	_, ok := books.GetBook(catalog, "nonexistent id")
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
	_, book_existed := books.GetBook(catalog, want.Id)

	if book_existed{
		t.Fatalf("Book with id %v exists already", want.Id)
	}

	books.AddBook(catalog, want)

	// if !ok{
	// 	t.Fatalf("Book %v with id %v exists already", got, got.Id)
	// }

	_, book_added := books.GetBook(catalog, "345")

	if !book_added{
		t.Fatalf("Added book not found")
	}
}

func getTestCatalog() map[string]books.Book{
	return map[string]books.Book{
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