package p2p

import (
	"net/http"

	"github.com/chiwon99881/chyocoin/utils"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

// Upgrade is function of router
func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utils.HandleError(err)
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for _, aConn := range conns {
			if aConn != conn {
				utils.HandleError(aConn.WriteMessage(websocket.TextMessage, p))
			}
		}
	}
}
