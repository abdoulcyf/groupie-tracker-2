package DataBase

import "fmt"

func createLocationTable() error {
	//--------------------------------

	// the structure of locations.json
	type LocationsJson struct {
		Index []struct {
			ID        int      `json:"id"`
			Locations []string `json:"locations"`
			Dates     string   `json:"dates"`
		} `json:"index"`
	}

	//--------------------------------

	var myLocationJsonContent LocationsJson
	//--------------------------------
	// fetch UrlLocations api and save it in  myLocationJsonContent
	errfetchLocationsUrl := fetchUrl(myUrlTable.UrlLocations, &myLocationJsonContent)
	if errfetchLocationsUrl != nil {
		errMsg = fmt.Sprintf("\n<---createLocationTable<---fetchUrl<---[%v]", errfetchLocationsUrl)
		logMsg = "error in fetching tha Locations Url" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)
	}
	//---------------------------------
	Mydb.TableOfLocations = make(map[int][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myLocationJsonContent.Index {
		Mydb.TableOfLocations[item.ID] = item.Locations
	}
	//-------------------------------
	return nil
}
