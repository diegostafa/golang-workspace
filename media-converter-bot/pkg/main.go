package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cleanUpTmp()

	token, err := readApiToken()
	checkError(err)

	bot, err := tg.NewBotAPI(token)
	checkError(err)

	err = eventLoop(bot)
	checkError(err)
}
