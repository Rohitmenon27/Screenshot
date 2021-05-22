package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	abc, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	url := "https://golang.org"
	filename := "golang.png"

	var imageBuf []byte
	if err := chromedp.Run(abc, screenshotTasks(url, &imageBuf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filename, imageBuf, 9544); err != nil {
		log.Fatal(err)
	}
}

func screenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(abc context.Context) (err error) {
			*imageBuf, _, err = page.PrintToPDF().WithPrintBackground(false).Do(abc)
			return err
		}),
	}
}
