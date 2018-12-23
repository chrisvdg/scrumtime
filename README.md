# Scrumtime
Simple slack bot that can be used to announce scrum time

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

Run container with image:
```
docker create --name scrumtime_helloworld scrumtime
docker cp config.yaml scrumtime_helloworld:/
docker start scrumtime_helloworld
```
