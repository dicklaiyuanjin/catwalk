package models

import (
  "github.com/gorilla/websocket"
  "fmt"
)


/******************************************************
* invitation send part
******************************************************/

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

    var ivtt JsIvtt
    if CwJSON.Unmarshal(msg, &ivtt) == true {
      fmt.Println("ivtt!!!!!!!!!!!!!!!!!!!!!!!!: ", ivtt)
      if ExistUsername(ivtt.Sender) && ExistUsername(ivtt.Receiver) {
        if InsertInvitation(&ivtt) == true {
          hub.broadcast <- msg
        }
      }
    }
  }
}

func (hub *IvttHub) SendMsg(ci *ConnInfo) {
  for {
    msg := <-hub.broadcast
    var ivtt JsIvtt
    if CwJSON.Unmarshal(msg, &ivtt) == true {
      if v, ok :=hub.conns[ivtt.Receiver]; ok {
        v.WriteMessage(1, msg)
      }
    }
  }
}

/**************************************************
* invitation reply part
**************************************************/
