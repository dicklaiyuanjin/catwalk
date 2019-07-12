package models

import (
  "github.com/gorilla/websocket"
  "fmt"
)

var Hub *HubModel

func init() {
  Hub = NewHub()
  go Hub.Run()
}


type ConnInfo struct {
  Username string
  Conn *websocket.Conn
}

type HubModel struct {
  conns map[string]*websocket.Conn
  register chan *ConnInfo
  unregister chan *ConnInfo
  broadcast chan []byte

  Data *DataModel
}

func NewHub() *HubModel {
  return &HubModel{
    conns: make(map[string]*websocket.Conn),
    register: make(chan *ConnInfo),
    unregister: make(chan *ConnInfo),
    broadcast: make(chan []byte),
  }
}

func (hub *HubModel) Run() {
  for {
    select {
      case conninfo := <-hub.register:
        hub.conns[conninfo.Username] = conninfo.Conn
      case conninfo := <-hub.unregister:
        if _, ok := hub.conns[conninfo.Username]; ok == true {
          delete(hub.conns, conninfo.Username)
          conninfo.Conn.Close()
        }
    }
  }
}

func (hub *HubModel) CloseConn(ci *ConnInfo) {
  if _, ok := hub.conns[ci.Username]; ok {
    delete(hub.conns, ci.Username)
    ci.Conn.Close()
  }
}

func (hub *HubModel) RegisterConn(ci *ConnInfo) {
  hub.register <- ci
}

func (hub *HubModel) RecMsg(ci *ConnInfo) {
  defer func(){
    hub.unregister <- ci
  }()

  for {
    _, msg, err := ci.Conn.ReadMessage()
    if err != nil {
      return
    }
    hub.broadcast <- msg
  }
}

func (hub *HubModel) SendMsg(ci *ConnInfo) {
  for {
    msg := <-hub.broadcast
    username, ok := hub.Data.Handler(msg)
    if !ok { return }
    v, ok := hub.Exist(username)
    if !ok { return }
    v.WriteMessage(1, msg)
  }
}

func (hub *HubModel) Exist(username string) (*websocket.Conn, bool) {
  v, ok := hub.conns[username]
  return v, ok
}


/*********************************************************
 * data handler
 ********************************************************/

type DataModel struct {
  name string
}

/* 
 * data handler
 * get the selected data
 * return target, ok
 */
func (d *DataModel) Handler(msg []byte) (string, bool) {
  var ws WsData
  ok := CwJSON.Unmarshal(msg, &ws)
  fmt.Println("wsData!!!!!!!!!!!!!!!!!!!!!:", ws)
  if !ok { return "", false }

  switch ws.Code {
  case 0:
    return d.Ivtt(&ws.Ivtt)
  case 1:
  }

  return "", false
}

func (d *DataModel) Ivtt(i *JsIvtt) (string, bool) {
  ok := Crud.User.Exist(i.Sender) && Crud.User.Exist(i.Receiver)
  if !ok { return "", false }

  ok = Crud.Invitation.Insert(i)
  if !ok { return "", false }

  return i.Receiver, true
}


