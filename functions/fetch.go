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

type Location struct {
	ID    int    `json:"id"`
	Name  string `json:"locations"`
	Dates string `json:"dates"`
}

type Date struct {
	ID   int    `json:"id"`
	Date string `json:"dates"`
}

type Relation struct {
	ID       int      `json:"id"`
	DateLocs []string `json:"datesLocations"`
}

type Data struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

var (
	artists   []Artist
	locations []Location
	dates     []Date
	relations []Relation
)

var apiURL = "https://groupietrackers.herokuapp.com/api/"

// FetchData fetches data from the API and stores it in the Data struct
// func FetchData() (*Data, error) {
// 	data := &Data{}

// 	endpoints := map[string]interface{}{
// 		"artists":   &data.Artists,
// 		"locations": &data.Locations,
// 		"dates":     &data.Dates,
// 		"relations": &data.Relations,
// 	}

// 	for endpoint, target := range endpoints {
// 		url := apiURL + endpoint
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			return nil, fmt.Errorf("failed to fetch %s: %s", endpoint, resp.Status)
// 		}

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if err := json.Unmarshal(body, target); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return data, nil
// }

func fetchData(url string, target interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second, // timeout the request if response is delayed
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("error fetching JSON from %s", url)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Reading from repsonse body")
		return err
	}

	return json.Unmarshal(body, &target)
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
	if err = fetchData("https://groupietrackers.herokuapp.com/api/relations", &relations); err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}
}
