package push

import (
	"github.com/astaxie/beego/context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var Clients = make(map[*websocket.Conn]bool, 1024)
var UserID2Client = make(map[int64]*websocket.Conn, 1024)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
	HandshakeTimeout:5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SetUpWebsocket(ctx context.Context) {
	log := logger(ctx)
	log.Info("SetUpWebsocket start")
	ws, err := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	Clients[ws] = true
	log.Info("SetUpWebsocket finish")
}

func PushMessage(ctx context.Context)  {
	log := logger(ctx)
	log.Info("PushMessage start")
	for client := range Clients {
		err := client.WriteMessage(websocket.BinaryMessage, []byte("push from server"))
		log.Info("PushMessage Success")
		go handleMessage(ctx, client)
		if err != nil {
			client.Close()
			delete(Clients, client)
		}
	}
}

func handleMessage(ctx context.Context, ws *websocket.Conn) {
	log := logger(ctx)
	log.Info("handleMessage start")
	for {
		msgType, bytes, err := ws.ReadMessage()
		if err != nil {
			log.WithError(err).Error("ws.ReadMessage Fail")
			delete(Clients, ws)
			break
		}
		log.Infof("msgType: %v, msg: %v", msgType, string(bytes))
	}
}

func logger(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	return logrus.WithFields(fields)
}