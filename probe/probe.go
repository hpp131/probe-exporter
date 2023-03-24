package probe

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func TcpProbe(target string, c chan<- map[string]int) {
	mapping := make(map[string]int)
	conn, err := net.DialTimeout("tcp", target, 3*time.Second)
	fmt.Println("after DialTimeout..")
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("recover the panic:", rec)
		}
	}()
	if err != nil {
		mapping[target] = 0
	} else {
		mapping[target] = 1
	}
	c <- mapping
	conn.Close()
	fmt.Println("already send mapping to c", target)

}

func HttpProbe(url string, c chan<- map[string]int) {
	mapping := make(map[string]int)
	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("recover the panic:", rec)
		}
	}()
	if err == nil {
		if resp.StatusCode == 200 {
			mapping[url] = 1
		} else {
			mapping[url] = 0
		}
	} else {
		mapping[url] = 0
	}
	c <- mapping
	defer resp.Body.Close()

}
