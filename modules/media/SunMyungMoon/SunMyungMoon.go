package SunMyungMoon

import (
    "github.com/PaulSonOfLars/gotgbot"
    "github.com/PaulSonOfLars/gotgbot/ext"
    "io/ioutil"
    "math/rand"
    "strings"
    "time"
)

func RandomSMMQ(b ext.Bot, u *gotgbot.Update) error {
    quotesFile, err := ioutil.ReadFile("Quotes.txt")
    if err != nil {
        return err
    }
    quotesStrings := strings.Split(string(quotesFile), "\n")
    rand.Seed(time.Now().UnixNano())
    _, err = b.SendMessage(u.Message.Chat.Id, quotesStrings[rand.Intn(len(quotesStrings))] + "\n© Преподобный Мун Сон Мён")
    return err
}
