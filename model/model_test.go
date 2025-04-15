package model

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {

	file, err := os.ReadFile("./model.json")
	if err != nil {
		return
	}
	var d map[string]interface{}
	err = json.Unmarshal(file, &d)
	if err != nil {
		return
	}
	data, err := NewVrit(d)
	if err != nil {
		return
	}
	fmt.Print(data.m)
}
