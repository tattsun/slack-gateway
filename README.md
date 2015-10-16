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
        - attachments(json, nullable) - see [Attachments Example](#attachmentsex)

### <a name="attachmentsex"></a>Attachments Example

See [Slack API Document](https://api.slack.com/docs/attachments) for more information.

```json
[
   {
      "fallback":"New open task [Urgent]: <http://url_to_task|Test out Slack message attachments>",
      "pretext":"New open task [Urgent]: <http://url_to_task|Test out Slack message attachments>",
      "color":"#D00000",
      "fields":[
         {
            "title":"Notes",
            "value":"This is much easier than I thought it would be.",
            "short":false
         }
      ]
   }
]
```
