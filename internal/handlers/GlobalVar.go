package handlers

import (
	"errors"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	ImageUrl string `json:"image"`
}

type Artist struct {
	ID           int
	Name         string
	ImageUrl     string
	CreationDate int
	FirstAlbum   string
	Members      []string
	Locations    []string // Add Locations field
	Dates        []string
	Relation     map[string][]string
}

type pageData struct {
	Type    string   // Indicates the type of data: "person" or "artist"
	//Persons []Artist // Slice of Person structs
	Artists []Artist // Slice of Artist structs
}

var ErrArtistNotFound = errors.New("artist not found")
