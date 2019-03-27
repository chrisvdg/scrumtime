# Scrumtime

Simple chat bot that can be used to announce scrum time by sending a message to a chat room with a defined schedule.

Supported chat platforms:
 - [Slack](https://slack.com/)
 - [Telegram](https://telegram.org/)

## Usage

Create a yaml file called `config.yaml` and insert the following data for each job:

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
    schedule: '0 0 0 * * 1-5' # More info on format https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format
```

Run
```sh
# this will take `./config.yaml` as config file by default
go run main.go
```

Alternatively to explicitly add the config file:
```sh
go run main.go -c <path/to/file>/myconfig.yaml
```

### Docker image

This project has a Dockerfile to create a small docker image of this project.

Generate image:
```sh
docker build -t scrumtime .
```

Specify Go version (Go version to build with)
```sh
docker build -t scrumtime . --build-arg go_version=1.11
```

Specify Alpine version (Final image base into which the binary is copied)
```sh
docker build -t scrumtime . --build-arg alpine_version=3.8.4
```


Run container with image:
```
docker create --name scrumtime_helloworld scrumtime
docker cp config.yaml scrumtime_helloworld:/
docker start scrumtime_helloworld
```
