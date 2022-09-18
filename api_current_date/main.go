package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const defaulTimeFormat = "2006-01-02 15:04:05 -0700 MST"

type DefaulTimeResponse struct {
	CurrentTime string `json:"current_time"`
}

type ManyTZResponse map[string]string

func main() {
	app := fiber.New()

	app.Get("/api/time", func(c *fiber.Ctx) error {
		timezone := c.Query("tz")

		if timezone == "" {
			return c.JSON(DefaulTimeResponse{
				CurrentTime: time.Now().UTC().Format(defaulTimeFormat),
			})
		}

		manyTZResponse := ManyTZResponse{}
		for _, tz := range strings.Split(timezone, ",") {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(fmt.Sprintf("invalid timezone: %s", tz))
			}

			manyTZResponse[tz] = time.Now().In(loc).Format(defaulTimeFormat)
		}

		return c.JSON(manyTZResponse)
	})

	app.Listen(":3333")
}
