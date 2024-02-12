package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *channel `xml:"channel"` //pointer bcz we have only one channel
}

type channel struct {
	Title    string  `xml:"title"`
	ItemList []Items `xml:"item"`
}
type Items struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\n Below are all the Google Search Trends for Today")
	fmt.Println("----------------------------------------------------")

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("search Term: ", r.Channel.ItemList[i].Title)
		fmt.Println("Link to the trend: ", r.Channel.ItemList[i].Link)
		fmt.Println("Headline: ", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("Headline: ", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)

		fmt.Println("----------------------------------------------------")
	}

}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}
