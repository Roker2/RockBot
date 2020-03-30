package sql

import (
  "database/sql"
  "github.com/Roker2/RockBot/modules/errors"
  _ "github.com/lib/pq"
  "log"
  "os"
  "strings"
)

const chatinfoTable = "CREATE TABLE IF NOT EXISTS chatinfo (id BIGINT PRIMARY KEY, warns_quantity INTEGER, welcome TEXT, rules TEXT);"

const usersTable = "CREATE TABLE IF NOT EXISTS users(id BIGINT PRIMARY KEY, warns INTEGER, ChatId BIGINT, UserId BIGINT);"

func GetWarnsQuantityOfChat (ChatId int) (int, error) {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  //log.Print(os.Getenv("DATABASE_URL"))
  if err != nil {
    return 0, err
  }
  _, err = db.Exec(chatinfoTable)
  if err != nil {
    return -1, err
  }
  var warns int
  err = db.QueryRow("SELECT warns_quantity FROM chatinfo WHERE id = $1 ;", ChatId).Scan(&warns)
  if warns == 0 {
    return 5, nil
  }
  return warns, nil
}

func SetWarnsQuantityOfChat (ChatId int, warns int) error {
  //log.Print("Chat: " + strconv.Itoa(ChatId) + " Warns: " + strconv.Itoa(warns))
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  //log.Print(os.Getenv("DATABASE_URL"))
  if err != nil {
    return err
  }
  _, err = db.Exec(chatinfoTable)
  if err != nil {
    return err
  }
  var warns_quantity int
  err = db.QueryRow("SELECT warns_quantity FROM chatinfo WHERE id = $1 ;", ChatId).Scan(&warns_quantity)
  if warns_quantity == 0 {
    _, err = db.Exec("INSERT INTO chatinfo (warns_quantity, id) VALUES ($1, $2);", warns, ChatId)
    if err != nil {
      return err
    }
  } else {
    _, err = db.Exec("UPDATE chatinfo SET warns_quantity = $1 WHERE id = $2 ;", warns, ChatId)
    if err != nil {
      return err
    }
  }
  return nil
}

func GetWelcome (ChatId int) (string, error) {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    return "Добро пожаловать, {firstName}!", err
  }
  _, err = db.Exec(chatinfoTable)
  if err != nil {
    return "Добро пожаловать, {firstName}!", err
  }
  var welcome string
  err = db.QueryRow("SELECT welcome FROM chatinfo WHERE id = $1;", ChatId).Scan(&welcome)
  if err != nil {
    return "Добро пожаловать, {firstName}!", err
  }
  return welcome, err
}

func AddUserWarn (ChatId int, UserId int) (int, error) {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  //log.Print(os.Getenv("DATABASE_URL"))
  if err != nil {
    return -1, err
  }
  _, err = db.Exec(usersTable)
  if err != nil {
    return -1, err
  }
  warns, err := GetUserWarns(ChatId, UserId)
  if err != nil {
    return -1, err
  }
  //log.Print("UserId + ChatId: " + strconv.Itoa(UserId + ChatId) + " warns: " + strconv.Itoa(warns))
  warns++
  if warns == 1 {
    _, err = db.Exec("INSERT INTO users(id, warns, ChatId, UserId) VALUES ($1, $2, $3, $4);", UserId + ChatId, warns, ChatId, UserId)
    if err != nil {
      return -1, err
    }
  } else {
    _, err = db.Exec("UPDATE users SET warns = $1 WHERE id = $2 ;", warns, UserId + ChatId)
    if err != nil {
      return -1, err
    }
  }
  return warns, nil
}

func GetUserWarns (ChatId int, UserId int) (int, error) {
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  //log.Print(os.Getenv("DATABASE_URL"))
  if err != nil {
    return -1, err
  }
  _, err = db.Exec(usersTable)
  if err != nil {
    return -1, err
  }
  log.Print("HELLO")
  var warns int
  err = db.QueryRow("SELECT warns FROM users WHERE id = $1 ;", UserId + ChatId).Scan(&warns)
  //log.Print(strconv.Itoa(UserId + ChatId))
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
  db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  //log.Print(os.Getenv("DATABASE_URL"))
  if err != nil {
    return err
  }
  _, err = db.Exec(usersTable)
  if err != nil {
    return err
  }
  _, err = db.Exec("DELETE from users where id = $1", UserId + ChatId)
  if err != nil {
    return err
  }
  return nil
}
