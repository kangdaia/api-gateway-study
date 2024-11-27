package client

import (
	"github.com/go-resty/resty/v2"
)

type ApiRequestTopic struct {
	Trace   trace               `json:"trace"`
	Header  map[string][]string `json:"header"`
	Request interface{}         `json:"request"`
	Code    int                 `json:"code"`
}

type trace struct {
	ConnectionTime float64 `json:"connectionTime"`
}

func NewApiRequestTopic(resp *resty.Response, request interface{}) ApiRequestTopic {
	t := resp.Request.TraceInfo()
	connectionTime := t.ConnTime.Seconds()
	header := resp.Header()
	return ApiRequestTopic{
		Trace:   trace{ConnectionTime: connectionTime},
		Header:  header,
		Request: request,
		Code:    resp.StatusCode(),
	}
}
