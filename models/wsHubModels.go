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
    hub.Data.Handler(msg, hub)
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
func (d *DataModel) Handler(data []byte, hub *HubModel) bool {
  var ws WsData
  fmt.Println("msg!!!!!!!!!!!!!!!!!!: ", string(data))
  ok := CwJSON.Unmarshal(data, &ws)
  fmt.Println("wsData!!!!!!!!!!!!!!!!!!!!!:", ws)
  if !ok { return false }

  switch ws.Code {
  case 0:
    return d.Ivtt(&ws.Ivtt, data, hub)
  case 1:
    return d.Rpl(&ws.Rpl, data, hub)
  case 3:
    return d.Msg(&ws.Msg, data, hub)
  }

  return false
}
/******************************************
 * ivtt handler
 *****************************************/
func (d *DataModel) Ivtt(i *JsIvtt, data []byte, hub *HubModel) bool {
  ok := Crud.User.Exist(i.Sender) && Crud.User.Exist(i.Receiver)
  if !ok { return false }

  ok = Crud.FriendList.ExistList(&JsFl{
    Username: i.Sender,
    Friusername: i.Receiver,
  })

  if ok { return false }

  ok = Crud.Invitation.Insert(i)

  if !ok { return false }

  v, ok := hub.Exist(i.Receiver)
  if !ok { return false }
  v.WriteMessage(1, data)

  return true
}

/*********************************************
 * reply handler
 ********************************************/
func (d *DataModel) Rpl(r *JsRpl, data []byte, hub *HubModel) bool {
  ok := App.Rpl.Check(r)
  if !ok { return false }

  if r.Attitude == "agree" {
    ok := App.Rpl.AddFri(r)
    if !ok { return false }

    ok = d.SendFif(r.Me, r.Obj, hub)
    if !ok { return false }

    ok = d.SendFif(r.Obj, r.Me, hub)
    if !ok { return false}
  }

  v, ok := hub.Exist(r.Obj)
  if !ok { return false }

  v.WriteMessage(1, data)
  return true
}

func (d *DataModel) SendFif(src string, obj string, hub *HubModel) bool {
  data, ok := App.Rpl.Marshal(src)
  if !ok { return false }

  v, ok := hub.Exist(obj)
  if !ok { return false } 

  v.WriteMessage(1, data)
  return true
}

/**************************************************
 * msg handler
 *************************************************/
func (d *DataModel) Msg(m *JsMsg, data []byte, hub *HubModel) bool {
  ok := App.Msg.CheckAndInsert(m)
  if !ok { return false }

  v, ok := hub.Exist(m.Receiver)
  v.WriteMessage(1, data)

  v, ok = hub.Exist(m.Sender)
  v.WriteMessage(1, data)

  return true
}
