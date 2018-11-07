# Scrumtime
Simple slack bot that can be used to announce scrum time

## Usage

Create a yaml file called `config.yaml` and insert the following data for each job:

* channel: Channel name to sent message too
* message: The message you want the bot to send
* api_key: Slack OAuth Access Token.  
Register app and add app to workspace at: https://api.slack.com/apps  
Add a bot user.  
Then get the key at `Features/OAuth & Permissions`  
* schedule: Cron schedule  
    e.g: '0 0 0 * * 1-5': This send the message every weekday at midnight  
    More info https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format

Config example:

```yaml
messengers:
  ex_slack:
    platform: slack
    api_key: 'xoxp-388...123'
    chat_id: test_channel
  ex_telegram:
    platform: telegram
    api_key: '123456789:AAF...E_Y'
    chat_id: '-123456789'

schedules:
  example:
    message: 'Hello world!'
    messengers:
      - ex_slack
      - ex_telegram
    schedule: '0 0 0 * * 1-5'
```

Run
```sh
# this will take `./config.yaml` as config file per default
go run main.go
```

Alternatively to explicitly add the config file:
```sh
go run main.go -c <path/to/file>/myconfig.yaml
```
