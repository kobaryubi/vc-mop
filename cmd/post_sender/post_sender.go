package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kobaryubi/vc-mop/internal/sender"
)

type Params struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

const URL = "http://localhost:8080"

func main() {
	params := new(Params)
	params.Id = 0
	params.Name = "hoge"

	paramsJSON, _ := json.Marshal(params)

	httpSender := sender.NewHTTPSender(URL, "POST")
	data, err := httpSender.Send(paramsJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
