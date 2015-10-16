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
- [/attachments] Post to slack using attackments
    - params
        - accessToken(string) - equals to SLACK_GATEWAY_ACCESS_TOKEN
        - username(string)
        - channel(string) - not channel name but ID
        - attachments(json) - see [Attachment Example](#attachmentex)

### <a name="attachmentex"></a>Attachment Example

See [Slack API Document](https://api.slack.com/docs/attachments) for more information.

```
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
