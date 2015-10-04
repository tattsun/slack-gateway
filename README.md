# slack-gateway

Simple slack gateway server.

## Install

```
$ go get github.com/tattsun/slack-gateway
$ SLACK_GATEWAY_SLACK_TOKEN=<your-slack-token> SLACK_GATEWAY_ACCESS_TOKEN=<gateway-token> slack-gateway
```

## Environment Variables

- SLACK_GATEWAY_PORT(string/optional)
  - default: ":8080"
- SLACK_GATEWAY_SLACK_TOKEN(string/required)
- SLACK_GATEWAY_ACCESS_TOKEN(string/required)

## API

- [/] Post to slack
	- params
		- accessToken(string) - equals to SLACK_GATEWAY_ACCESS_TOKEN
		- username(string)
		- channel(string) - not channel name but ID
		- message(string)
