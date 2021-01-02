package websocket

import (
	 //"fmt"
	 //"io"
	 "log"
	 "net/http"
	 "github.com/gorilla/websocket"
)

// Setting Upgrader 
var upgrader = websocket.Upgrader {
	 ReadBufferSize: 1024,
	 WriteBufferSize: 1024,

	 // Check origin function. Might edit later
	 CheckOrigin: func (r *http.Request) bool { return true },
}

func Upgrade (w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	 ws, err := upgrader.Upgrade(w, r, nil)
	 if err != nil {
			log.Println (err)
			return ws, err
	 }
	 return ws, nil
}

