package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page represents an individul page on the website
type Page struct {
	Title string
	Body  []byte
}

func (p Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("File not found")
		// return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", r, "p.Body")
		}
	}()

	fmt.Println("Accessing Page File")
	p, _ := loadPage(title)
	fmt.Println("Accessed Page file")
	// if err != nil {
	// 	fmt.Println("Page not found so yeah")
	// }
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
