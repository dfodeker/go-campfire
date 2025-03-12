package main

import (
	"fmt"

	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	Template interface {
		ExecuteTemplate(wr io.Writer, name string, data any) error
	}
}
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Camp struct {
	Name       string
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

var Camp1 = &Camp{
	Name:       "Camp Halfblood",
	Attributes: []string{"Strong"},
}

var camps = []Camp{
	{
		Name:       "Camp Halfblood",
		Attributes: []string{"Greek", "Montauk", "Strong"},
		Campers: []Camper{
			{Name: "Percy Jackson", Age: 12},
			{Name: "Annabeth Chase", Age: 13},
		},
	},
	{
		Name:       "Camp Jupiter",
		Attributes: []string{"Roman", "San Francisco", "Disciplined"},
		Campers: []Camper{
			{Name: "Jason Grace", Age: 16},
			{Name: "Piper McLean", Age: 15},
		},
	},
}

func joinStrings(sep string, items []string) string {
	return strings.Join(items, sep)
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {

	fmt.Println("=============================")

	if Camp1.Name == "Camp Jupiter" {
		Camp1.Attributes = append(Camp1.Attributes, "The Legion is the core military unit of Camp Jupiter, with different cohorts representing various gods.")
		Camp1.Attributes = append(Camp1.Attributes, "The Fort, where campers train and live.")
		Camp1.Attributes = append(Camp1.Attributes, "The Augury, a priestess who predicts the future.")
	} else {
		Camp1.Attributes = append(Camp1.Attributes, "Magic Barrier")
		Camp1.Attributes = append(Camp1.Attributes, "The Arena, where campers train in combat.")
		Camp1.Attributes = append(Camp1.Attributes, "The Campfire, where stories and songs are shared.")
	}

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define route
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// something??
	renderer := &echo.TemplateRenderer{
		Template: template.Must(template.New("").Funcs(template.FuncMap{
			"join": joinStrings, // Register join function
		}).ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer
	// Routes
	// e.GET("/", hello)
	e.GET("/", func(c echo.Context) error {
		// camp := &Camp{
		// 	Name:       "Camp Halfblood",
		// 	Attributes: []string{"Greek", "Montauk", "Strong"},
		// }
		return c.Render(http.StatusOK, "index.html", camps)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
