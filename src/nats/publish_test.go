package nats

import (
	"encoding/json"
	"testing"
)

func TestPublish(t *testing.T) {

	data := RequestData{
		Method:  "testing",
		URL:     "www.testing.com",
		UUID:    "test1543215671",
		Status:  "",
		Subject: "",
	}
	body, _ := json.Marshal(data)
	Publish(body, "nodeScan")
}
