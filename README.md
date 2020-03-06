# Scrumtime

Simple chat bot that can be used to announce scrum time by sending a message to a chat room with a defined schedule.

Supported chat platforms:
 - [Slack](https://slack.com/)
 - [Telegram](https://telegram.org/)

## Usage

Create a yaml file called `config.yaml` and insert data according the the following structure

Config example:

```yaml
bots:
  ex_slack:
    platform: slack
    api_key: 'xoxp-388...123'
  ex_telegram:
    platform: telegram
    api_key: '123456789:AAF...E_Y'

messages:
  example:
    body: 'Hello world!'
    messengers:
      - bot: ex_slack
        chat_ids: test_channel
      - bot: ex_telegram
        chat_ids: ['-123456789']
    schedule: '0 0 0 * * 1-5' # More info on format https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format
    expiretime: 600 # Optional time in seconds to delete the message again
```

### Schedule
Schedule is in a Quarts Scheduler format, [more info can he found here.](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html)

Run
```sh
# this will take `./config.yaml` as config file by default
go run main.go
```

Alternatively to explicitly add the config file:
```sh
go run main.go -c <path/to/file>/myconfig.yaml
```

## Docker
### Build Docker image

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

### Pull docker image

Alternatively, the image pushed on Docker Hub can be used.
This command also mounts (`-v`) the config file into the container instead of copying to it  
and runs it in the background (`-d`).

```sh
docker run --name scrumtime_helloworld -d -v $PWD/config.yaml:/config.yaml chrisvdg/scrumtime # Or with version tag: chrisvdg/scrumtime:0.0.1
```

### Setting timezone of container

By default the image uses the UTC timezone. To set the timezone of your container, set the `TZ` environmental variable.  
Using this, you can't set it with abbreviated, use the full name of the `TZ` Database. (Timezone may be changed but time will still be UTC)

```sh
docker run -e "TZ=Europe/Brussels" -d -v $PWD/config.yaml:/config.yaml chrisvdg/scrumtime
```

The timezone can also be configured in the Dockerfile so there won't be a need to set it when starting the container. For this the abbreviations can be used.

Replace the line `RUN apk add tzdata` with:
Replace `EST` (2x) with your timezone.
```Dockerfile
RUN apk add tzdata && cp /usr/share/zoneinfo/EST /etc/localtime && echo "EST" >  /etc/timezone && apk del tzdata
```
