package main

import (
	"fmt"
	"main/internal/DataBase"
)

// ----------------------------------------------
func main() {
	//----------------------------------------
	errLoadDataBase := DataBase.LoadDataBase()
	if errLoadDataBase != nil {
		errMsg := fmt.Sprintf("\n<---main-Cli-App<---LoadDataBase<---[%v]", errLoadDataBase)
		logMsg := "error in fetching tha main Url" + errMsg
		//loader.Error(logMsg)
		DataBase.LoggerExt.Error(logMsg)
		fmt.Println(errMsg)
		return
	}
	//----------------------------------------
	DataBase.ReadDB(7)
	//----------------------------------------

}
