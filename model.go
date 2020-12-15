package main

import "time"

type Contact struct {
	ContactType string `json:"type"`
	Details string `json:"details"`
}

type Person struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Birthdate time.Time `json:"birthdate"`
	Contacts []*Contact `json:"contacts"`
}
