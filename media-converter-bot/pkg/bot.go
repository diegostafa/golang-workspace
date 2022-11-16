package main

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const inputFile = "tmp/input"
const outputFile = "tmp/output"

const helpMsg = `
Media file converter using ffmpeg
Send a video file to get started

Avaiable commands:
- /send - wait for a file to be sent
- /formats - show all the avaiable formats
- /help - show this message`

func eventLoop(bot *tg.BotAPI) error {
	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	receivedFile := false

	for update := range updates {
		msg := update.Message
		if msg != nil {
			if msg.IsCommand() {
				response := tg.NewMessage(msg.Chat.ID, "")

				switch msg.Command() {
				case "send":
					if receivedFile {
						receivedFile = false
						cleanUpTmp()
					}
					response.Text = "Send me a video file"
				case "formats":
					response.Text = stringifyFormats()
				case "help":
					response.Text = helpMsg
				default:
					response.Text = "Unregognized command, try /help."
				}

				_, err := bot.Send(response)
				checkError(err)

			} else if msg.Document != nil {
				doc := msg.Document
				if isMimeTypeVideo(doc.MimeType) {
					_, err := bot.Send(tg.NewMessage(msg.Chat.ID, "Retrieving the file..."))
					checkError(err)

					fileUrl, err := bot.GetFileDirectURL(doc.FileID)
					checkError(err)

					err = downloadFile(fileUrl, inputFile)
					checkError(err)

					_, err = bot.Send(tg.NewMessage(msg.Chat.ID, "Received: "+doc.FileName+"\nType: "+doc.MimeType))
					checkError(err)

					_, err = bot.Send(tg.NewMessage(msg.Chat.ID, "Send me a new format"))
					checkError(err)

					receivedFile = true
				} else {
					_, err := bot.Send(tg.NewMessage(msg.Chat.ID, "The file you sent me is not a video"))
					checkError(err)
				}
			} else if receivedFile {
				if msg.Text != "" && isFormatAvaiable(msg.Text) {
					_, err := bot.Send(tg.NewMessage(msg.Chat.ID, "Found the format! encoding with "+msg.Text+"..."))
					checkError(err)

					transcodedFile := outputFile + "." + msg.Text
					trans := initTranscoder(inputFile, transcodedFile)

					done := trans.Run(false)
					err = <-done
					if err != nil {
						_, _ = bot.Send(tg.NewMessage(msg.Chat.ID, "Couldn't convert: "+msg.Text))
					} else {
						_, err = bot.Send(tg.NewMessage(msg.Chat.ID, "Done! uploading the file..."))
						checkError(err)

						file := tg.FilePath(transcodedFile)
						_, err = bot.Send(tg.NewDocument(msg.Chat.ID, file))
						checkError(err)

						_, err = bot.Send(tg.NewMessage(msg.Chat.ID, "You can send me another format, or start over by sending a new file"))
						checkError(err)
					}
				} else {
					_, err := bot.Send(tg.NewMessage(msg.Chat.ID, "Unrecognized format, check /formats"))
					checkError(err)
				}
			} else {
				_, err := bot.Send(tg.NewMessage(msg.Chat.ID, "I have no file to work on"))
				checkError(err)
			}
		}
	}
	return nil
}
