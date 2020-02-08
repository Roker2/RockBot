package sqlite

import (
	"../errors"
  "database/sql"
  "os"
  "log"
	"strconv"
  "strings"
  _ "github.com/mattn/go-sqlite3"
)

func GetWarnsQuantityOfChat (ChatId int) int {
  os.Mkdir("databases",0770)
  db, err := sql.Open("sqlite3", "./databases/chatinfo.db")
  if err != nil {
    errors.SendError(err)
    return -1
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS chatinfo (id INTEGER PRIMARY KEY, warns_quantity INTEGER)")
  if err != nil {
    errors.SendError(err)
    return -2 //error: can not to create table
  }
  statement.Exec()
  var warns int
  statement, err = db.Prepare("SELECT warns_quantity FROM chatinfo WHERE id = ?")
  if err != nil {
    errors.SendError(err)
    return -3//error: can not select
  }
  err = statement.QueryRow(ChatId).Scan(&warns)
  if warns == 0 {
    return 5
  }
  return warns
}

func SetWarnsQuantityOfChat (ChatId int, warns int) {
  log.Print("Chat: " + strconv.Itoa(ChatId) + " Warns: " + strconv.Itoa(warns))
  os.Mkdir("databases",0770)
  db, err := sql.Open("sqlite3", "./databases/chatinfo.db")
  if err != nil {
    errors.SendError(err)
    return// -1
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS chatinfo (id INTEGER PRIMARY KEY, warns_quantity INTEGER)")
  if err != nil {
    errors.SendError(err)
    return// -2
  }
  statement.Exec()
  statement, err = db.Prepare("SELECT warns_quantity FROM chatinfo WHERE id = ?")
  if err != nil {
    errors.SendError(err)
    return
  }
  var warns_quantity int
  err = statement.QueryRow(ChatId).Scan(&warns_quantity)
  if warns_quantity == 0 {
    statement, err = db.Prepare("INSERT INTO chatinfo (warns_quantity, id) VALUES (?, ?)")
  } else {
    statement, err = db.Prepare("UPDATE chatinfo SET warns_quantity = ? WHERE id = ?")
  }
  if err != nil {
    errors.SendError(err)
    return// err
  }
  statement.Exec(warns, ChatId)
}

func AddUserWarn (ChatId int, UserId int) int {
  os.Mkdir("databases",0770)
  db, err := sql.Open("sqlite3", "./databases/warns.db")
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, warns INTEGER)")
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  statement.Exec()
  var warns int
  statement, err = db.Prepare("SELECT warns FROM user WHERE id = ?")
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  err = statement.QueryRow(UserId + ChatId).Scan(&warns)
  if warns == 0 {
    statement, err = db.Prepare("INSERT INTO user (warns, id) VALUES (?, ?)")
  } else {
    statement, err = db.Prepare("UPDATE user SET warns = ? WHERE id = ?")
  }
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  warns++
  log.Print("UserId + ChatId: " + strconv.Itoa(UserId + ChatId) + " warns: " + strconv.Itoa(warns))
  statement.Exec(warns, UserId + ChatId)
  return warns
}

func GetUserWarns (ChatId int, UserId int) int {
  os.Mkdir("databases",0770)
  db, err := sql.Open("sqlite3", "./databases/warns.db")
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, warns INTEGER)")
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  statement.Exec()
  var warns int
  statement, err = db.Prepare("SELECT warns FROM user WHERE id = ?")
  if err != nil {
    errors.SendError(err)
    return -1// err
  }
  err = statement.QueryRow(UserId + ChatId).Scan(&warns)
  if err != nil {
    errors.SendError(err)
    if strings.Contains(err.Error(), "sql: no rows in result set") {
      return 0
    } else {
      return -1// err
    }
  }
  return warns
}
