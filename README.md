# Slack Timer

## Commands

`@bot` is your BOT's name.

### `@bot ping`

Check to see if the BOT server is working properly.  
If it is working, the BOT posts the message `pong` to the same channel.

### `@bot timer`

Measure the time.  
Valid arguments are the following.

- `@bot timer 3`: Measures 3 seconds.
- `@bot timer 3 sec`: Measures 3 seconds.
- `@bot timer 3 min`: Measures 3 minutes.
- `@bot timer 3 sec Hello, world!`: Measures 3 seconds with the memo: "Hello, world!".

## Develop

Add a BOT with `app_mentions:read` and `chat:write` permissions to your Slack channel and set .env.

```
docker-compose up --build
```