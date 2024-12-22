package main

import (
	"fmt"
	"os"
)

type Book struct {
	ID        string
	Title     string
	Author    string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists && !book.IsBorrowed {
		book.IsBorrowed = true
		l.Books[id] = book
		fmt.Printf("Book '%s' has been borrowed.\n", book.Title)
	} else if exists {
		fmt.Println("This book is already borrowed.")
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists && book.IsBorrowed {
		book.IsBorrowed = false
		l.Books[id] = book
		fmt.Printf("Book '%s' has been returned.\n", book.Title)
	} else if exists {
		fmt.Println("This book is not borrowed.")
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	library := Library{
		Books: make(map[string]Book),
	}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add – Add a book")
		fmt.Println("2. Borrow – Borrow a book")
		fmt.Println("3. Return – Return a book")
		fmt.Println("4. List – List available books")
		fmt.Println("5. Exit – Exit the program")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, author string
			fmt.Println("Enter book ID:")
			fmt.Scan(&id)
			fmt.Println("Enter book title:")
			fmt.Scan(&title)
			fmt.Println("Enter book author:")
			fmt.Scan(&author)
			book := Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			library.AddBook(book)
			fmt.Println("Book added.")
		case 2:
			var id string
			fmt.Println("Enter book ID to borrow:")
			fmt.Scan(&id)
			library.BorrowBook(id)
		case 3:
			var id string
			fmt.Println("Enter book ID to return:")
			fmt.Scan(&id)
			library.ReturnBook(id)
		case 4:
			library.ListBooks()
		case 5:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
