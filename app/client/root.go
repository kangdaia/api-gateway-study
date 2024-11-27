package client

import (
	"api-gateway-study/common"
	"api-gateway-study/config"
	"api-gateway-study/kafka"
	"sync"

	"github.com/go-resty/resty/v2"
)

const (
	_defaultBatchTime = 2
)

type HttpClient struct {
	client *resty.Client
	cfg    config.App

	producer kafka.Producer

	batchTime    float64
	fetchLock    sync.Mutex // 추후 redis mutex lock 적용
	mapper       []ApiRequestTopic
	fetchChannel chan ApiRequestTopic
}

func NewHttpClient(
	cfg config.App,
	producer map[string]kafka.Producer,
) *HttpClient {
	batchTime := cfg.Producer.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	if cfg.Http.BaseUrl == "" {
		panic("base url does not exists.")
	}

	httpClient := HttpClient{
		cfg:          cfg,
		producer:     producer[cfg.App.Name],
		batchTime:    batchTime,
		mapper:       make([]ApiRequestTopic, 0),
		fetchChannel: make(chan ApiRequestTopic),
	}

	httpClient.client = resty.New().
		SetJSONMarshaler(common.JsonHander.Marshal).
		SetJSONUnmarshaler(common.JsonHander.Unmarshal).
		SetBaseURL(cfg.Http.BaseUrl)

	if len(cfg.Producer.URL) > 0 {
		go func() {
			httpClient.loop()
		}()
	}

	return &httpClient
}

func (h *HttpClient) GET(url string, router config.Router) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = getRequest(h.client, router)
		resp, err = req.Get(url)

		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	defer h.handleRequestDefer(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return string(resp.Body()), nil
}

func (h *HttpClient) POST(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = getRequest(h.client, router).SetBody(requestBody)
		resp, err = req.Post(url)

		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	defer h.handleRequestDefer(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return string(resp.Body()), nil
}

func (h *HttpClient) DELETE(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = getRequest(h.client, router).SetBody(requestBody)
		resp, err = req.Delete(url)

		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	defer h.handleRequestDefer(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return string(resp.Body()), nil
}

func (h *HttpClient) PUT(url string, router config.Router, requestBody interface{}) (interface{}, error) {
	var err error
	var req *resty.Request
	var resp *resty.Response

	_, err = common.CB.Execute(func() ([]byte, error) {
		req = getRequest(h.client, router).SetBody(requestBody)
		resp, err = req.Put(url)

		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	defer h.handleRequestDefer(resp, req.Body)

	if err != nil {
		return nil, err
	}

	return string(resp.Body()), nil
}

func getRequest(client *resty.Client, router config.Router) *resty.Request {
	//h.client.R().SetAuthScheme().SetAuthToken().SetHeader()
	req := client.R().EnableTrace()
	if router.Auth != nil {
		if len(router.Auth.Key) != 0 {
			req.SetAuthScheme(router.Auth.Key)
		}
		req.SetAuthToken(router.Auth.Token)
	}
	if router.Header != nil {
		req.SetHeaders(router.Header)
	}
	return req
}
