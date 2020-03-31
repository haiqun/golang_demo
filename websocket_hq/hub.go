package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// 注册客户
	clients map[*Client]bool

	// 来自客户端的入站消息
	broadcast chan []byte

	// 注册 来自客户端的请求
	register chan *Client

	// 取消注册 来自客户端的请求
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}