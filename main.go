package main

import (
	model "github.com/yydspg/model"
	"os"
	"encoding/json"
	"fmt"
)

func main(){
	file, err := os.ReadFile("./model/model.json")
	if err != nil {
		fmt.Print(err)
	}
	var d map[string]interface{}
	err = json.Unmarshal(file, &d)
	if err != nil {
		return
	}
	data, err := model.NewVrit(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(data)
}