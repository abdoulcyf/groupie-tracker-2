package servers

import (
	"fmt"
	"net/http"
	"os"
)

func FileServer(directoryPath string) (string, error) {
	//directoryPath := "../../static"
	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		fmt.Printf("\n<---FileServer<---os.Stat<---[%v]", err)
	}
	fileServer := http.FileServer(http.Dir(directoryPath))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	return directoryPath, nil
}
