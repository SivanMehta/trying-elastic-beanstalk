package twitter

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// TBD: getting actual tweets from Twitter
var accessToken = ""
var token = ""
var secret = ""

const apiBase = "https://api.twitter.com/"

//
// GetTweets will return a body of text from trending topics
// This is currently just the text of the Declaration of Independence, but stay tuned.
//
func GetTweets(trend string) []string {
	// exactly containing the trend, without tweets that contain links
	query := "\"" + trend + "\" -filter:links"
	query = url.QueryEscape(query)
	responseString := makeAuthedRequest("GET", "1.1/search/tweets.json?q="+query+"&include_entities=false&lang=en&result_type=popular")

	fmt.Println(string(responseString))

	totallyTweets, _ := ioutil.ReadFile("corpus.txt")

	tweets := strings.Split(string(totallyTweets), ".")
	return tweets
}

//
// GetTrends goes to twitter and returns a list of trending topics in the US
//
func GetTrends() []string {
	responseString := makeAuthedRequest("GET", "1.1/trends/place.json?id=1")
	fmt.Println(string(responseString))

	// would normally be fetched from twitter, just hardcoded for now
	return []string{"nba", "dogs", "marvel"}
}
