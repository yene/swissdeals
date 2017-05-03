package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func DaydealScrape(o chan<- offer) {
	doc, err := goquery.NewDocument("https://www.daydeal.ch")
	if err != nil {
		log.Fatal(err)
		return
	}

	s := doc.Find(".ym-wbox")
	price := s.Find(".salePrice .price").Text()
	price = removeNewline(price)

	originalPrice := s.Find(".originalPrice").Text()
	originalPrice = strings.Replace(originalPrice, "*", "", -1)
	originalPrice = removeNewline(originalPrice)

	price = "CHF " + price + " statt vorher " + originalPrice

	title := s.Find(".meta-first-line").Text()
	title = title + " " + s.Find(".meta-second-line").Text()
	title = removeNewline(title)
	img, _ := s.Find(".productImage").Attr("src")

	o <- offer{
		Name:  title,
		Price: price,
		Link:  "https://www.daydeal.ch",
		Image: img,
		Site:  "daydeal.ch",
	}
}
