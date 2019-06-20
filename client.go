package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"time"
)

func main() {
	hostName := "localhost"
	portNum := "4245"

	service := hostName + ":" + portNum

	remoteAddr, err := net.ResolveUDPAddr("udp", service)

	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Client running at address: %s", conn.LocalAddr().String())

	counter := 0

	for {
		msg := strconv.Itoa(counter)
		counter++

		// send message to server
		now := time.Now()
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Client: Sent '%s'\n", msg)

		// get message from server
		buff := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buff)
		if err != nil {
			log.Fatal(err)
		}
		then := time.Now()
		latency := float64(then.Sub(now).Nanoseconds()) / math.Pow10(6)

		fmt.Printf("Client: Got '%s'\n", string(buff[:n]))
		fmt.Printf("Latency: %f ms\n", latency)
	}
}
