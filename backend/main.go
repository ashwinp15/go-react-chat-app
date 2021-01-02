package main

import (
	 "fmt"
	 //"log"
	 "net/http"
	 "github.com/ashwinp15/chat-app/pkg/websocket"
)

// Websocket endpoint
func serveWs (pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	 fmt.Println(r.Host)

	 // upgrading the connection to websocket
	 conn, err := websocket.Upgrade(w, r)
	 if err != nil {
			fmt.Fprintf (w, "%+V\n", err)
	 }

	 client := &websocket.Client {
			Conn: conn,
			Pool: pool,
	 }

	 pool.Register <- client
	 client.Read()
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf (w, "Simple web server")
	 }

func setupRoutes () {
	 http.HandleFunc("/", indexHandler)

	 // Now, serve /ws endpoint using the serveWs function
	 pool := websocket.NewPool()
	 go pool.Start()
	 http.HandleFunc ("/ws", func (w http.ResponseWriter, r *http.Request) {serveWs(pool, w, r)} )
}

func main () {
	 setupRoutes()
	 fmt.Println("Backend running")
	 http.ListenAndServe(":8080", nil)
}

