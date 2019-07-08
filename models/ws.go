package models

import (
  "github.com/gorilla/websocket"
  "fmt"
)

type ConnInfo struct {
  Username string
  Conn *websocket.Conn
}


/*
 * datatype determine the type of hub
 * 0: invitation data
 */
type Hub struct {
  conns map[string]*websocket.Conn
  register chan *ConnInfo
  unregister chan *ConnInfo
  broadcast chan []byte
  datatype int
}

func NewHub(dt int) *Hub {
  return &Hub{
    conns: make(map[string]*websocket.Conn),
    register: make(chan *ConnInfo),
    unregister: make(chan *ConnInfo),
    broadcast: make(chan []byte),
    datatype: dt,
  }
}

func (hub *Hub) Run() {
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

func (hub *Hub) CloseConn(conninfo *ConnInfo) {
  if _, ok := hub.conns[conninfo.Username]; ok {
    delete(hub.conns, conninfo.Username)
    conninfo.Conn.Close()
  }
}

func (hub *Hub) RegisterConn(conninfo *ConnInfo) {
  hub.register <- conninfo
}

func (hub *Hub) RecMsg(conninfo *ConnInfo) {
  defer func(){
    hub.unregister <- conninfo
  }()

  for {
    _, msg, err := conninfo.Conn.ReadMessage()
    if err != nil {
      return
    }

    switch t := hub.datatype; t {
    case 0: //invitation
      var ivtt InvitationJSON
      if AnalyzeInvitationJson(&ivtt, msg) == true {
        if InsertInvitation(&ivtt) == true {
          hub.broadcast <- msg
        }
      }
    }

    //hub.broadcast <- msg
  }
}

func (hub *Hub) SendMsg(conninfo *ConnInfo) {
  for {
    msg := <-hub.broadcast

    switch t := hub.datatype; t {
    case 0: //invitation
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
}
