package main

import (
	"encoding/json"
	"html/template"
	"io"
	"os"

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
	Id     int      `json: "id"`
	Name   string   `json: "name"`
	Types  []string `json: "types"`
	Sprite string   `json: "sprite`
}

func readPokemonList() []Pokemon {
	var pokemonList []Pokemon
	pokemonJson, err := os.Open("data/pokemon.json")
	if err != nil {
		panic(err)
	}
	defer pokemonJson.Close()

	pokemonBytes, _ := io.ReadAll(pokemonJson)

	json.Unmarshal(pokemonBytes, &pokemonList)

	return pokemonList[0:151]
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Static("/", "public")

	pokemonList := readPokemonList()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", pokemonList)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
