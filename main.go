package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/profiler"
)

func main() {
	if err := profiler.Start(profiler.Config{}); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type incomingUpdate struct {
	// UpdateID int `json:"update_id"`
	Message struct {
		// MessageId uint64 `json:"message_id"`
		// From      struct {
		// 	Id           uint64
		// 	IsBot        bool   `json:"is_bot"`
		// 	FirstName    string `json:"first_name"`
		// 	Username     string
		// 	LanguageCode string `json:"language_code"`
		// }
		Chat struct {
			ID int `json:"id"`
			// FirstName string `json:"first_name"`
			// Username  string
			// ChatType  string `json:"type"`
		}
		// Date     uint64
		Text string
		// Entities []struct {
		// 	Offset     uint64
		// 	Length     uint64
		// 	EntityType string `json:"type"`
		// }
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", body)
	var data incomingUpdate
	if err := json.Unmarshal(body, &data); err != nil {
		println(err.Error())
		return
	}

	command := strings.Replace(data.Message.Text, "@mktovw_bot", "", -1)
	var message string

	switch command {
	case "/cur":
		message = getCurrencyOverview()
	case "/com":
		message = getCommodityOverview()
	case "/start":
		message = start
	case "/help":
		message = help
	case "/settings":
		message = settings
	case "":
	default:
		message = default_msg
	}

	resp_struct := struct {
		ChatId int    `json:"chat_id"`
		Method string `json:"method"`
		Text   string `json:"text"`
	}{
		ChatId: data.Message.Chat.ID,
		Method: "sendMessage",
		Text:   message,
	}
	resp, err := json.Marshal(resp_struct)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(resp))
}

const start = "I'm written by eepnv. Please support by donating me your code.\n https://github.com/eepnt/tgbot_mktovw"
const help = "Exchange code:\nBFX: Bitfinex\nFFA: FreeForexApi"
const settings = "I have nothing to set yet"
const default_msg = "i can't understand what you mean"
