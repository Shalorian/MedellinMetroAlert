package mediator

import (
	"log"
	"regexp"
	"time"

	"github.com/MedellinMetroAlert/utils"
	"github.com/dghubble/go-twitter/twitter"
)

const (
	timeformat = "Mon Jan 2 15:04:05 +0000 2006"
)

var (
	emojiRx = regexp.MustCompile(`[\x{1F4E2}]`)
)

type Event struct {
	twitter.Tweet
	Events *[]twitter.Tweet
}

func (event Event) FetchEvents(tweets *[]twitter.Tweet) error {
	found := false
	for _, tweet := range *tweets {
		createdAt, err := time.Parse(timeformat, tweet.CreatedAt)
		if err != nil {
			log.Printf("Time format error :: %s \n", err.Error())
			return err
		}

		if emojiRx.MatchString(tweet.Text) && utils.IsRelevant(createdAt) && event.isNew(tweet.Text) {
			found = true
			*event.Events = append(*event.Events, tweet)
			log.Printf("%s \n Check https://twitter.com/metrodemedellin/status/%v for more details", tweet.Text, tweet.ID)
		}
	}

	if !found {
		log.Println("No new events found.")
	}

	return nil
}

func (event Event) isNew(newEvent string) bool {
	for _, oldEvent := range *event.Events {
		if oldEvent.Text == newEvent {
			return false
		}
	}
	return true
}
