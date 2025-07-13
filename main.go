package main

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"time"
	"github.com/chromedp/chromedp"
)

func main() {

	keys := []string{"ТНС%20энерго%20НН", "энергосбыт", "ТНС"}

	regularExpressions := regexp.MustCompile(`https://vk\.com/(wall[-]?\d+_\d+|public\d+)`)

	operaPath := `C:\Users\TheBoss\AppData\Local\Programs\Opera GX\opera.exe`

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(operaPath),
		chromedp.Flag("headless", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64)"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	for _, k := range keys {

		var content string

		escQuery := url.QueryEscape(k)

		vkURL := "https://vk.com/search?q=" + escQuery

		err := chromedp.Run(ctx,
			chromedp.Navigate(vkURL),
			chromedp.Sleep(4*time.Second),
			chromedp.OuterHTML("html", &content),
		)

		if err != nil {
			fmt.Printf("Ошибка при загрузке страницы для \"%s\": %v\n", k, err)
			continue
		}

		matches := regularExpressions.FindAllString(content, -1)
		if len(matches) == 0 {
			fmt.Println("Увы")
		}

		fmt.Printf("Ссылки по запросу \"%s\":\n", k)
		for _, link := range matches {
			fmt.Println(link)
		}
	}

}
