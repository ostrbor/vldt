package vldt

import (
	"encoding/json"
	"fmt"
	"testing"
)

type resp struct {
	Name string `json:"name"`
	Problem
}

func TestName(t *testing.T) {
	p := Str("name", "AAAAAAAAAAA").Max(2).Errs()

	r := resp{
		Name:    "myname",
		Problem: p[0],
	}

	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
