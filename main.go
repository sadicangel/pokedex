package main

import (
	"encoding/json"
	"fmt"
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

const MAX = 151

type Pokemon struct {
	Id     int      `json: "id"`
	Name   string   `json: "name"`
	Types  []string `json: "types"`
	Sprite string   `json: "sprite`
}

type PokemonStats struct {
	HP        int `json: "HP"`
	Attack    int `json: "Attack"`
	Defense   int `json: "Defense"`
	SpAttack  int `json: "SpAttack"`
	SpDefense int `json: "SpDefense"`
	Speed     int `json: "Speed"`
}

type PokemonDetails struct {
	Id    int          `json: "id"`
	Name  string       `json: "name"`
	Types []string     `json: "types"`
	Stats PokemonStats `json: "stats"`
	Image string       `json: "image`
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

	return pokemonList[0:MAX]
}

func readPokemonDetails() map[string]PokemonDetails {
	var pokemonList []PokemonDetails
	pokemonJson, err := os.Open("data/pokemon_details.json")
	if err != nil {
		panic(err)
	}
	defer pokemonJson.Close()

	pokemonBytes, _ := io.ReadAll(pokemonJson)

	json.Unmarshal(pokemonBytes, &pokemonList)

	details := make(map[string]PokemonDetails)
	for i := 0; i < MAX; i++ {
		details[fmt.Sprint(pokemonList[i].Id)] = pokemonList[i]
	}
	return details
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Static("/public", "public")

	pokemonList := readPokemonList()
	pokemonDetails := readPokemonDetails()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", pokemonList)
	})
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.Render(200, "details", pokemonDetails[id])
	})
	e.Logger.Fatal(e.Start(":3000"))
}
