package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	pageTitle string
	Data      Data
}

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

	data := Data{
		Artists:   artists,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}

	dataResponse := Response{
		pageTitle: "Groupie Tracker",
		Data:      data,
	}
	tmpl.Execute(w, dataResponse)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if strings.ToUpper(r.Method) != http.MethodGet {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := fetchData(apiURL+"/artists", &artists); err != nil {
		ServeError(w, fmt.Sprintf("Failed to fetch artists' data: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataRes := Response{
		pageTitle: "Artists",
		Data: Data{
			Artists:   artists,
			Locations: locations,
			Dates:     dates,
			Relations: relations,
		},
	}

	// fmt.Println(dataRes.Data)

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

	tmpl, err := template.ParseFiles("templates/concerts.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataConcerts := Response{
		pageTitle: "Concerts",
		Data: Data{
			Artists:   artists,
			Locations: locations,
			Dates:     dates,
			Relations: relations,
		},
	}

	tmpl.Execute(w, dataConcerts)
}

func ArtistDetail(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID from the URL (e.g., /artists/{id})
	idStr := strings.TrimPrefix(r.URL.Path, "/artists/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ServeError(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Find the artist by ID
	var artist Artist
	loadData()
	fmt.Println("Artists to search: ", id)
	fmt.Println("All artists: ", artists)
	for _, a := range artists {
		if a.ID == id {
			fmt.Println(artist)
			artist = a
			break
		}
	}
	fmt.Println("The ID: ", artist.ID)
	fmt.Println("The artist: ", artist)
	if artist.ID == 0 {
		ServeError(w, "Artist not found", http.StatusNotFound)
		return
	}

	var location []Location
	var artistDates []Date
	var artistConcertDates []ConcertDate
	var artistRelations []Relation

	for _, loc := range locations.Index {
		if loc.ID == id {
			location = append(location, loc)
		}
	}

	for _, date := range dates.Index {
		if date.ID == id {
			artistDates = append(artistDates, date)
		}
	}

	for _, concert := range concertDates.Index {
		if concert.ID == id {
			artistConcertDates = append(artistConcertDates, concert)
		}
	}

	for _, relation := range relations.Index {
		if relation.ID == id {
			artistRelations = append(artistRelations, relation)
		}
	}

	// Load the artist detail template
	tmpl, err := template.ParseFiles("templates/band.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Prepare the response data
	data := Response{
		pageTitle: artist.Name,
		Data: Data{
			Artists:      []Artist{artist},
			Locations:    Locations{Index: location},
			Dates:        Dates{Index: artistDates},
			ConcertDates: ConcertDates{Index: artistConcertDates},
			Relations:    Relations{Index: artistRelations},
		},
	}

	tmpl.Execute(w, data)
}
