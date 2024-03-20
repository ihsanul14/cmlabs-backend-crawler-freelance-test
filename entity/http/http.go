package http

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

type Http struct{}

type Result struct {
	Link []string
	Body string
}

type IHttp interface {
	Call(context.Context, string) (*Result, error)
}

func NewHttp() IHttp {
	return &Http{}
}

func (m *Http) Call(ctx context.Context, u string) (*Result, error) {
	chromeCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(chromeCtx, opts...)
	defer cancel()

	chromeCtx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	if err := chromedp.Run(chromeCtx,
		chromedp.Navigate(u)); err != nil {
		return nil, fmt.Errorf("entity.http.Call: %v", err)
	}

	var links []string
	var body string
	if err := chromedp.Run(chromeCtx,
		chromedp.Evaluate(`Array.from(document.getElementsByTagName('a')).map(a => a.href)`, &links),
		chromedp.OuterHTML("html", &body),
	); err != nil {
		return nil, fmt.Errorf("entity.http.Call: %v", err)
	}

	return &Result{
		Link: links,
		Body: body,
	}, nil
}
