<h1 align="center">GitHub Telegram Notify</h1>
<h4 align="center">Actions that sends commit updates of a repository to any chat in Telegram</h4>
<p align="center"><a href="https://github.com/marketplace/actions/github-telegram-notifier">🏪 View on Github Marketplace</a>&emsp;🏷️ <a href="https://github.com/EverythingSuckz/github-telegram-notify/releases/tag/v1.1.2">v1.1.2</a></p>

## Prerequisites

### - `bot_token`
A Telegram Bot Token is required for using the Telegram bot from which the commit updates are being send.
### Obtaining a Telegram Bot Token
 - Goto  [@BotFather](https://telegram.dog/BotFather)
 - After sending `/start` command, send `/newbot`
 - Follow the onscreen instructions and at the end, you'll obtain a bot token.

### - `chat_id`
### Obtaining Chat ID of a group
 - Go to the group of your choice  
 - Add [@MissRose_bot](https://telegram.dog/MissRose_bot)
 - Type the command `/id` and send it to the group.

### - `topic_id` (optional)
Use this only if you have topics enabled.

## How to use?

### Basic example

Add the following lines of code in your YML file.

```sh
  - name: Notify the commit on Telegram
    uses: EverythingSuckz/github-telegram-notify@main
    with:
      bot_token: '${{ secrets.TELEGRAM_BOT_TOKEN }}'
      chat_id: '${{ secrets.TELEGRAM_CHAT_ID }}'
```

### Advanced example

You can also replace the content of your YML file with the following lines of code.


```sh
name: Telegram Notification
on:
  push:
  fork:
  watch:
  issues:
    types: [created, closed, opended, reopened, locked, unlocked]
  issue_comment:
    types: [created, deleted]
  pull_request:
    types: [created, closed, opened, reopened, locked, unlocked, synchronize]
  pull_request_target:
    types: [created, closed, opened, reopened, locked, unlocked, synchronize]
  pull_request_review_comment:
    types: [created, deleted]
  release:
    types: [published, released]

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
        - uses: actions/checkout@v2
        - name: Notify everything on Telegram
          uses: EverythingSuckz/github-telegram-notify@main
          with:
            bot_token: '${{ secrets.TELEGRAM_BOT_TOKEN }}'
            chat_id: '${{ secrets.TELEGRAM_CHAT_ID }}'
            topic_id: '${{ secrets.TELEGRAM_TOPIC_ID }}'
```

## Supported events

- [x] Commits
- [x] Forks
- [x] Watch
  - [x] stars
- [x] Issues
  - [x] created
  - [x] closed
  - [x] opened
  - [x] reopened
  - [x] locked
  - [x] unlocked
- [x] Issue comments
  - [x] created
  - [x] deleted
- [x] Pull Request
  - [x] created
  - [x] closed
  - [x] opened
  - [x] reopened
  - [x] locked
  - [x] unlocked
  - [x] synchronize
- [x] Pull Request comments
  - [x] created
  - [x] deleted
- [x] Releases
  - [x] published
  - [x] released
- [ ] Discussions
- [ ] Discussion comments

## Contributing

If you can't find the event you are looking for, assume that it's not tested yet. You are free to open a pull request.

Also, feel free to open a PR in case of any minor fix and please open an issue first for major changes.

## License

<a href="https://opensource.org/licenses/MIT"><img src="https://upload.wikimedia.org/wikipedia/commons/0/0c/MIT_logo.svg" alt="mit" width="150"/></a>   
[Licensed under MIT License](https://opensource.org/licenses/MIT) 