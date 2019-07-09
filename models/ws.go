package models

import (
  "github.com/gorilla/websocket"
  "fmt"
)

type IvttHub struct {
  conns map[string]*websocket.Conn
  register chan *ConnInfo
  unregister chan *ConnInfo
  broadcast chan []byte
}

func NewIvttHub() *IvttHub {
  return &IvttHub{
    conns: make(map[string]*websocket.Conn),
    register: make(chan *ConnInfo),
    unregister: make(chan *ConnInfo),
    broadcast: make(chan []byte),
  }
}

func (hub *IvttHub) Run() {
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

func (hub *IvttHub) CloseConn(ci *ConnInfo) {
  if _, ok := hub.conns[ci.Username]; ok {
    delete(hub.conns, ci.Username)
    ci.Conn.Close()
  }
}

func (hub *IvttHub) RegisterConn(ci *ConnInfo) {
  hub.register <- ci
}

func (hub *IvttHub) RecMsg(ci *ConnInfo) {
  defer func(){
    hub.unregister <- ci
  }()

  for {
    _, msg, err := ci.Conn.ReadMessage()
    if err != nil {
      return
    }

    var ivtt InvitationJSON
    if AnalyzeInvitationJson(&ivtt, msg) == true {
      if InsertInvitation(&ivtt) == true {
        hub.broadcast <- msg
      }
    }
  }
}

func (hub *IvttHub) SendMsg(ci *ConnInfo) {
  for {
    msg := <-hub.broadcast
    var ivtt InvitationJSON
    if AnalyzeInvitationJson(&ivtt, msg) == true {
      fmt.Println("ivtt!!!!!!!!!!!!!!!!!!!!!!!!!!: ", ivtt)
      rec := ReadUserInfoUsername(ivtt.Receiver)
      if rec != "" {
        if v, ok :=hub.conns[rec]; ok {
          v.WriteMessage(1, msg)
        }
      }
    }
  }
}
