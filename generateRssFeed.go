package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}

func getAllHsPhotos() {
	pageNumber := 1
	isConnectionSuccessful := true

	rssFeed := `<?xml version="1.0" encoding="UTF-8"?>
		<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:dc="http://purl.org/dc/elements/1.1/">
			<channel>
				<title>Hedi Silmane's Diary</title>
				<atom:link href="http://EXAMPLESITE.COM/RSS" rel="self" type="application/rss+xml" />
				<link>https://www.hedislimane.com/diary/</link>
				<description>hedislimane.com diary</description>
	`

	for isConnectionSuccessful {
		res, err := http.Get("https://www.hedislimane.com/diary/" + fmt.Sprint(pageNumber))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		// We loop unitl we no longer get a 200 status code
		if res.StatusCode != 200 {
			isConnectionSuccessful = false
			pageNumber--
			continue
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find HS posts
		doc.Find(".post").Each(func(i int, post *goquery.Selection) {
			// For each item found, get the title
			photoSrc := post.Find("img").First().AttrOr("src", "no photo")
			publishDate := post.Find(".date").First().Text()
			publishDate = strings.ReplaceAll(publishDate, "\n", "")
			publishDate = strings.Trim(strings.Split(publishDate, "/")[0], " ")

			// Dates in Go are fuuuccctt.... 😰
			inputDate := publishDate
			layout := "2006-01-02"
			t, err := time.Parse(layout, inputDate)
			if err != nil {
				panic(err)
			}
			outputFormat := "Mon, 02 Jan 2006 15:04:05 MST"
			rfc822DateString := t.In(time.FixedZone("EST", -5*60*60)).Format(outputFormat)

			rssFeed += "<item>"
			rssFeed += "<pubDate>" + rfc822DateString + "</pubDate>"
			rssFeed += "<dc:creator><![CDATA[ Hedi Silmane ]]></dc:creator>"
			rssFeed += "<description><![CDATA[<img src=" + photoSrc + " />]]></description>"
			rssFeed += "<guid>" + photoSrc + "</guid>"
			rssFeed += "</item>"

			fmt.Printf("Page: %d Post: %d\n", pageNumber, i)
		})

		pageNumber++
	}

	rssFeed += "</channel></rss>"

	if pageNumber < 1 {
		log.Fatal("Error while connecting to Hedi Slimane's Diary")
	}

	// Write our RSS feed to a file
	file, err := os.Create("./public/output.xml")
	checkErrors(err)

	_, e := file.WriteString(rssFeed)
	if e != nil {
		log.Fatal(e)
	}

	fileSyncError := file.Sync()
	if fileSyncError != nil {
		log.Fatal(fileSyncError)
	}

	// Close the file
	defer file.Close()

	fmt.Println("Got a total of ", pageNumber, " pages.")
}

func generateRssFeed() {
	getAllHsPhotos()
}
