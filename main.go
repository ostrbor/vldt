package vldt

import (
	"encoding/json"
	"fmt"
)

var (
	MaxStringLength = 100
)

type Problem struct {
	Detail string `json:"detail"`
}

func JSON[T any](b []byte) (v T, errs []Problem) {
	req := new(T)
	if err := json.Unmarshal(b, req); err != nil {
		detail := ""
		if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			detail = fmt.Sprintf("%s: invalid type '%s'", jsonErr.Field, jsonErr.Value)
		} else {
			detail = err.Error()
		}
		p := Problem{Detail: detail}
		errs = append(errs, p)
		return *req, errs
	}
	return *req, nil
}
