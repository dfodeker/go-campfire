package endpoints

// import (
// 	"html/template"
// 	"io"

// 	"github.com/labstack/echo/v4"
// )

// type Template struct {
// 	templates *template.Template
// }
// type TemplateRenderer struct {
// 	templates *template.Template
// }

// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

// 	// Add global methods if data is a map
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}

// 	return t.templates.ExecuteTemplate(w, name, data)
// }

// // func JoinStrings(sep string, items []string) string {
// // 	return strings.Join(items, sep)
// // }

// // NewTemplateRenderer initializes a new TemplateRenderer
// func NewTemplateRenderer() *TemplateRenderer {

// 	tmpl := template.Must(template.ParseFiles(
// 		"./templates/index.html",
// 		"./templates/camp.html",
// 	))

// 	return &TemplateRenderer{
// 		templates: tmpl,
// 	}
// }

// // Render renders a template with data

// func HandleIndex(c echo.Context) error {
// 	return c.Render(200, "index.html", nil)
// }
