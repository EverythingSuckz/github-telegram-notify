# Github Telegram Notify

Actions that sends commit updates of a repository to any chat in Telegram

## Prerequisites

### - `bot_token`
A Telegram Bot Token is required for using the Telegram bot from which the commit updates are being send.
### Obtaining a Telegram Bot Token
 - Goto  [@BotFather](https://telegram.dog/BotFather)
 - After sending `/start` command, send `/newbot`
 - Follow the onscreen instructions and at the end, you'll obtain a bot token.

### - `chat_id`
### Obtaining a Telegram Bot Token
 - Go to the group of your choice  
 - Add [@MissRose_bot](https://telegram.dog/MissRose_bot)
 - Type the command `/id` and send it to the group.

Add the following lines of code in your YML file.

```sh
  - uses: actions/checkout@v3
  - name: Notify the commit on Telegram
    uses: EverythingSuckz/github-telegram-notify@v1.0.0
    with:
      bot_token: '${{ secrets.BOT_TOKEN }}'
      chat_id: '${{ secrets.CHAT_ID }}'
```

## Contributing
Feel free to open a PR in case of any minor fix and please open an issue first for major changes.

## License

<a href="https://opensource.org/licenses/MIT"><img src="https://upload.wikimedia.org/wikipedia/commons/0/0c/MIT_logo.svg" alt="mit" width="150"/></a>   
[Licensed under MIT License](https://opensource.org/licenses/MIT) 