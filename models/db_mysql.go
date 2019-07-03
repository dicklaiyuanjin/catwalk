package models

import (
  "github.com/astaxie/beego/orm"
   _ "github.com/go-sql-driver/mysql"
)

func init() {
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default", "mysql", "dick:12345678@/catwalk?charset=utf8")
  orm.Debug = true
}


