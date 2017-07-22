package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Scrape digitec or galaxus
func DigitecGalaxusScrape(o chan<- offer, site string) {
	doc, err := goquery.NewDocument("https://www." + site + ".ch/LiveShopping")
	if err != nil {
		log.Fatal(err)
		return
	}

	doc.Find("article.daily-offer-new").Each(func(i int, s *goquery.Selection) {
		// TODO: only post one deal for debugging
		if i > 0 {
			return
		}

		count := s.Find(".product-count__text").Text()
		if strings.Contains(count, "beendet") {
			return
		}

		price := s.Find(".product-price").Text()
		// TODO: get product price without sub nodes, instead of doing this string manipulation.
		price = strings.Replace(price, ".–1", ".-", -1)
		price = strings.Replace(price, ".–", ".– ", -1)
		price = removeNewline(price)
		if price == "" {
			return
		}

		title := s.Find(".product-name").Text()
		title = removeNewline(title)

		link, _ := s.Find("a.overlay").Attr("href")
		link = "https://www." + site + ".ch" + link

		/*
			TODO: currently the image is inside the javascript lazy load, it does not work without javascript.
				img := s.Find(".daily-offer__image noscript").Text()
				img = "https:" + extractSrc(img)
				img = removeShitFromURL(img) + "?fit=inside%7C258:318&output-format=progressive-jpeg"
		*/

		o <- offer{
			Name:  title,
			Price: price,
			Link:  link,
			//Image: img,
			Site: site + ".ch",
		}

	})

}
