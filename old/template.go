package main

import (
	"html/template"
	"log"
	"net/http"
)

// used to pass to template
// name corosponds to the name on the idnex
type Gopher struct{
	Name string
}


func helloGopherHandler (w http.ResponseWriter, r *http.Request){
	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == ""{
		gophername = "Gopher"
	}
	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./index.html", gopher)
}


func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil{
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

//basic server and handle call
func main(){
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.ListenAndServe(":8080",nil)
}
