package main

import (
	"encoding/json"
	"github-telegram-notify/types"
	"github-telegram-notify/utils"
	"os"
)

func main() {
	tgToken := os.Getenv("INPUT_BOT_TOKEN")
	chatID := os.Getenv("INPUT_CHAT_ID")
	topicID := os.Getenv("INPUT_TOPIC_ID")
	gitEventRaw := os.Getenv("INPUT_GIT_EVENT")
	authorTag := os.Getenv("INPUT_AUTHOR_TAG")
	print(gitEventRaw)
	var gitEvent *types.Metadata
	err := json.Unmarshal([]byte(gitEventRaw), &gitEvent)
	if err != nil {
		panic(err)
	}
	text, markupText, markupUrl, err := utils.CreateContents(gitEvent, authorTag)
	if err != nil {
		panic(err)
	}
	errMessage := utils.SendMessage(tgToken, chatID, text, markupText, markupUrl, topicID)
	if errMessage.Description != "" {
		panic(errMessage.String())
	}

}