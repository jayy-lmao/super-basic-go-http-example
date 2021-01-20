package main

import (
  "net/http"
  "strconv"
	"github.com/webx-top/echo"
  "github.com/webx-top/echo/engine/standard"
)

type (
  Snack struct {
    Id int `json: "snackId"`
    Name string `json:"name" query:"name"`
};
)

var (
	snacks = map[int]*Snack{}
	seq   = 1
)
func createSnack(c echo.Context) error {
    // Create an ID
    s := &Snack{
      Id: seq,
    }

    // Add other properties
    if err := c.Bind(s);err!=nil{
      return err
    }

    snacks[s.Id] = s
    seq++

    return  c.JSON(s,http.StatusCreated)
}

func getSnack(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(err, http.StatusBadRequest)
  }
  return c.JSON(snacks[id], http.StatusOK)
}


func main() {
	e := echo.New()

	e.Get("/", func(c echo.Context) error {
		return c.String("Hello, World!", http.StatusOK)
	})
  e.Get("/snacks/:id", getSnack)
  e.Post("/snacks", createSnack)

  e.Run(standard.New(":1323"));
}
