package supervisor

import (
	"fmt"
	"log"
	"time"

	"github.com/MedellinMetroAlert/supervisor/mediator"
	"github.com/dghubble/go-twitter/twitter"
)

const (
	user             = "metrodemedellin"
	requestFrequency = time.Second * 20
)

func Start() {
	log.Println("Initializing event cache ...")
	event := mediator.Event{
		Events: &[]twitter.Tweet{},
	}

	for {
		fmt.Println()
		log.Printf("Searching for new events from @%s... \n", user)

		tweets, err := mediator.GetTweetsFromUserTimeline(user)
		if err != nil {
			log.Printf("Error while getting tweets from @%s :: %s \n Retrying ... \n", user, err.Error())
			continue
		}

		err = event.FetchEvents(tweets)
		if err != nil {
			log.Printf("Error while fetching events from @%s tweets :: %s \n Retrying ... \n", user, err.Error())
			continue
		}
		time.Sleep(requestFrequency)
	}
}
