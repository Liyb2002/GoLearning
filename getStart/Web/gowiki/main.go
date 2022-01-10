package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
//	"github.com/gin-gonic/gin"


)


type Page struct {
    Title  string 
    Body []byte  
}

//Local save
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    return &Page{Title: title, Body: body}, nil
}


//web functions

//Get

func viewHandler(w http.ResponseWriter, r *http.Request) {

	//r.URL.Path = /view/TestPage

    title := r.URL.Path[len("/view/"):]
  	p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
//	fmt.Fprintf(w, "hi %s", r.URL.Path)
}




func main() {

    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}

