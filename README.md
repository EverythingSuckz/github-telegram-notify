# github-telegram-notif

```sh
- uses: actions/checkout@v3
- name: Notify the commit on Telegram
    uses: EverythingSuckz/github-telegram-notify@v1
    with:
    bot_token: '${{ secrets.BOT_TOKEN }}'
    chat_id: '${{ secrets.CHAT_ID }}'
```