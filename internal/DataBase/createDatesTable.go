package DataBase

import "fmt"

func createDatesTable() error {
	// -----------------------------------------
	// the structure of Dates.json
	type DatesJson struct {
		Index []struct {
			ID    int      `json:"id"`
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	//---------------------------------------
	var myDatesJsonContent DatesJson
	//---------------------------------------
	// fetch UrlDates api and save it in  myDatesJsonContent
	errfetchDatesUrl := fetchUrl(myUrlTable.UrlDates, &myDatesJsonContent)
	if errfetchDatesUrl != nil {
		errMsg = fmt.Sprintf("\n<---createDatesTable<---fetchUrl<---[%v]", errfetchDatesUrl)
		logMsg = "error in fetching tha Dates Url" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)
	}
	//------initiate map for Table OF data-----------
	Mydb.TableOfDates = make(map[int][]string)
	//------Read data from content and put them in Table-------
	for _, item := range myDatesJsonContent.Index {
		Mydb.TableOfDates[item.ID] = item.Dates
	}
	//---------------------------------------
	return nil
}
