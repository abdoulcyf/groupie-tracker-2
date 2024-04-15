package DataBase

import "fmt"

func LoadDataBase() error {
	//---------------fetchUrl---------------
	// fetching mainApiUrl api and save the datas in a var
	errfetchmainUrl := fetchUrl(mainApiUrl, &myUrlTable)
	if errfetchmainUrl != nil {
		errMsg = fmt.Sprintf("\n<---LoadDataBase<---mainApiUrl<---[%v]", errfetchmainUrl)
		logMsg = "error in fetching tha main Url" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	//------------CreateArtistsAndMemberTable---------------
	// create database Artists table and members table
	errCreatingArtistsAndMemberTable := CreateArtistsAndMemberTable()
	if errCreatingArtistsAndMemberTable != nil {
		errMsg = fmt.Sprintf("\n<---LoadDataBase<---CreateArtistsAndMemberTable<---[%v]", errCreatingArtistsAndMemberTable)
		logMsg = "error in creating Artists and member Table" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	//--------------createLocationTable-------------
	//create Location table
	errCreateLocationTable := createLocationTable()
	if errCreateLocationTable != nil {
		errMsg = fmt.Sprintf("\n<---LoadDataBase<---createLocationTable<---[%v]", errCreateLocationTable)
		logMsg = "error in creating Locations Table" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	//-----------------createDatesTable-------------
	//create datas table
	erCreateDatesTable := createDatesTable()
	if erCreateDatesTable != nil {
		errMsg = fmt.Sprintf("\n<---LoadDataBase<---createDatesTable<---[%v]", erCreateDatesTable)
		logMsg = "error in creating Dates Table" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	//---------createRelationsTable---------------
	//create Relations table
	errCreateDatesTable := createRelationsTable()
	if errCreateDatesTable != nil {
		errMsg = fmt.Sprintf("\n<---LoadDataBase<---createRelationsTable<---[%v]", errCreateDatesTable)
		logMsg = "error in creating Relations Table" + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}

	//------------------------------------------
	return nil
}
