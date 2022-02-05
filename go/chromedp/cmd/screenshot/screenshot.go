package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	start := time.Now()
	// navigate to a page, wait for an element, click
	var res string
	var imageBuf []byte
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(`https://github.com`),
		chromedp.Text(`h1`, &res, chromedp.NodeVisible, chromedp.ByQuery),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.ScrollIntoView(`h1`),
		chromedp.WaitVisible(`h1`, chromedp.ByQuery),
		chromedp.Screenshot(`h1`, &imageBuf, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("h1-screenshot.png", imageBuf, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())
}
