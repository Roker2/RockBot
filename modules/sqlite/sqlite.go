package sqlite

import (
  "github.com/Roker2/RockBot/modules/errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "os"
  "strconv"
  "strings"
)

func GetWarnsQuantityOfChat (ChatId int) (int, error) {
  db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
  if err != nil {
    return 0, err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS chatinfo (id INTEGER PRIMARY KEY, warns_quantity INTEGER)")
  if err != nil {
    errors.SendError(err)
    return 0, err //error: can not to create table
  }
  statement.Exec()
  var warns int
  statement, err = db.Prepare("SELECT warns_quantity FROM chatinfo WHERE id = ?")
  if err != nil {
    errors.SendError(err)
    return 0, err//error: can not select
  }
  err = statement.QueryRow(ChatId).Scan(&warns)
  if warns == 0 {
    return 5, nil
  }
  return warns, nil
}

func SetWarnsQuantityOfChat (ChatId int, warns int) error {
  log.Print("Chat: " + strconv.Itoa(ChatId) + " Warns: " + strconv.Itoa(warns))
  db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
  if err != nil {
    return err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS chatinfo (id INTEGER PRIMARY KEY, warns_quantity INTEGER)")
  if err != nil {
    return err
  }
  statement.Exec()
  statement, err = db.Prepare("SELECT warns_quantity FROM chatinfo WHERE id = ?")
  if err != nil {
    return err
  }
  var warns_quantity int
  err = statement.QueryRow(ChatId).Scan(&warns_quantity)
  if warns_quantity == 0 {
    statement, err = db.Prepare("INSERT INTO chatinfo (warns_quantity, id) VALUES (?, ?)")
  } else {
    statement, err = db.Prepare("UPDATE chatinfo SET warns_quantity = ? WHERE id = ?")
  }
  if err != nil {
    return err
  }
  statement.Exec(warns, ChatId)
  return nil
}

func AddUserWarn (ChatId int, UserId int) (int, error) {
  db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
  if err != nil {
    return -1, err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, warns INTEGER, ChatId INTEGER, UserId INTEGER)")
  if err != nil {
    return -1, err
  }
  statement.Exec()
  warns, err := GetUserWarns(ChatId, UserId)
  if err != nil {
    return -1, err
  }
  log.Print("UserId + ChatId: " + strconv.Itoa(UserId + ChatId) + " warns: " + strconv.Itoa(warns))
  warns++
  if warns == 1 {
    statement, err = db.Prepare("INSERT INTO user (id, warns, ChatId, UserId) VALUES (?, ?, ?, ?)")
    if err != nil {
      return -1, err
    }
    statement.Exec(UserId + ChatId, warns, ChatId, UserId)
  } else {
    statement, err = db.Prepare("UPDATE user SET warns = ? WHERE id = ?")
    if err != nil {
      return -1, err
    }
    statement.Exec(warns, UserId + ChatId)
  }
  return warns, nil
}

func GetUserWarns (ChatId int, UserId int) (int, error) {
  db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
  if err != nil {
    return -1, err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, warns INTEGER, ChatId INTEGER, UserId INTEGER)")
  if err != nil {
    return -1, err
  }
  statement.Exec()
  var warns int
  statement, err = db.Prepare("SELECT warns FROM user WHERE id = ?")
  if err != nil {
    return -1, err
  }
  err = statement.QueryRow(UserId + ChatId).Scan(&warns)
  log.Print(strconv.Itoa(UserId + ChatId))
  if err != nil {
    errors.SendError(err)
    if strings.Contains(err.Error(), "sql: no rows in result set") {
      return 0, nil
    } else {
      return -1, err
    }
  }
  return warns, nil
}

func ResetUserWarns (ChatId int, UserId int) error {
  db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
  if err != nil {
    return err
  }
  statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, warns INTEGER, ChatId INTEGER, UserId INTEGER)")
  if err != nil {
    return err
  }
  statement.Exec()
  statement, err = db.Prepare("DELETE from user where id = ?")
  if err != nil {
    return err
  }
  statement.Exec(UserId + ChatId)
  return nil
}
