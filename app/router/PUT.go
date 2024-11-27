package router

import (
	"api-gateway-study/app/client"
	"api-gateway-study/config"

	"github.com/gofiber/fiber/v2"
)

type put struct {
	cfg    config.Router
	client *client.HttpClient
}

func AddPut(
	cfg config.Router,
	client *client.HttpClient,
) func(c *fiber.Ctx) error {
	r := put{cfg: cfg, client: client}
	return r.handleRequest
}

func (r put) handleRequest(c *fiber.Ctx) error {
	apiResut, err := r.client.DELETE(r.cfg.Path, r.cfg, c.Request().Body())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(apiResut)
}
