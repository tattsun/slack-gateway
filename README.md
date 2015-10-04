# slack-gateway

Simple slack gateway server.

## Install

```
$ go get github.com/tattsun/slack-gateway
$ SLACK_GATEWAY_SLACK_TOKEN=<your-slack-token> SLACK_GATEWAY_ACCESS_TOKEN=<gateway-token> slack-gateway
```

## API

- [/] Post to slack
	- params
		- accessToken(string) - equals to SLACK_GATEWAY_ACCESS_TOKEN
		- username(string)
		- channel(string) - not channel name but ID
		- message(string)
