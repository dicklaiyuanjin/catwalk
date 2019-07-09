package models

import (
  "github.com/gorilla/websocket"
)

type ConnInfo struct {
  Username string
  Conn *websocket.Conn
}

//websocket hub
type WsHub interface {
  Run()
  CloseConn(ci *ConnInfo)
  RegisterConn(ci *ConnInfo)
  RecMsg(ci *ConnInfo)
  SendMsg(ci *ConnInfo)
}
