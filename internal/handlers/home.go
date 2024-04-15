package handlers

import (
	"fmt"
	"html/template"
	"main/internal/DataBase"
	"net/http"
)

// func HomeHandler(w http.ResponseWriter, r *http.Request) error {
// 	//------------------------------------
// 	errForDataBase := DataBase.LoadDataBase()
// 	if errForDataBase != nil {
// 		fmt.Println("Error in loading data base")
// 	} else {
// 		fmt.Println("Data base loaded successfully")
// 	}
// 	person1 := Person{
// 		Name:  DataBase.Mydb.TableOfArtists[1].Name,
// 		Image: DataBase.Mydb.TableOfArtists[1].ImageUrl,
// 	}

// 	person2 := Person{
// 		Name: DataBase.Mydb.TableOfArtists[2].Name,

// 		Image: DataBase.Mydb.TableOfArtists[2].ImageUrl,
// 	}

// 	person3 := Person{
// 		Name: DataBase.Mydb.TableOfArtists[3].Name,

// 		Image: DataBase.Mydb.TableOfArtists[3].ImageUrl,
// 	}

// 	var templateAdrress string = "../../templates/index.html"
// 	temp, err := template.ParseFiles(templateAdrress)
// 	if err != nil {
// 		fmt.Println("Error in parsing template")
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	} else {
// 		fmt.Println("Template parsed successfully")
// 	}

// 	data := pageData{
// 		Type:    "Person",
// 		Persons: []Person{person1, person2, person3},
// 		Artists: []Artist{},
// 	}

//		err = temp.Execute(w, data)
//		if err != nil {
//			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		}
//		return nil
//	}
func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	//------------------------------------
	errForDataBase := DataBase.LoadDataBase()
	if errForDataBase != nil {
		fmt.Println("Error in loading data base")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return errForDataBase
	}

	// Initialize an empty slice to hold persons
	var artists []Artist

	// Iterate over the artists in the database and populate the slice
	for _, artist := range DataBase.Mydb.TableOfArtists {
		artists = append(artists, Artist{
			ID:       artist.ID,
			Name:     artist.Name,
			ImageUrl: artist.ImageUrl,
		})
	}

	// Prepare data for the template
	data := pageData{
		Type:    "artist",
		Artists: artists,
		//Artists: nil, // We are not using the Artists field
	}

	// Parse the template
	templateAddress := "../../templates/index.html"
	temp, err := template.ParseFiles(templateAddress)
	if err != nil {
		fmt.Println("Error in parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	// Execute the template with the data
	err = temp.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	return nil
}
