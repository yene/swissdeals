package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
)

type offer struct {
	Name  string
	Price string
	Link  string
	Image string
	Site  string
}

var slacksHook = flag.String("slack", "", "Slacks Webhook")

func main() {
	flag.Parse()

	if *slacksHook == "" {
		fmt.Println("Please pass in a slack Webhook")
		fmt.Println("If you don't have one visit https://my.slack.com/services/new/incoming-webhook/")
		return
	}

	offers := make(chan offer)

	go func() {
		for o := range offers {
			postOffer(o)
		}
	}()
	DaydealScrape(offers)
	GalaxusScrape(offers)
	DigitecScrape(offers)

	// genius or just lazy?
	time.Sleep(time.Second * 2)
}

func postOffer(p offer) {
	text := fmt.Sprintf("<%s|%s> für nur %s <%s|Image>", p.Link, p.Name, p.Price, p.Image)

	payload := slack.Payload{
		Text:     text,
		Username: p.Site,
		Channel:  "#spam",
		IconUrl:  "https://shiro.ch/mymoney.png",
	}
	err := slack.Send(*slacksHook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}

/*
 * Helper functions
 */

func removeNewline(s string) string {
	s = strings.Replace(s, "\n", " ", -1)
	s = strings.TrimSpace(s)
	return s
}

func removeShitFromURL(url string) string {
	que := strings.Index(url, "?")
	if que == -1 {
		return url
	}
	return url[0:que]
}

func cutoffafterprice(p string) string {
	pos := strings.LastIndex(p, ".–")
	if pos == -1 {
		return p
	}
	pos += len(".–") // The dash has length of 3
	return p[:pos]
}

func extractSrc(s string) string {
	start := strings.Index(s, "src=\"")
	if start == -1 {
		return ""
	}
	start += 5
	substring := s[start:len(s)]

	end := strings.Index(substring, "\"")
	if start == -1 {
		return ""
	}
	return substring[0:end]
}
