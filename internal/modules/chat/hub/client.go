package hub

import "github.com/gorilla/websocket"

type Client struct {
	username string
	outbound chan []byte
	socket   *websocket.Conn
}

func NewClient(socket *websocket.Conn) *Client {
	return &Client{
		outbound: make(chan []byte),
		socket:   socket,
	}
}

func (c *Client) Write() {
	for {
		select {
		case data, ok := <-c.outbound:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (c *Client) Close() {
	c.socket.Close()
	close(c.outbound)
}
