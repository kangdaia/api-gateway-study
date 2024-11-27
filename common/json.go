package common

import (
	"log"

	"github.com/bytedance/sonic"
)

type jsonHandler struct {
	marshal   func(v interface{}) ([]byte, error)
	unmarshal func(buffer []byte, v interface{}) error
}

var JsonHander jsonHandler

func init() {
	JsonHander = jsonHandler{
		marshal:   sonic.Marshal,
		unmarshal: sonic.Unmarshal,
	}
}

func (j jsonHandler) Marshal(v interface{}) ([]byte, error) {
	bytes, err := j.marshal(v)

	if err != nil {
		log.Println("Failed to marshal json", err.Error())
		return nil, err
	}
	return bytes, nil
}

func (j jsonHandler) Unmarshal(buffer []byte, v interface{}) error {
	err := j.unmarshal(buffer, v)
	if err != nil {
		log.Println("Failed to unmarshal json", err.Error())
		return err
	}
	return nil
}
