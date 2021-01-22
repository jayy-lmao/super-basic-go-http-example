package main

import (
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/standard"
	"net/http"
	"strconv"
)

type (
	Snack struct {
		Id   int    `json: "snackId"`
		Name string `json:"name" query:"name"`
	}
)

var (
	snacks = map[int]*Snack{}
	seq    = 1
)

func getSnacks(c echo.Context) error {
  return c.JSON(snacks)
}

func createSnack(c echo.Context) error {
	// Create an ID
	s := &Snack{
		Id: seq,
	}

	// Add other properties
	if err := c.Bind(s); err != nil {
		return err
	}

	snacks[s.Id] = s
	seq++

	return c.JSON(s, http.StatusCreated)
}

func deleteSnack(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	delete(snacks, id)
	if err != nil {
		return c.JSON(err, http.StatusBadRequest)
	}
	return c.NoContent(http.StatusNoContent)
}

func getSnack(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(err, http.StatusBadRequest)
	}
  snack := snacks[id]

  if snack == nil {
    return c.NoContent(http.StatusNotFound)
  }

	return c.JSON(snacks[id], http.StatusOK)
}

func main() {
	e := echo.New()

	e.Get("/", func(c echo.Context) error {
		return c.String("Hello, World!", http.StatusOK)
	})
	e.Get("/snacks", getSnacks)
	e.Get("/snacks/:id", getSnack)
	e.Delete("/snacks/:id", deleteSnack)
	e.Post("/snacks", createSnack)

	e.Run(standard.New(":1323"))
}
