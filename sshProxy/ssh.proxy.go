package main

import (
	ssh_proxy "ssh_proxy/code/ssh.proxy"
	"time"
)

func main() {
	ssh_proxy.Start()
	for {
		time.Sleep(1 * time.Hour)
	}
}
