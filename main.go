package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp4", ":8080")
	fmt.Println("listening in port 8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer c.Close()

		fmt.Printf("Serving %s\n", c.RemoteAddr().String())
		for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			command := strings.TrimSpace(string(netData))
			if command == "ping" {
				fmt.Println("This is pong")
				c.Write([]byte(string("pong\n")))
			} else {
				c.Write([]byte(string("unknown command\n")))
			}
		}
	}
}
