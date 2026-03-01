package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextId = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}
	newContact := Contact{
		ID:    nextId,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextId++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %v\n", newContact)
}

func findContactByName(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func ListContacts() {
	fmt.Println("Contact List:")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func main() {
	ListContacts()

	addContact("Alice", "alice@example.com", "123-456-7890")
	addContact("Bob", "bob@example.com", "098-765-4321")
	addContact("Alice", "alice@example.com", "123-456-7890") // duplicate to test addContact
	ListContacts()

	name := "Hugo"
	contact := findContactByName(name)
	if contact != nil {
		fmt.Printf("Contact found: %v\n", *contact)
	} else {
		fmt.Printf("Contact not found: %v\n", name)
	}
}
