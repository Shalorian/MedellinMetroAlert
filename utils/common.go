package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/spf13/viper"
)

var (
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
	Twitterclient  *twitter.Client
)

func InitializeConfigUtils() {
	ex, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("Error while returning executable path: %s", err))
	}
	confPath := filepath.Join(ex, "conf/"+string(filepath.Separator))

	viper.SetConfigType("yaml")
	viper.SetConfigName("app_conf")
	viper.AddConfigPath(confPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error while reading config file: %s", err))
	}

	consumerKey = viper.GetString("consumerKey")
	consumerSecret = viper.GetString("consumerSecret")
	accessToken = viper.GetString("accessToken")
	accessSecret = viper.GetString("accessSecret")

	newTwitterClient()
}

func newTwitterClient() {
	// utils.GetTweetsFromUser("metrodemedellin", true, false, true)
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	Twitterclient = twitter.NewClient(httpClient)
}

func IsRelevant(createdAt time.Time) bool {
	return time.Now().Day() == createdAt.Day() && time.Now().Month() == createdAt.Month() && time.Now().Year() == createdAt.Year()
}
