# udp-echo
UDP echo server, client. Written as a means to measure latency of the vanilla UDP.

## server.go
- listens for UDP connections to address localhost:4245
- echoes back any data it recieves

## client.go
- sends data to UDP server running at localhost:4245
- reads data it recieves as response from UDP server

## author
Sayan Goswami
