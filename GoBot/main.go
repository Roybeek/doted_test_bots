package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	botToken := "1958943268:AAFJkPyMNRN4cUVJhcVgwHIL3SnMMsH7Gn8"
	// https://api.telegram.org/bot<token>/METHOD_NAME
	botAPI := "https://api.telegram.org/bot"
	botURL := botAPI + botToken
	offset := 0
	poll_flag := 0
	name_surname := ""
	for {
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			log.Println("что-то пошло не так", err.Error())
		}
		for _, update := range updates {
			msg_text := update.Message.Text
			respond_text := ""
			switch poll_flag {
			case 0:
				if msg_text == "/start" {
					poll_flag = 1 // получили компанду старт
					respond_text = "Привет, введи свое имя"
				}
			case 1:
				if msg_text != "" {
					name_surname = msg_text
					poll_flag = 2
					respond_text = "Введи свою фамилию"
				}
			case 2:
				if msg_text != "" {
					poll_flag = 0
					respond_text = "Приятно познакомиться, " + msg_text + " " + name_surname
				}
			}

			err := respond(botURL, update, respond_text)
			if err != nil {
				log.Println("getting error when try to respond: ", err.Error())
			}
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
}

func getUpdates(botURL string, offset int) ([]Update, error) {
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil

}

func respond(botURL string, update Update, respond_text string) error {
	log.Println("try resp")
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = respond_text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))

	if err != nil {
		return err
	}

	return nil
}
