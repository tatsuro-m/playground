package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	start := time.Now()
	var res string
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate("https://github.com"),
		chromedp.ScrollIntoView("footer"),
		//chromedp.WaitVisible("footer < div"),
		chromedp.Text("h1", &res, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
}
