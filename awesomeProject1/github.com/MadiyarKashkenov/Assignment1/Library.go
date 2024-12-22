package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	ID         string
	Author     string
	Title      string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

var lib Library

func AddBook(book Book) {
	lib.Books[book.ID] = book

}
func BorrowBook(id string) {
	book, exists := lib.Books[id]
	if !exists {
		fmt.Println("Book not found")

	} else if book.IsBorrowed {
		fmt.Println("Book is borrowed")
	} else {
		book.IsBorrowed = true
		lib.Books[id] = book
		fmt.Println("Book borrowed successfully!")
	}
}
func ReturnBook(id string) {
	book, exists := lib.Books[id]
	if !exists {
		fmt.Println("Book not found!")
	}
	if !book.IsBorrowed {
		fmt.Println("Book is not borrowed!")
	}
	book.IsBorrowed = false
	lib.Books[id] = book
	fmt.Println("Book returned successfully!")
}
func ListBooks() {
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}
func main() {
	var sw int
	lib.Books = make(map[string]Book)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1: Add Book:")
		fmt.Println("2: Borrow Book")
		fmt.Println("3: Return Book")
		fmt.Println("4: List Books")
		fmt.Println("5: Exit:")
		fmt.Println(" ")
		fmt.Print("Enter a request:")

		fmt.Scan(&sw)
		if sw == 1 {
			var id, title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Book Title: ")
			scanner.Scan()
			title = scanner.Text()
			fmt.Print("Enter Book Author: ")
			scanner.Scan()
			author = scanner.Text()
			AddBook(Book{ID: id, Title: title, Author: author})
		} else if sw == 2 {
			var id string
			fmt.Print("Enter Book ID to borrow: ")
			fmt.Scan(&id)
			BorrowBook(id)
		} else if sw == 3 {
			var id string
			fmt.Print("Enter Book ID to return: ")
			fmt.Scan(&id)
			ReturnBook(id)
		} else if sw == 4 {
			ListBooks()
		} else if sw == 5 {
			break
		} else {
			fmt.Println("Invalid input")

		}
	}
}
