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

/******************************************************
* invitation send part
******************************************************/
type HubModel struct {
  conns map[string]*websocket.Conn
  register chan *ConnInfo
  unregister chan *ConnInfo
  broadcast chan []byte
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

    var ivtt JsIvtt
    if CwJSON.Unmarshal(msg, &ivtt) == true {
      fmt.Println("rec ivtt!!!!!!!!!!!!!!!!!!!!!!!!: ", ivtt)
      if Crud.User.Exist(ivtt.Sender) && Crud.User.Exist(ivtt.Receiver) {
        if Crud.Invitation.Insert(&ivtt) == true {
          hub.broadcast <- msg
        }
      }
    }
  }
}

func (hub *HubModel) SendMsg(ci *ConnInfo) {
  for {
    msg := <-hub.broadcast
    var ivtt JsIvtt
    if CwJSON.Unmarshal(msg, &ivtt) == true {
      fmt.Println("send msg!!!!!!!!!!!!!!!!!!!!!!!!: ", ivtt)
      if v, ok := hub.conns[ivtt.Receiver]; ok {
        v.WriteMessage(1, msg)
      }
    }
  }
}

/**************************************************
* invitation reply part
**************************************************/
