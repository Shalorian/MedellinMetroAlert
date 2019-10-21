package mediator

import (
	"errors"
	"log"

	"github.com/MedellinMetroAlert/utils"
	"github.com/dghubble/go-twitter/twitter"
)

func GetTweetsFromUserTimeline(screenName string) (*[]twitter.Tweet, error) {
	trimUser := true
	excludeReplies := true
	includeRetweets := true

	timelineParams := &twitter.UserTimelineParams{
		ScreenName:      screenName,
		TrimUser:        &trimUser,
		ExcludeReplies:  &excludeReplies,
		IncludeRetweets: &includeRetweets,
		Count:           100,
	}

	client := utils.Twitterclient

	log.Printf("Getting tweets from @%s ...\n", screenName)
	tweets, resp, err := client.Timelines.UserTimeline(timelineParams)
	if err != nil {
		log.Printf("Twitter Api response error %s\n", err.Error())
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Println("Twitter Api response not successfull")
		return nil, errors.New("Response not successfull")
	}

	return &tweets, nil
}
