# scrumtime
Simple slack bot announcing scrum time

## Usage

Create a yaml file called `config.yaml` and insert the following data for each job:

* channel: Channel name to sent message too
* message: The message you want the bot to send
* api_key: Slack OAuth Access Token
* schedule: Cron schedule  
    e.g: '0 0 * * 1-5': This send the message every weekday at midnight  
    More info https://godoc.org/github.com/robfig/cron

Config example:

```yaml
schedules:
  a_job: # Name of your job to schedule
    channel: 'my channel'
    message: 'Hello world'
    api_key: 'xoxp-38811....'
    schedule: '0 0 * * 1-5'
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
