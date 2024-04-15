package DataBase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ----------------------------------------------
func fetchUrl(url string, structure interface{}) error {
	//--------GET API CALL----------------
	resp, errGet := http.Get(url)
	//-------------------------------------
	if errGet != nil {
		errMsg = fmt.Sprintf("\n<---readArtistsAndMemberjson<---http.Get<---[%v]", errGet)
		logMsg = "myUrlTable.UrlArtists url server not reachable " + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	defer resp.Body.Close()
	//-------------------------------------
	// Check response status
	if resp.StatusCode != http.StatusOK {
		errMsg = fmt.Sprintf("\n<---readArtistsAndMemberjson<---resp.StatusCode<---[%v]", resp.StatusCode)
		logMsg = "myUrlTable.UrlArtists url server return not 200 status code " + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)
	}
	//----------------read response body--------------------
	contentByte, errByte := io.ReadAll(resp.Body)
	if errByte != nil {
		errMsg = fmt.Sprintf("\n<---readArtistsAndMemberjson<---io.ReadAll<---[%v]", errByte)
		logMsg = "myUrlTable.UrlArtists can not read body of response " + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)

	}
	//--------------Unmarshal----------
	errUnmarshal := json.Unmarshal(contentByte, &structure)
	if errUnmarshal != nil {
		errMsg = fmt.Sprintf("\n<---readArtistsAndMemberjson<---Unmarshal<---[%v]", errUnmarshal)
		logMsg = "myUrlTable.UrlArtists can not Unmarshal json " + errMsg
		logger.Error(logMsg)
		return fmt.Errorf(errMsg)
	}
	//--------------------------
	return nil
}
