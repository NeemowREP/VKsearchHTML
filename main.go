package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	keys := []string{"ТНС%20энерго%20НН", "энергосбыт", "ТНС"}

	for _, k := range keys {

		escQuery := url.QueryEscape(k)

		vkURL := "https://vk.com/search?q=" + escQuery

		parsedvkURL, err := url.Parse(vkURL)
		if err != nil {
			fmt.Println("Ошибка при парсе URL:", err)
			return
		}

		req, err := http.NewRequest("GET", parsedvkURL.String(), nil)
		if err != nil {
			fmt.Println("Ошибка при создании GET запроса:", err)
			return
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка при чтении ответа:", err)
			return
		}
		fmt.Println(string(body))
	}

}
