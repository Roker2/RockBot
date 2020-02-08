package sqlite

import (
	"../errors"
  "database/sql"
  "os"
  "log"
	"strconv"
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
