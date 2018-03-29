package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume1/section5/socialmedia"
)

func displaySocialMediaPostHandler(w http.ResponseWriter, r *http.Request){
	myPost := socialmedia.NewPost("Name", socialmedia.Moods["thrilled"], "Go is neat!", "Checkot the Go website!", "https://golang.org", "./images/gogopher.png", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
	renderTemplate(w, "./templates/socialMediaPostTemplate.html", myPost)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}){
	t, err := template.ParseFiles(templateFile)
	if err != nil{
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func main(){
	http.HandleFunc("/display-social-media-post", displaySocialMediaPostHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}