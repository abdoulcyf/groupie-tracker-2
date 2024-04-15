package DataBase

import (
	"fmt"
)

func createRelationsTable() error {
	//----------------------------------------
	// the structure of relation.json
	type RelationsJson struct {
		Index []struct {
			ID             int `json:"id"`
			DatesLocations map[string][]string
		} `json:"index"`
	}
	//-------------------------------------------
	var myRelationsJsonContent RelationsJson
	//-------------------------------------------------
	// fetch UrlRelation api and save it in  myRelationsJsonContent
	errfetchRelationUrl := fetchUrl(myUrlTable.UrlRelation, &myRelationsJsonContent)
	if errfetchRelationUrl != nil {
		errMsg = fmt.Sprintf("\n<---createRelationsTable<---fetchUrl<---[%v]", errfetchRelationUrl)
		logMsg = "error in fetching tha Relation Url" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	//--------------------------------------
	//initiate the map (create an empty map)
	Mydb.TableOfRelation = make(map[int]map[string][]string)
	//------Read data from content and put them in Table-------

	for _, item := range myRelationsJsonContent.Index {

		Mydb.TableOfRelation[item.ID] = item.DatesLocations

	}
	//----------------------------------
	return nil
}
