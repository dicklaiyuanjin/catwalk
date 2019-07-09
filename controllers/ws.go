package controllers

import (
	"github.com/astaxie/beego"
  "github.com/gorilla/websocket"
  "net/http"
  "catwalk/models"
)

type WsController struct {
	beego.Controller
}

var ivttHub *models.IvttHub

func init() {
  ivttHub = models.NewIvttHub()
  go ivttHub.Run()
}

func (this *WsController) JoinIvttUser() {
  username := this.GetString("username")
  if len(username) == 0 {
    this.Redirect("/", 302)
    return
  }

  conn, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
  if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

  conninfo := models.ConnInfo{
    Username: username,
    Conn: conn,
  }

  ivttHub.RegisterConn(&conninfo)
  defer func(){
    ivttHub.CloseConn(&conninfo)
  }()

  go ivttHub.RecMsg(&conninfo)
  ivttHub.SendMsg(&conninfo)
}

func (this *WsController) ReplyIvttUser() {
  for{}
}
