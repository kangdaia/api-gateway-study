package router

import (
	"api-gateway-study/app/client"
	"api-gateway-study/config"
	"api-gateway-study/types/http"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type get struct {
	cfg    config.Router
	client *client.HttpClient
}

func AddGet(
	cfg config.Router,
	client *client.HttpClient,
) func(c *fiber.Ctx) error {
	r := get{cfg: cfg, client: client}
	return r.handleRequest
}

func (r get) handleRequest(c *fiber.Ctx) error {
	switch r.cfg.GetType {
	case http.QUERY:
		return r.queryType(c)
	case http.URL:
		return r.urlType(c)
	default:
		panic("Failed to find get type.")
	}
}

func (r get) queryType(c *fiber.Ctx) error {
	var builder strings.Builder
	builder.WriteString(r.cfg.Path)

	for i, k := range r.cfg.Variable {
		v := utils.CopyString(c.Query(k))

		if i != 0 {
			builder.WriteString(fmt.Sprintf("?%s=%s", k, v))
		} else {
			builder.WriteString(fmt.Sprintf("&%s=%s", k, v))
		}
	}
	fullUrl := builder.String()
	apiResult, err := r.client.GET(fullUrl, r.cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(apiResult)
}

func (r get) urlType(c *fiber.Ctx) error {
	var builder strings.Builder
	builder.WriteString(string(c.Request().URI().Path()))

	//routes := c.Route()
	//params := routes.Params
	//for i, p := range params {
	//	v := utils.CopyString(c.Params(p))
	//	builder.WriteString(v)
	//	if i < len(params)-1 {
	//		builder.WriteString("/")
	//	}
	//}

	fullUrl := builder.String()
	apiResult, err := r.client.GET(fullUrl, r.cfg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(apiResult)
}
