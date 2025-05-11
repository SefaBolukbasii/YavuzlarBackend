package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func WriteJson(c *fiber.Ctx, v any) {
	c.Set("Content-Type", "application/json")

	data, err := json.Marshal(v)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Error encoding JSON")
		return
	}

	c.Send(data)
}
