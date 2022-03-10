package main

import (
	"fmt"
	"os"
	"strings"
)

// books variable stores books that system has.
var books = []string{"A Tale of Two Cities", "The Hobbit","Harry Potter and the Philosopher's Stone", "The Little Prince", "Dream of the Red Chamber", "And Then There Were None", "The Lion, the Witch and the Wardrobe", "She: A History of Adventure", "The Da Vinci Code", "The Alchemist", }

func main() {
	command := os.Args[1] 
	switch command { 
	case "list":
		list()
	case "search":
		searchedBook := strings.Join(os.Args[2:]," ")
		fmt.Println(search(searchedBook, books)) 
	default:
		fmt.Println("Please enter a valid command") 
	}
}

// search function searches given string(searchItem) in the given list(list) and returns a string.
func search(searchItem string, list []string) string{
	for _, book := range books {
		if strings.EqualFold(book, searchItem){
			return book
	}
}
	return "Entered book is not in the list"
}

// list function prints books in the list.
func list(){
	for _, book := range books{ 
		fmt.Println(book)
	}
}