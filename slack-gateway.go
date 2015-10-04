package main

import (
	"fmt"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

func postToSlack(username string, channel string, message string) {
	param := slack.NewPostMessageParameters()
	param.Username = username
	_, _, err := ctx.slack.PostMessage(channel, message, param)
	if err != nil {
		log.WithFields(log.Fields{
			"username": username,
			"channel":  channel,
			"message":  message,
		}).Error(err.Error())
		return
	}
	log.WithFields(log.Fields{
		"username": username,
		"channel":  channel,
		"message":  message,
	}).Info("posted")
}

func handler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("accessToken")
	if token != ctx.accessToken {
		log.WithFields(log.Fields{
			"accessToken": token,
		}).Error("unauthorized")
		http.Error(w, "unauthorized", 400)
		return
	}

	username := r.FormValue("username")
	channel := r.FormValue("channel")
	message := r.FormValue("message")
	go func() {
		postToSlack(username, channel, message)
	}()
	fmt.Fprintf(w, "ok")
}

func main() {
	initGatewayCtx()
	http.HandleFunc("/", handler)
	http.ListenAndServe(ctx.port, nil)
}
