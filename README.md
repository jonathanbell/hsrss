# Scrape HTML Content Into an RSS Feed

If you still like RSS as a technology (like I do), then you might be interested in creating an RSS feed for a website that no longer supports RSS. That way, you can point your RSS reader to a serverless function online and have your RSS reader update when new content is created. Yes folks, we are [screen scrapping](https://en.wikipedia.org/wiki/Data_scraping) with Golang 🙃!

One of my favourite photographs is [Hedi Slimane](https://www.hedislimane.com/). I really like his style and expulsive use of black and white. However [his blog site](https://www.hedislimane.com/diary/) odens't support RSS. Since I really want to know when new content is posted, I created a little serverless function that will check his site and return an XML RSS feed.
