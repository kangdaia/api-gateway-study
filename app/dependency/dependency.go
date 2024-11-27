package dependency

import (
	"api-gateway-study/app/client"
	"api-gateway-study/app/router"
	"api-gateway-study/config"
	"api-gateway-study/kafka"
	"flag"

	"go.uber.org/fx"
)

var (
	yamlPath = flag.String("yamlPath", "./deploy.yaml", "yaml path")
)

func init() {
	flag.Parse()
}

var Cfg = fx.Module(
	"config",
	fx.Provide(func() config.Config {
		return config.NewCfg(*yamlPath)
	}),
)

var Producer = fx.Module(
	"kafka_producer",
	fx.Provide(func(cfg config.Config) map[string]kafka.Producer {
		clients := make(map[string]kafka.Producer, len(cfg.App))
		for _, a := range cfg.App {
			clients[a.App.Name] = kafka.NewProducer(a.Producer)
		}
		return clients
	}),
)

var HttpClient = fx.Module(
	"http_client",
	fx.Provide(func(cfg config.Config, producer map[string]kafka.Producer) map[string]*client.HttpClient {
		clients := make(map[string]*client.HttpClient, len(cfg.App))
		for _, a := range cfg.App {
			clients[a.App.Name] = client.NewHttpClient(a, producer)
		}
		return clients
	}),
)

var Router = fx.Module(
	"router",
	fx.Provide(func(cfg config.Config, client map[string]*client.HttpClient) map[string]router.Router {
		clients := make(map[string]router.Router, len(cfg.App))
		for _, a := range cfg.App {
			clients[a.App.Name] = router.NewRouter(a, client)
		}
		return clients
	}),
)
