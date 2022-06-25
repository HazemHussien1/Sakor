package crawl

import (
	"fmt"

	"github.com/gocolly/colly"
)

//use colly to crawl the url and print each one
func crawl(url string, depth int, threads int) {
	fmt.Printf("%s\n", url)

	c := colly.NewCollector(
		// set MaxDepth to the specified depth, and specify Async for threading
		colly.MaxDepth(depth),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: threads})

	// Print every href found, and visit it
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		absoluteURL := e.Request.AbsoluteURL(link)

		if absoluteURL != "" {
			// fmt.Fprintf(w, "[href] %s\n", e.Request.AbsoluteURL(link))
			fmt.Printf("%s\n", absoluteURL)
			// Visit link found on page on a new thread
			e.Request.Visit(link)
		}
	})
	c.Visit(url)

	// Wait until threads are finished
	c.Wait()
}
