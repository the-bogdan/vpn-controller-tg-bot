package main

import (
	"log"
	"vpn-controller-tg-bot/src/bot"
)

const botToken = "<<you telegram bot token>>"

var allowedUsers = map[int64]bool{
	<<telegram_user_id>>: true,
}

func main() {
	tgBot, err := bot.NewBot(botToken, allowedUsers)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Bot started")

	tgBot.Run()
}
