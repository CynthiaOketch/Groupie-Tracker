package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Artist struct {
	ID           int            `json:"id"`
	Image        string         `json:"image"`
	Name         string         `json:"name"`
	Members      []string `json:"members"`
	CreationDate int            `json:"creationDate"`
	FirstAlbum   string         `json:"firstAlbum"`
	Locations    string         `json:"locations"`
	ConcertDates string         `json:"concertDates"`
	Relations    string         `json:"relations"`
}

type Artists []Artist

var data Artists

func fetchJSON() {
	client := &http.Client{
		Timeout: 10 * time.Second, // timeout the request if response is delayed
	}
	resp, err := client.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("error fetching JSON")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error : Received status code %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Reading from repsonse body")
		return
	}
	
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON body", err)
		fmt.Println("body: ", string(body))
		return
	}
	fmt.Println("fetched data: ", data)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	fetchJSON()
	
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home")
}

func handlerFunctions() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artists", artistHandler)
	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	handlerFunctions()
}
