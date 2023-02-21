package youtube

import (
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

func GetSubscribers() (Item, error) {
	var response Response

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/" , nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	q := req.URL.Query()
	// notice how I'm using os.Getenv() to pick up the environment
    // variables that we defined earlier. No hard coded credentials here
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

    // finally we make the request to the URL that we have just
    // constructed
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return Item{}, err
    }
    defer resp.Body.Close()

}
