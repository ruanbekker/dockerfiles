package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey       = Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = Getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = Getenv("TWITTER_ACCESS_TOKEN_SECRET")
  keyword           = Getenv("TWITTER_KEYWORD")
)

func Getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing envvar " + name)
	}
	return v
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

    stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#" + keyword},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			logrus.Warning("received unexpected error %T", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}
		_, err := api.Retweet(t.Id, false)
		if err != nil {
			logrus.Errorf("could not retweet %d: %v", t.Id, err)
			continue
		}
		logrus.Infof("retweetedId: %d, Tweet: %v", t.Id, t.FullText)
	}

}
