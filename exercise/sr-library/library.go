//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	title string
}

type LibraryItem struct {
	book     Book
	checkIn  time.Time
	checkOut time.Time
	landed   bool
}

type Member struct {
	id        int
	name      string
	curBookid int
}

type Library struct {
	books   map[string]LibraryItem
	members map[int]Member
}

func printLendedBooks(lib Library) {
	for key, value := range lib.books {
		if value.landed {
			fmt.Println("%s is landed", key)
		}
	}
}

func main() {
	lib := Library{
		books: {

		},
	}
}
