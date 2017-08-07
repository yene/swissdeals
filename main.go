package main

import (
	"flag"
	"fmt"
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
	DigitecGalaxusScrape(offers, "digitec")
	DigitecGalaxusScrape(offers, "galaxus")

	// genius or just lazy?
	time.Sleep(time.Second * 2)
}

func postOffer(p offer) {
	var text string
	if p.Image == "" { // digitec/galaxus
		text = fmt.Sprintf("*%s* %s", p.Price, p.Link)
	} else { // daydeal
		text = fmt.Sprintf("<%s|%s> für nur %s <%s|Image>", p.Link, p.Name, p.Price, p.Image)
	}

	payload := slack.Payload{
		Text:        text,
		Username:    p.Site,
		Channel:     "#spam",
		IconUrl:     "https://shiro.ch/mymoney.png",
		UnfurlLinks: (p.Image == ""), // only unfurl digitec links
		UnfurlMedia: true,
	}
	err := slack.Send(*slacksHook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
