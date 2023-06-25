package main

import (
	"fmt"
)

// BookClub is a type that contains information about a book club 
type BookClub struct {
	name string
	theme string
	members []string
	readings []string
	discussions []string
}

// NewBookClub creates a new BookClub
func NewBookClub(name, theme string) *BookClub {
	return &BookClub{name: name, theme: theme}
}

// AddMember adds a new member to the book club
func (b *BookClub) AddMember(memberName string) {
	b.members = append(b.members, memberName)
}

// AddReading adds a book to the book club's list of readings
func (b *BookClub) AddReading(bookName string) {
	b.readings = append(b.readings, bookName)
}

// AddDiscussion adds a new discussion topic to the book club
func (b *BookClub) AddDiscussion(discussion string) {
	b.discussions = append(b.discussions, discussion)
}

// PrintInfo prints out all the information for a book club
func (b *BookClub) PrintInfo() {
	fmt.Printf("The book club is named %s and its theme is %s.\n", 
		b.name, b.theme)
	fmt.Printf("Members: %v\n", b.members)
	fmt.Printf("Readings: %v\n", b.readings)
	fmt.Printf("Discussions: %v\n", b.discussions)
}

func main() {
	// Create a new book club
	bc := NewBookClub("Literary Lovers", "Book Discussion")

	// Add members
	bc.AddMember("John")
	bc.AddMember("Jane")
	bc.AddMember("Bob")
	bc.AddMember("Alice")

	// Add readings
	bc.AddReading("The Great Gatsby")
	bc.AddReading("Pride and Prejudice")
	bc.AddReading("To Kill a Mockingbird")

	// Add discussions
	bc.AddDiscussion("Main characters")
	bc.AddDiscussion("Themes")
	bc.AddDiscussion("Symbols")
	bc.AddDiscussion("Analysis")

	// Print out info
	bc.PrintInfo()
}