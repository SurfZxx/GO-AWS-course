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
var nextId int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)
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
	for _, contact := range contactList {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Printf("")
}

func main() {
	ListContacts()
}
