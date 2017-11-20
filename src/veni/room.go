package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type room struct {
	// forward is a channel that holds incoming messages
	// that shoud be forwarded to other clients.
	forward chan []byte

	join chan *client

	leave chan *client

	clients map[*client]bool
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// joining
			fmt.Println("someone join the room", client)
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			fmt.Println("someone leave the room", client)
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients

			for client := range r.clients {
				select {
				case client.send <- msg:
					//send the message
				default:
					// failed to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/// we get the socket by calling
	socket, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
