package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Person is a person
type Person struct {
	Name       string `json: "name"`
	Age        int    `json: "age"`
	Profession string `json: "profession"`
	HairColor  string `json: "hairColor"`
}

var peopleMap = make(map[string]Person)

func main() {
	http.HandleFunc("/people", People)
	http.HandleFunc("/people/", SearchPeople)
	http.ListenAndServe(":8080", nil)
}

// SearchPeople searches the map for a person
func SearchPeople(w http.ResponseWriter, req *http.Request) {
	var searchQuery = req.URL.Path[8:]
	searchResult, prs := peopleMap[searchQuery]

	if prs != true {
		fmt.Fprintf(w, `A person with the name "%s" was not found!`, searchQuery)
		return
	}
	fmt.Fprintf(w, "%+v", searchResult)
}

// People Prints a message
func People(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		log.Println(req.URL.Path[1])
		if req.URL.Path == "/people" {
			e, err := json.Marshal(peopleMap)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Write all the people to people.json
			err = ioutil.WriteFile("people.json", e, 0644)
			if err != nil {
				log.Fatal(err)
			}

			// Show all the people on the webpage
			fmt.Fprintf(w, "%s", e)
		}

	case "POST":
		// Decode the body of the POST req
		decoder := json.NewDecoder(req.Body)
		var newPerson Person
		err := decoder.Decode(&newPerson)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		peopleMap[newPerson.Name] = newPerson
		fmt.Fprintf(w, "The new person has been added to the map!")
	}
}
