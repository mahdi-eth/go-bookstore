package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(req *http.Request, x interface{}) {
	if body, err := io.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
		  return
		}
	}
}