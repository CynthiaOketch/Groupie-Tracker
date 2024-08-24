package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Define structs for API response
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Index []Location
}
type Location struct {
	ID    int      `json:"id"`
	Venues  []string `json:"locations"`
	Dates string   `json:"dates"`
}

type Dates struct {
	Index []Date
}

type Date struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}

// type ConcertDate struct {
// 	ID   int       `json:"id"`
// 	Date time.Time `json:"date"`
// }

// type ConcertDates struct {
// 	Index []ConcertDate
// }

type Relations struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID       int      `json:"id"`
	DateLocs map[string][]string `json:"datesLocations"`
}

// type DateLoc struct {
// 	Location string
// 	Dates []time.Time
// }

type Data struct {
	Artists      []Artist
	Locations    Locations
	Dates        Dates
	// ConcertDates ConcertDates
	Relations    Relations
}

type BandDetails struct {
	Artist Artist
	Location Location
	Dates Date
	Relation Relation
}

var (
	artists      []Artist
	locations    Locations
	dates        Dates
	// concertDates ConcertDates
	relations    Relations
)

var apiURL = "https://groupietrackers.herokuapp.com/api/"

// FetchData fetches data from the API and stores it in the Data struct
func fetchData(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func loadData() {
	var err error
	if err = fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	if err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations); err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}
	if err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates); err != nil {
		log.Fatalf("Error fetching dates: %v", err)
	}
	if err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations); err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}
}
