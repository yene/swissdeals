package main

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GalaxusScrape(o chan<- offer) {
	doc, err := goquery.NewDocument("https://www.galaxus.ch/LiveShopping")
	if err != nil {
		log.Fatal(err)
		return
	}

	doc.Find(".daily-offer").Each(func(i int, s *goquery.Selection) {
		price := s.Find(".product-content .product-price").Text()
		price = strings.Replace(price, ".–", ".– ", -1)
		price = cutoffafterprice(price)
		price = removeNewline(price)
		title := s.Find(".product-content .product-name").Text()
		title = removeNewline(title)

		stock := s.Find(".daily-offer__stock .top").Text()
		if strings.TrimSpace(stock) == "0" {
			return
		}

		day := s.Find(".daily-offer__date .top").Text()
		day = removeNewline(day)
		link, _ := s.Find("a.overlay").Attr("href")
		link = "https://www.galaxus.ch" + link

		img := s.Find(".daily-offer__image noscript").Text()
		img = "https:" + extractSrc(img)
		img = removeShitFromURL(img) + "?fit=inside%7C258:318&output-format=progressive-jpeg"

		t := time.Now()
		if day == t.Format("02") {
			o <- offer{
				Name:  title,
				Price: price,
				Link:  link,
				Image: img,
				Site:  "galaxus.ch",
			}
		}
	})

}
