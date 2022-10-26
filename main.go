package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var TGBOTTOKEN string
var DEBUG string
var MASTER_ID int64
var CHANNEL_REPORT int64
var CHANNEL_SLEEP int64
var err error

func main() {

	//读取环境变量到全局变量中
	TGBOTTOKEN = os.Getenv("TGBOTTOKEN")
	DEBUG = os.Getenv("DEBUG")
	MASTER_ID, err = strconv.ParseInt(os.Getenv("MASTER_ID"), 10, 64)
	if err != nil {
		log.Panic(err.Error())
	}
	CHANNEL_REPORT, err = strconv.ParseInt(os.Getenv("CHANNEL_REPORT"), 10, 64)
	if err != nil {
		log.Panic(err.Error())
	}
	CHANNEL_SLEEP, err = strconv.ParseInt(os.Getenv("CHANNEL_SLEEP"), 10, 64)
	if err != nil {
		log.Panic(err.Error())
	}

	//使用TGBOTTOKEN登录机器人
	bot, err := tgbotapi.NewBotAPI(TGBOTTOKEN)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Default().Printf("Authorized on account %s", bot.Self.UserName)

	if DEBUG == "true" {
		bot.Debug = true
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() { // If we got a message
			log.Default().Printf("[COMMAND::%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Command() == "start" || update.Message.Command() == "echo" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
