package client

import (
	"api-gateway-study/common"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func (h *HttpClient) fetchToKafka() {
	h.fetchLock.Lock()
	defer h.fetchLock.Unlock()
	if len(h.mapper) > 0 {
		ent := h.mapper
		h.mapper = make([]ApiRequestTopic, 0)
		v, err := common.JsonHander.Marshal(ent)
		if err == nil {
			h.producer.SendEvent(v)
		} else {
			fmt.Println("Failed to marshal json", err)
		}
	}
}

func (h *HttpClient) loop() {
	ticker := time.NewTicker(time.Duration(h.batchTime) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		h.fetchToKafka()
	}
}

func (h *HttpClient) handleRequestDefer(resp *resty.Response, request interface{}) {
	if len(h.cfg.Producer.URL) > 0 {
		h.mapper = append(h.mapper, NewApiRequestTopic(resp, request))
	}
}
