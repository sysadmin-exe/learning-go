package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	var body []byte
	if body, err := io.ReadAll(r.Body); err == nil {
		_ = r.Body.Close()
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
	if err := json.Unmarshal([]byte(body), x); err != nil {
		return
	}

}
