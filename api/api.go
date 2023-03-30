package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type dogs struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
type catfacts struct {
	Fact   string `json:"fact"`
	Length string `json:"length"`
}

var APIs = map[string]string{
	"catfacts": "https://catfact.ninja/fact",
	"dogs":     "https://dog.ceo/api/breeds/image/random",
	"boredapi": "https://www.boredapi.com/api/activity",
}

func fetch(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func GetApis(c *fiber.Ctx) error {
	var apistring string
	for key := range APIs {
		apistring += key
		apistring += " "
	}

	return c.SendString(apistring)
}

func GetApi(c *fiber.Ctx) error {
	var fetchDog dogs

	usedAPI, ok := APIs[c.Params("api")]
	if !ok {
		return c.SendString("api not available")
	}
	fetchRes := fetch(usedAPI)

	if err := json.Unmarshal(fetchRes, &fetchDog); err != nil {
		log.Fatalln(err)
	}

	return c.JSON(fetchDog)
}
