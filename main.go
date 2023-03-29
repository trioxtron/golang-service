package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type dogImage struct {
    Message string `json:"message"`
    Status  string `json:"status"`
}

func fetch(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return res
}


func main() {
    var fetchDog dogImage
    
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        fetchRes := fetch("https://dog.ceo/api/breeds/image/random").Body
        body, err := io.ReadAll(fetchRes)
        if err != nil {
            log.Fatalln(err)
        }

        if err :=json.Unmarshal(body, &fetchDog); err != nil {
            log.Fatalln(err)
        }
        return c.JSON(fetchDog)
    })

    app.Listen(":3000")

}
