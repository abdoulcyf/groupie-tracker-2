package handlers

import (
	"fmt"
	"html/template"
	"main/internal/DataBase"
	"net/http"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	//------------------------------------
	errForDataBase := DataBase.LoadDataBase()
	if errForDataBase != nil {
		fmt.Println("Error in loading data base")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Extract artist name from URL
	artistName := strings.TrimPrefix(r.URL.Path, "/artist/")

	// Retrieve artist details from the database based on the name
	artist, found := FindArtistByName(artistName)
	if !found {
		http.NotFound(w, r)
		return
	}

	data := pageData{
		Type:    "artist",
		Artists: []Artist{artist},
	}

	// Parse the template
	templateAddress := "../../templates/artist_details.html"
	temp, err := template.ParseFiles(templateAddress)
	if err != nil {
		fmt.Println("Error in parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data
	err = temp.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func FindArtistByName(name string) (Artist, bool) {
	for _, artist := range DataBase.Mydb.TableOfArtists {
		if artist.Name == name {
			// Retrieve additional information associated with the artist
			members, found := DataBase.Mydb.TableOfMembers[artist.ID]
			if !found {
				members = []string{} // Set an empty slice if no members found
			}

			locations, found := DataBase.Mydb.TableOfLocations[artist.ID]
			if !found {
				locations = []string{} // Set an empty slice if no locations found
			}

			dates, found := DataBase.Mydb.TableOfDates[artist.ID]
			if !found {
				dates = []string{} // Set an empty slice if no dates found
			}

			relation, found := DataBase.Mydb.TableOfRelation[artist.ID]
			if !found {
				relation = make(map[string][]string) // Set an empty map if no relation found
			}

			// Populate the fields of the Artist struct
			artist.Members = members
			artist.Locations = locations
			artist.Dates = dates
			artist.Relation = relation

			return Artist(artist), true
		}
	}
	return Artist{}, false
}
