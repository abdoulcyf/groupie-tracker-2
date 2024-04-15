package servers

import (
	//"main/internal/art"
	"fmt"
	"net/http"
)

func WebServer() error {

	//-------------------------------
	directoryPath := "../../static"

	_, err := FileServer(directoryPath)
	if err != nil {
		fmt.Println("Directory path does not exist")
	} else {
		fmt.Println("Directory path received succesfully")
	}
	//-------------------------------

	serverWebGuiApp := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/", MainHandler)

	errForServer := serverWebGuiApp.ListenAndServe()

	if errForServer != nil {
		errMsg := fmt.Sprintf("\n<---WebServer<---ListenAndServe<---[%v]", errForServer)
		return fmt.Errorf(errMsg)
	} else {
		fmt.Println("Server is running on port 8080")
	}

	//-------------------------------
	return nil
}

/*
{{.Domain}}
"http://localhost:8080/ascii-art"
*/
