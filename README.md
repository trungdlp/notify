<div align="center">
<img
    width=40%
    src="assets/gopher-letter.svg"
    alt="notify logo"
/>

---

<p><i>A dead simple Go library for sending notifications to various messaging services.</i></p>

&nbsp;

[![CI](https://github.com/trungdlp/notify/actions/workflows/ci.yml/badge.svg)](https://github.com/trungdlp/notify/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/trungdlp/notify)](https://goreportcard.com/report/github.com/trungdlp/notify)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/trungdlp/notify)

</div>

&nbsp;

## About <a id="about"></a>

This repository is an actively maintained fork of [nikoksr/notify](https://github.com/nikoksr/notify). It preserves Notify's simple API for sending notifications across different messaging platforms while keeping dependencies current, replacing unmaintained client libraries, and maintaining service integrations.

## Disclaimer <a id="disclaimer"></a>

Any misuse of this library is your own liability and responsibility and cannot be attributed to the authors of this library.  See [license](LICENSE) for more.

Spamming through the use of this library **may get you permanently banned** on most supported platforms.

Since Notify is highly dependent on the consistency of the supported external services and the corresponding latest client libraries, we cannot guarantee its reliability nor its consistency, and therefore you should probably not use or rely on Notify in critical scenarios.

## Install <a id="install"></a>

```sh
go get -u github.com/trungdlp/notify
```

## Example usage <a id="usage"></a>

```go
// Create a telegram service. Ignoring error for demo simplicity.
telegramService, _ := telegram.New("your_telegram_api_token")

// Passing a telegram chat id as receiver for our messages.
// Basically where should our message be sent?
telegramService.AddReceivers(-1234567890)

// Optional: send messages to a specific forum topic (message thread).
telegramService.SetMessageThreadID(12345)

// Tell our notifier to use the telegram service. You can repeat the above process
// for as many services as you like and just tell the notifier to use them.
// Inspired by http middlewares used in higher level libraries.
notify.UseServices(telegramService)

// Send a test message.
_ = notify.Send(
	context.Background(),
	"Subject/Title",
	"The actual message - Hello, you awesome gophers! :)",
)
```

#### Recommendation <a id="recommendation"></a>

In this example, we use the global `Send()` function. Similar to most logging libraries such as
[zap](https://github.com/uber-go/zap), we provide global functions for convenience. However, as with most logging
libraries, we also recommend avoiding the use of global functions as much as possible. Instead, use one of our versatile
constructor functions to create a new local `Notify` instance and pass it down the stream.

## Contributing <a id="contributing"></a>

Contributions of all kinds are welcome. Check the [open issues](https://github.com/trungdlp/notify/issues), read the [contribution guidelines](https://github.com/trungdlp/notify/blob/main/CONTRIBUTING.md), or open a pull request to help keep the library and its service integrations up to date.

> Psst, don't forget to check the list of [missing services](https://github.com/trungdlp/notify/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Aaffects%2Fservices+label%3A%22help+wanted%22+no%3Aassignee) waiting to be added by you or create [a new issue](https://github.com/trungdlp/notify/issues/new?assignees=&labels=affects%2Fservices%2C+good+first+issue%2C+hacktoberfest%2C+help+wanted%2C+type%2Fenhancement%2C+up+for+grabs&template=service-request.md&title=feat%28service%29%3A+Add+%5BSERVICE+NAME%5D+service) if you want a new service to be added.

## Supported services <a id="supported_services"></a>

> Click [here](https://github.com/trungdlp/notify/issues/new?assignees=&labels=affects%2Fservices%2C+good+first+issue%2C+hacktoberfest%2C+help+wanted%2C+type%2Fenhancement%2C+up+for+grabs&template=service-request.md&title=feat%28service%29%3A+Add+%5BSERVICE+NAME%5D+service) to request a missing service.

| Service                                                                           | Path                                     | Credits                                                                                         |       Status       |
|-----------------------------------------------------------------------------------|------------------------------------------|-------------------------------------------------------------------------------------------------|:------------------:|
| [Amazon SES](https://aws.amazon.com/ses)                                          | [service/amazonses](service/amazonses)   | [aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)                                       | :heavy_check_mark: |
| [Amazon SNS](https://aws.amazon.com/sns)                                          | [service/amazonsns](service/amazonsns)   | [aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)                                       | :heavy_check_mark: |
| [Bark](https://apps.apple.com/us/app/bark-customed-notifications/id1403753865)    | [service/bark](service/bark)             | -                                                                                               | :heavy_check_mark: |
| [DingTalk](https://www.dingtalk.com)                                              | [service/dinding](service/dingding)      | [blinkbean/dingtalk](https://github.com/blinkbean/dingtalk)                                     | :heavy_check_mark: |
| [Discord](https://discord.com)                                                    | [service/discord](service/discord)       | [bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)                                     | :heavy_check_mark: |
| [Email](https://wikipedia.org/wiki/Email)                                         | [service/mail](service/mail)             | [jordan-wright/email](https://github.com/jordan-wright/email)                                   | :heavy_check_mark: |
| [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging)      | [service/fcm](service/fcm)               | [appleboy/go-fcm](https://github.com/appleboy/go-fcm)                                           | :heavy_check_mark: |
| [Google Chat](https://workspace.google.com/intl/en/products/chat/)                | [service/googlechat](service/googlechat) | [googleapis/google-api-go-client](https://google.golang.org/api/chat/v1)                        | :heavy_check_mark: |
| [HTTP](https://wikipedia.org/wiki/Hypertext_Transfer_Protocol)                    | [service/http](service/http)             | -                                                                                               | :heavy_check_mark: |
| [Lark](https://www.larksuite.com/)                                                | [service/lark](service/lark)             | [go-lark/lark](https://github.com/go-lark/lark)                                                 | :heavy_check_mark: |
| [Line](https://line.me)                                                           | [service/line](service/line)             | [line/line-bot-sdk-go](https://github.com/line/line-bot-sdk-go)                                 | :heavy_check_mark: |
| [Line Notify](https://notify-bot.line.me)                                         | [service/line](service/line)             | [utahta/go-linenotify](https://github.com/utahta/go-linenotify)                                 | :heavy_check_mark: |
| [Mailgun](https://www.mailgun.com)                                                | [service/mailgun](service/mailgun)       | [mailgun/mailgun-go](https://github.com/mailgun/mailgun-go)                                     | :heavy_check_mark: |
| [Matrix](https://www.matrix.org)                                                  | [service/matrix](service/matrix)         | [mautrix/go](https://github.com/mautrix/go)                                                     | :heavy_check_mark: |
| [Microsoft Teams](https://www.microsoft.com/microsoft-teams)                      | [service/msteams](service/msteams)       | [atc0005/go-teams-notify](https://github.com/atc0005/go-teams-notify)                           | :heavy_check_mark: |
| [PagerDuty](https://www.pagerduty.com)                                            | [service/pagerduty](service/pagerduty)   | [PagerDuty/go-pagerduty](https://github.com/PagerDuty/go-pagerduty)                             | :heavy_check_mark: |
| [Plivo](https://www.plivo.com)                                                    | [service/plivo](service/plivo)           | [plivo/plivo-go](https://github.com/plivo/plivo-go)                                             | :heavy_check_mark: |
| [Pushover](https://pushover.net/)                                                 | [service/pushover](service/pushover)     | [gregdel/pushover](https://github.com/gregdel/pushover)                                         | :heavy_check_mark: |
| [Pushbullet](https://www.pushbullet.com)                                          | [service/pushbullet](service/pushbullet) | [cschomburg/go-pushbullet](https://github.com/cschomburg/go-pushbullet)                         | :heavy_check_mark: |
| [Reddit](https://www.reddit.com)                                                  | [service/reddit](service/reddit)         | [vartanbeno/go-reddit](https://github.com/vartanbeno/go-reddit)                                 | :heavy_check_mark: |
| [RocketChat](https://rocket.chat)                                                 | [service/rocketchat](service/rocketchat) | [RocketChat/Rocket.Chat.Go.SDK](https://github.com/RocketChat/Rocket.Chat.Go.SDK)               | :heavy_check_mark: |
| [SendGrid](https://sendgrid.com)                                                  | [service/sendgrid](service/sendgrid)     | [sendgrid/sendgrid-go](https://github.com/sendgrid/sendgrid-go)                                 | :heavy_check_mark: |
| [Slack](https://slack.com)                                                        | [service/slack](service/slack)           | [slack-go/slack](https://github.com/slack-go/slack)                                             | :heavy_check_mark: |
| [Syslog](https://wikipedia.org/wiki/Syslog)                                       | [service/syslog](service/syslog)         | [log/syslog](https://pkg.go.dev/log/syslog)                                                     | :heavy_check_mark: |
| [Telegram](https://telegram.org)                                                  | [service/telegram](service/telegram)     | [OvyFlash/telegram-bot-api](https://github.com/OvyFlash/telegram-bot-api)                       | :heavy_check_mark: |
| [TextMagic](https://www.textmagic.com)                                            | [service/textmagic](service/textmagic)   | [textmagic/textmagic-rest-go-v2](https://github.com/textmagic/textmagic-rest-go-v2)             | :heavy_check_mark: |
| [Twilio](https://www.twilio.com/)                                                 | [service/twilio](service/twilio)         | [kevinburke/twilio-go](https://github.com/kevinburke/twilio-go)                                 | :heavy_check_mark: |
| [Twitter](https://twitter.com)                                                    | [service/twitter](service/twitter)       | [drswork/go-twitter](https://github.com/drswork/go-twitter)                                     | :heavy_check_mark: |
| [Viber](https://www.viber.com)                                                    | [service/viber](service/viber)           | [mileusna/viber](https://github.com/mileusna/viber)                                             | :heavy_check_mark: |
| [WeChat](https://www.wechat.com)                                                  | [service/wechat](service/wechat)         | [silenceper/wechat](https://github.com/silenceper/wechat)                                       | :heavy_check_mark: |
| [Webpush Notification](https://developer.mozilla.org/en-US/docs/Web/API/Push_API) | [service/webpush](service/webpush)       | [SherClockHolmes/webpush-go](https://github.com/SherClockHolmes/webpush-go/)                    | :heavy_check_mark: |
| [WhatsApp](https://www.whatsapp.com)                                              | [service/whatsapp](service/whatsapp)     | [Rhymen/go-whatsapp](https://github.com/Rhymen/go-whatsapp)                                     |        :x:         |

## Special Thanks <a id="special_thanks"></a>

### Original project <a id="original-project"></a>

Notify was originally created by [@nikoksr](https://github.com/nikoksr). This project is based on the original [nikoksr/notify](https://github.com/nikoksr/notify) repository and is grateful for the foundation built by its original author and contributors.

### Maintainers <a id="maintainers"></a>

- [@svaloumas](https://github.com/svaloumas)

### Logo <a id="logo"></a>

The [logo](https://github.com/MariaLetta/free-gophers-pack) was made by the amazing [MariaLetta](https://github.com/MariaLetta).
