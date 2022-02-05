package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/dom"
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
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(`https://github.com`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.ScrollIntoView(`footer`),
		chromedp.Text(`h1`, &res, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			fmt.Print(res) // print HTML source code
			return err
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\nTook: %f secs\n", time.Since(start).Seconds())
}
