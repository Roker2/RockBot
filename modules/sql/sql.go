package sql

import (
  "database/sql"
  "github.com/PaulSonOfLars/gotgbot/ext"
  "github.com/Roker2/RockBot/modules/errors"
  "github.com/Roker2/RockBot/modules/texts"
  _ "github.com/lib/pq"
  "os"
  "strings"
)

const chatinfoTable = "CREATE TABLE IF NOT EXISTS chatinfo (id BIGINT PRIMARY KEY, warns_quantity INTEGER, welcome TEXT, rules TEXT, disabled_commands TEXT DEFAULT '');"

const usersTable = "CREATE TABLE IF NOT EXISTS users(id BIGINT PRIMARY KEY, warns INTEGER, ChatId BIGINT, UserId BIGINT);"

const allUsersTable = "CREATE TABLE IF NOT EXISTS AllUsers(id BIGINT PRIMARY KEY, UserName TEXT);"

func openDataBase() (*sql.DB, error) {
  return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}

func GetWarnsQuantityOfChat (ChatId int) (int, error) {
  db, err := openDataBase()
  defer db.Close()
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
  db, err := openDataBase()
  defer db.Close()
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
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return texts.DefaultWelcome, err
  }
  _, err = db.Exec(chatinfoTable)
  if err != nil {
    return texts.DefaultWelcome, err
  }
  var welcome string
  err = db.QueryRow("SELECT welcome FROM chatinfo WHERE id = $1;", ChatId).Scan(&welcome)
  if err != nil {
    return texts.DefaultWelcome, err
  }
  return welcome, err
}

func SetWelcome(ChatId int, welcomeText string) error {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return err
  }
  _, err = db.Exec(chatinfoTable)
  var chatExist bool
  err = db.QueryRow("SELECT count (1) FROM chatinfo WHERE id = $1", ChatId).Scan(&chatExist)
  if err != nil {
    return err
  }
  if !chatExist {
    _, err = db.Exec("INSERT INTO chatinfo (id, welcome) VALUES ($1, $2)", ChatId, welcomeText)
  } else {
    _, err = db.Exec("UPDATE chatinfo SET welcome = $2 WHERE id = $1;", ChatId, welcomeText)
  }
  return err
}

func GetRules (ChatId int) (string, error) {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return texts.DefaultRules, err
  }
  _, err = db.Exec(chatinfoTable)
  if err != nil {
    return texts.DefaultRules, err
  }
  var rules string
  err = db.QueryRow("SELECT rules FROM chatinfo WHERE id = $1;", ChatId).Scan(&rules)
  if err != nil {
    return texts.DefaultRules, err
  }
  return rules, err
}

func SetRules(ChatId int, rulesText string) error {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return err
  }
  _, err = db.Exec(chatinfoTable)
  var chatExist bool
  err = db.QueryRow("SELECT count (1) FROM chatinfo WHERE id = $1", ChatId).Scan(&chatExist)
  if err != nil {
    return err
  }
  if !chatExist {
    _, err = db.Exec("INSERT INTO chatinfo (id, rules) VALUES ($1, $2)", ChatId, rulesText)
  } else {
    _, err = db.Exec("UPDATE chatinfo SET rules = $2 WHERE id = $1;", ChatId, rulesText)
  }
  return err
}

func AddUserWarn (ChatId int, UserId int) (int, error) {
  db, err := openDataBase()
  defer db.Close()
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

func RemoveUserWarn (ChatId int, UserId int) (int, error) {
  db, err := openDataBase()
  defer db.Close()
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
  warns--
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
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return -1, err
  }
  _, err = db.Exec(usersTable)
  if err != nil {
    return -1, err
  }
  var warns int
  err = db.QueryRow("SELECT warns FROM users WHERE id = $1 ;", UserId + ChatId).Scan(&warns)
  //log.Print(strconv.Itoa(UserId + ChatId))
  if err != nil {
    if strings.Contains(err.Error(), "sql: no rows in result set") {
      errors.SendError(err)
      return 0, nil
    } else {
      return -1, err
    }
  }
  return warns, nil
}

func ResetUserWarns (ChatId int, UserId int) error {
  db, err := openDataBase()
  defer db.Close()
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

func SaveUser(user *ext.User) error {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return err
  }
  _, err = db.Exec(allUsersTable)
  if err != nil {
    return err
  }
  var userExist bool
  err = db.QueryRow("SELECT count (1) FROM AllUsers WHERE id = $1", user.Id).Scan(&userExist)
  if !userExist {
    _, err = db.Exec("INSERT INTO AllUsers(id, UserName) VALUES ($1, $2);", user.Id, user.Username)
  } else {
    var needToChange bool
    err = db.QueryRow("SELECT count (1) FROM AllUsers WHERE id = $1 AND UserName = $2", user.Id, user.Username).Scan(&needToChange)
    needToChange = !needToChange
    if needToChange {
      _, err = db.Exec("UPDATE AllUsers SET UserName = $1 WHERE id = $2 ;", user.Username, user.Id)
    }
  }
  return err
}

func GetUserId(userName string) (int, error) {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    //return 0, err
    errors.SendError(err)
  }
  id := 0
  err = db.QueryRow("SELECT id FROM AllUsers WHERE UserName = $1 ;", userName).Scan(&id)
  return id, err
}

func GetDisabledCommands(ChatId int) (string, error) {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return "", err
  }
  var disabledCommands string
  err = db.QueryRow("SELECT disabled_commands FROM chatinfo WHERE id = $1;", ChatId).Scan(&disabledCommands)
  if err != nil {
    if strings.Contains(err.Error(), "sql: no rows in result set") || strings.Contains(err.Error(), "converting NULL to string is unsupported") {
      errors.SendError(err)
      return "", nil
    } else {
      return "", err
    }
  }
  return disabledCommands, err
}

func SetDisabledCommands(ChatId int, commands string) error {
  db, err := openDataBase()
  defer db.Close()
  if err != nil {
    return err
  }
  _, err = db.Exec(chatinfoTable)
  if err != nil {
    return err
  }
  var chatExist bool
  err = db.QueryRow("SELECT count (1) FROM chatinfo WHERE id = $1", ChatId).Scan(&chatExist)
  if !chatExist {
    _, err = db.Exec("INSERT INTO chatinfo (disabled_commands, id) VALUES ($1, $2);", commands, ChatId)
    if err != nil {
      return err
    }
  } else {
    _, err = db.Exec("UPDATE chatinfo SET disabled_commands = $1 WHERE id = $2 ;", commands, ChatId)
    if err != nil {
      return err
    }
  }
  return nil
}