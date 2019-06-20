.PHONY: all

all: server client

server: server.go
	go build server.go

client: client.go
	go build client.go

clean:
	rm server client
