package main

// import (
// 	"cloud.google.com/go/firestore"
// 	"context"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// )

// var (
// 	projectID        string = "xxxxxx"
// 	FirestoreClient  *firestore.Client
// 	FirestoreContext context.Context
// 	botToken         *string
// )

// func init() {
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, projectID)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}
// 	FirestoreClient = client
// 	FirestoreContext = ctx
// }

// func botURL() string {
// 	if botToken == nil {
// 		if data, err := FirestoreClient.Collection("config").Doc("production").Get(FirestoreContext); err != nil {
// 			panic(err)
// 		} else {
// 			result := data.Data()
// 			s := result["market_overviewer"].(string)
// 			botToken = &s
// 		}
// 	}
// 	return "https://api.telegram.org/bot" + *botToken
// }

// func sendMessage(chatID int, text string) error {
// 	sendURL := botURL() + "/sendMessage"
// 	v := url.Values{}
// 	v.Set("chat_id", strconv.Itoa(chatID))
// 	v.Set("text", text)
// 	_, err := http.PostForm(sendURL, v)
// 	return err
// }
