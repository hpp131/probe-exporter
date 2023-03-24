package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Target struct {
	Tcp  []string `json:"tcp"`
	Http []string `json:"http"`
}

func (t *Target) ConvertFunc(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(content, t); err != nil {
		fmt.Println(err)
	}
}
