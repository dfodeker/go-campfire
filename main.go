package main

import (
	"fmt"
	"log"

	"html/template"

	"net/http"
	"strings"
)

type Camp struct {
	Name       string
	Photo      string
	Attributes []string
	Campers    []Camper
}
type Camper struct {
	Name string
	Age  int
}

func (c *Camp) UpdateCamp(newCamp string) {
	c.Name = newCamp
}
func (c Camp) Join(arr []string) string {
	return strings.Join(arr, ", ")
}

var Camp1 = &Camp{
	Name:       "Camp Halfblood",
	Attributes: []string{"Strong"},
}

var camps = []Camp{
	{
		Name:       "Camp Halfblood",
		Attributes: []string{"Greek", "Montauk", "Strong"},
		Photo:      "https://wccls.bibliocommons.com/events/uploads/images/full/ae2133848439b0be478b4c4ecfcb51f6/CampHalf-Blood.png",
		Campers: []Camper{
			{Name: "Percy Jackson", Age: 12},
			{Name: "Annabeth Chase", Age: 13},
		},
	},
	{
		Name:       "Camp Jupiter",
		Attributes: []string{"Roman", "San Francisco", "Disciplined"},
		Photo:      "https://image-cdn.neatoshop.com/styleimg/163827/none/purple/default/587347-20;1728171486y.jpg",
		Campers: []Camper{
			{Name: "Jason Grace", Age: 16},
			{Name: "Piper McLean", Age: 15},
		},
	},
}

func JoinStrings(sep string, items []string) string {
	return strings.Join(items, sep)
}

var campsTemplate *template.Template
var campersTemplate *template.Template

func main() {

	fmt.Println("=============================")

	_, err := template.ParseFiles(
		"./templates/home.html",
		"./templates/camp.html",
	)
	if err != nil {
		log.Fatalf("could not init templates %v\n", err)

	}
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"join":  strings.Join, // Converts a string to uppercase.
	}

	campsTemplate, err = template.New("layout.html").Funcs(funcMap).ParseFiles("templates/layout.html", "templates/camps.html")
	if err != nil {
		log.Fatalf("Error parsing camps template: %v", err)
	}

	// Template for campers page
	campersTemplate, err = template.New("layout.html").Funcs(funcMap).ParseFiles("templates/layout.html", "templates/campers.html")
	if err != nil {
		log.Fatalf("Error parsing campers template: %v", err)
	}

	http.HandleFunc("/", campsHandler)
	http.HandleFunc("/camp/", campersHandler) // expects URL of the form /camp/{name}

	log.Println("Server starting on :1230")
	if err := http.ListenAndServe(":1230", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
func campsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Camps []Camp
	}{
		Title: "List of Camps",
		Camps: camps,
	}

	if err := campsTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// campersHandler renders a list of campers for a specific camp.
// It extracts the camp name from the URL.
func campersHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the camp name from the URL path. For example, "/camp/CampAdventure"
	path := strings.TrimPrefix(r.URL.Path, "/camp/")
	if path == "" {
		http.Error(w, "Camp name not specified", http.StatusBadRequest)
		return
	}
	campName := path
	var camp *Camp
	for i := range camps {
		if camps[i].Name == campName {
			camp = &camps[i]
			break
		}
	}
	data := struct {
		Title    string
		CampName string
		Camp     *Camp
	}{
		Title:    "Campers at " + campName,
		CampName: campName,
		Camp:     camp,
	}

	if err := campersTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
