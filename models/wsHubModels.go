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

func (d *DataModel) Rpl(r *JsRpl, data []byte, hub *HubModel) bool {
  //1.查看是否有这么一条好友邀请
  //1.1若有，则删除，继续下一步（2）
  //1.2否则，返回错误信息
  ok := Crud.Invitation.Exist(r.Obj, r.Me)
  if !ok { return false }

  ok = Crud.Invitation.Delete(r.Obj, r.Me)
  if !ok { return false }

  //另一侧的删除
  Crud.Invitation.Delete(r.Me, r.Obj)

  //2.查看回复是否同意邀请
  //2.1同意邀请，则为双方添加好友列表，并且生成双方用户信息作为好友信息互发，继续下一步（3）
  //2.2否则，跳转到（3）
  if r.Attitude == "agree" {
    ok = Crud.FriendList.Insert(&JsFl{
      Username: r.Me,
      Friusername: r.Obj,
    })
    if !ok { return false }

    //Me to Obj
    ok = d.SendFif(r.Me, r.Obj, hub)
    if !ok { return false }

    //Obj to Me
    ok = d.SendFif(r.Obj, r.Me, hub)
    if !ok { return false}
  }

  //发送reply给对方，让对方获悉自己的邀请是否成功
  v, ok := hub.Exist(r.Obj)
  if !ok { return false }

  v.WriteMessage(1, data)
  return true
}

//将src的info发送给obj在线用户
func (d *DataModel) SendFif(src string, obj string, hub *HubModel) bool {
  Src := JsUif{Username: src}
  ok := Crud.Uif.Read(&Src, "Username")
  if !ok { return false }

  ws := WsData{
    Code: 2,
    Fif: Src,
  }

  data, ok := CwJSON.Marshal(ws)
  if !ok { return false }

  v, ok := hub.Exist(obj)
  if !ok { return false }

  fmt.Println("fif!!!!!!!!!!!!!!!!: ", ws.Fif.Username);

  v.WriteMessage(1, data)
  return true
}

func (d *DataModel) Msg(m *JsMsg, data []byte, hub *HubModel) bool {
  ok := App.Msg.CheckAndInsert(m)
  if !ok { return false }

  v, ok := hub.Exist(m.Receiver)
  v.WriteMessage(1, data)

  v, ok = hub.Exist(m.Sender)
  v.WriteMessage(1, data)

  return true
}
