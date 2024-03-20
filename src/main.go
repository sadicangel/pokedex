package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Pokemon struct {
	Id    int
	Name  string
	Types []string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	var pokemons = []Pokemon{
		{
			Id:    1,
			Name:  "Bulbasaur",
			Types: []string{"Grass", "Poison"},
		},
		{
			Id:    4,
			Name:  "Charmander",
			Types: []string{"Fire"},
		},
		{
			Id:    7,
			Name:  "Squirtle",
			Types: []string{"Water"},
		},
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", pokemons)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
