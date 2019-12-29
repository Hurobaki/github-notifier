# Github Notifier

Small server written in Go to send message to Slack workspace whenever an action is made to a pull request.

## TODO

* Script to automate deployment to Now.sh
* Handle application/x-www-form-urlencoded

## Documentation

[Github Api](https://developer.github.com/v3/activity/events/types/#pullrequestevent)
[Slack Block Builder](https://api.slack.com/tools/block-kit-builder?mode=message&blocks=%5B%5D)

## Develop / Deploy

```shell script
now dev

now --prod
```

## Troubleshooting

Use Ngrok to send Github payload instead of Postman. It leads to copying issue. 