# Sample InfluxDB

This sample count the number of tweets for a given hashtag every minute.

This sample runs a worker to update the InfluxDB database and a web container to display a graphic
showing the results.

Mandatory environment variable:

- HASHTAG: which hashtag to follow
- TWITTER_ACCESS_SECRET, TWITTER_ACCESS_TOKEN, TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET:
  credentials to access the Twitter API. Generate these variables [here](https://apps.twitter.com/).

This sample is running on: https://influxdb.is-easy-on-scalingo.com/

## Deploy via Git

Create an application on https://scalingo.com, then:

```shell
scalingo --app my-app git-setup
git push scalingo master
```

And that's it!

## Deploy via One-Click

[![Deploy to Scalingo](https://cdn.scalingo.com/deploy/button.svg)](https://dashboard.scalingo.com/create/app?source=https://github.com/Scalingo/sample-influxdb#master)

## Running Locally

```shell
docker compose up
```

The app listens by default on the port 8086