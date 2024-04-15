package handlers

import (
	// "fmt"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, errMsg string) {

	//templateFileAddress := "../../templates/"
	templateFileName := "error.html"
	templateFile := templateFileAddress + templateFileName
	// Define the template
	t := template.Must(template.ParseFiles(templateFile))
	myUrl := myDomain + "/"
	data := struct {
		Msg string
		Url string
	}{
		Msg: errMsg,
		Url: myUrl,
	}
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
