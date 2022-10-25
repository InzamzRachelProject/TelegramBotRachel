package main

import (
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {

	var TGBotToken string

	//命令行参数读取
	flag.StringVar(&TGBotToken, "tgbottoken", "", "机器人Token,必填")
	flag.Parse()

	//环境变量读取
	if TGBotToken == "" {
		TGBotToken = os.Getenv("TGBOTTOKEN")
	}

	//官方示例
	bot, err := tgbotapi.NewBotAPI(TGBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() { // If we got a message
			log.Printf("[COMMAND::%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Command() == "start" || update.Message.Command() == "echo" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
