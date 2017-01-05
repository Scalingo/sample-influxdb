# Sample InfluxDB

This sample count the number of tweets for a given hashtag every minute.

This sample runs a worker to update the InfluxDB database and a web container to display a graphic
showing the results.

Mandatory environment variable:

- HASHTAG: which hashtag to follow
- TWITTER_ACCESS_SECRET, TWITTER_ACCESS_TOKEN, TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET:
  credentials to access the Twitter API. Generate these variables [here](https://apps.twitter.com/).

