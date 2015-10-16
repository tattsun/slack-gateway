package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

type gatewayCtx struct {
	accessToken string
	slack       *slack.Client
	port        string
}

var ctx = gatewayCtx{}

func initGatewayCtx() {
	slackToken := os.Getenv("SLACK_GATEWAY_SLACK_TOKEN")
	if slackToken == "" {
		panic("SLACK_GATEWAY_SLACK_TOKEN isn't set")
	}
	ctx.slack = slack.New(slackToken)

	ctx.accessToken = os.Getenv("SLACK_GATEWAY_ACCESS_TOKEN")
	if ctx.accessToken == "" {
		panic("SLACK_GATEWAY_ACCESS_TOKEN isn't set")
	}

	ctx.port = os.Getenv("SLACK_GATEWAY_PORT")
	if ctx.port == "" {
		ctx.port = ":8080"
	}
}

func authorize(w http.ResponseWriter, r *http.Request) bool {
	token := r.FormValue("accessToken")
	if token != ctx.accessToken {
		log.WithFields(log.Fields{
			"accessToken": token,
		}).Error("unauthorized")
		http.Error(w, "unauthorized", 401)
		return false
	}
	return true
}

func postToSlack(username string, channel string, message string, attachments string) {
	param := slack.NewPostMessageParameters()
	param.Username = username

	dec := json.NewDecoder(strings.NewReader(attachments))
	_ = dec.Decode(&param.Attachments)
	fmt.Printf("%v\n", param.Attachments)

	_, _, err := ctx.slack.PostMessage(channel, message, param)
	if err != nil {
		log.WithFields(log.Fields{
			"username":    username,
			"channel":     channel,
			"message":     message,
			"attachments": attachments,
		}).Error(err.Error())
		return
	}
	log.WithFields(log.Fields{
		"username":    username,
		"channel":     channel,
		"message":     message,
		"attachments": attachments,
	}).Info("posted")
}

func handler(w http.ResponseWriter, r *http.Request) {
	ok := authorize(w, r)
	if !ok {
		return
	}

	username := r.FormValue("username")
	channel := r.FormValue("channel")
	message := r.FormValue("message")
	attachments := r.FormValue("attachments")
	go func() {
		postToSlack(username, channel, message, attachments)
	}()
	fmt.Fprintf(w, "ok")
}

func main() {
	initGatewayCtx()
	http.HandleFunc("/", handler)
	http.ListenAndServe(ctx.port, nil)
}
