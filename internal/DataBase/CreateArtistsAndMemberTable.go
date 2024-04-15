package DataBase

import "fmt"

// -------------------------------------------
func CreateArtistsAndMemberTable() error {
	// ---------------------------------------
	// the structure as helper to read json data
	type ArtistsJson []struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations    string   `json:"locations"`
		ConcertDates string   `json:"concertDates"`
		Relations    string   `json:"relations"`
	}
	//---------------------------------------------
	var myArtistsJsonContent ArtistsJson
	//-----------------------------
	// fetch UrlArtists api and save it in  myArtistsJsonContent
	errfetchArtistsUrl := fetchUrl(myUrlTable.UrlArtists, &myArtistsJsonContent)
	if errfetchArtistsUrl != nil {
		errMsg = fmt.Sprintf("\n<---CreateArtistsAndMemberTable<---fetchUrl<---[%v]", errfetchArtistsUrl)
		logMsg = "error in fetching tha Artists Url" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)
	}
	//----------------------------------
	var myArtistData Artist
	//---------initiate empty map for Tables------------------
	Mydb.TableOfArtists = make(map[int]Artist)
	Mydb.TableOfMembers = make(map[int][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myArtistsJsonContent {

		myArtistData = Artist{
			ID:           item.ID,
			Name:         item.Name,
			ImageUrl:     item.Image,
			CreationDate: item.CreationDate,
			FirstAlbum:   item.FirstAlbum,
		}

		Mydb.TableOfArtists[item.ID] = myArtistData
		Mydb.TableOfMembers[item.ID] = item.Members
	}
	//----------------------------------
	return nil
}

//--------------------------------
