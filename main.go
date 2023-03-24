package main

import (
	"blackbox-exporter/probe"
	"fmt"
)

const fileName = "./target.json"

func main() {
	target := new(Target)
	target.ConvertFunc(fileName)
	chNum := len(target.Tcp) + len(target.Http)
	chQueue := make(chan map[string]int)
	defer close(chQueue)

	for _, addr := range target.Tcp {
		addr := addr
		go probe.TcpProbe(addr, chQueue)
	}

	for _, url := range target.Http {
		url := url
		go probe.HttpProbe(url, chQueue)
	}
	for i := 0; i < chNum; i++ {
		fmt.Println("iteration chQueue", <-chQueue)
	}
	fmt.Printf("after close chan ")

}
