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

Add a BOT with `app_mentions:read` and `chat:write` permissions to your Slack workspace and set .env.

```
docker-compose up --build
```

## Deploy

I recommend using [Cloud Run](https://cloud.google.com/run) provided by GCP.  
You just select the forked-repository, set two environment variables, and set an endpoint for the Slack event API.
