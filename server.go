package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	hostName := "localhost"
	portNum := "4245"

	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	if err != nil {
		log.Fatal(err)
	}

	udpListener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer udpListener.Close()

	fmt.Println("UDP server started at: ", service)

	for {
		buff := make([]byte, 1024)

		// get data from client
		n, addr, err := udpListener.ReadFromUDP(buff)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Server: Got '%s'\n", string(buff[:n]))

		// echo data back to client
		_, err = udpListener.WriteToUDP(buff[:n], addr)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Server: Sent '%s'\n", string(buff[:n]))
	}
}
