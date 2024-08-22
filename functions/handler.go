package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Response struct {
	pageTitle string
	Data      Data
}

var dataRes Response

// const apiURL = "https://groupietrackers.herokuapp.com/api"

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Println("OK: ", http.StatusOK)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataRes = Response{
		pageTitle: "Artists",
		Data: Data{
			Artists:   artists,
			Locations: locations.Index,
			Dates:     dates.Index,
			Relations: relations.Index,
		},
	}
	tmpl.Execute(w, dataRes)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if strings.ToUpper(r.Method) != http.MethodGet {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	// if err := fetchData(apiURL+"/artists", &artists); err != nil {
	// 	ServeError(w, fmt.Sprintf("Failed to fetch artists' data: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataRes = Response{
		pageTitle: "Artists",
		Data: Data{
			Artists:   artists,
			Locations: locations.Index,
			Dates:     dates.Index,
			Relations: relations.Index,
		},
	}
	tmpl.Execute(w, dataRes)
}

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/about.html")
		if err != nil {
			ServeError(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Concerts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/concerts" {
		ServeError(w, "Page not found", http.StatusNotFound)
	}

	if strings.ToUpper(r.Method) != http.MethodGet {
		ServeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// tmpl, err := template.ParseFiles("/templates/concerts.html")
	// if err != nil {
	// 	fmt.Println(err)
	// 	ServeError(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }
}
