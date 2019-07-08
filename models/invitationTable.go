package models

import (
  "github.com/astaxie/beego/orm"
)
//here, sender and receiver are nickname


func InsertInvitation(i *InvitationJSON) bool {
  o := orm.NewOrm()
  o.Using("default")

  invite := Invitation{
    Sender: i.Sender,
    Receiver: i.Receiver,
    Msg: i.Msg,
  }

  _, err := o.Insert(&invite)
  if err == nil {
    return true
  }

  return false
}

func ReadInvitation(result *[]InvitationJSON, name string, key string) bool {
  o := orm.NewOrm()
  o.Using("default")

  var invites []Invitation

  _, err := o.QueryTable("invitation").Filter(key, name).All(&invites)
  if err == nil {
    for _, v := range invites {
      var item InvitationJSON
      item.Sender = v.Sender
      item.Receiver = v.Receiver
      item.Msg = v.Msg
      *result = append(*result, item)
    }
    return true
  }

  return false

}

func ReadInvitationId(name string, key string) (int, bool) {
  o := orm.NewOrm()
  o.Using("default")

  var invite Invitation
  if key == "Sender" {
    invite.Sender = name
  } else if key == "Receiver" {
    invite.Receiver = name
  } else {
    return -1, false
  }

  err := o.Read(&invite, key)

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return -1, false
  } else {
    return invite.Iid, true
  }
}

func DeleteInvitation(name string, key string) bool {
  o := orm.NewOrm()
  o.Using("default")

  id, myerr := ReadInvitationId(name, key)
  if  myerr == false {
    return false
  }

  if _, err := o.Delete(&Invitation{Iid: id}); err == nil {
    return true
  }
  return false
}


