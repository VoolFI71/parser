package main

import (
	"fmt"
	"io/ioutil"
 	"net/http"
    "encoding/json"

)

func main() {
	url := "https://kuper.ru/api/v3/stores/48037/departments?per_page=100&page=1"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}	
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 YaBrowser/25.2.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "ru,en;q=0.9,be;q=0.8,zh;q=0.7")
	req.Header.Set("Baggage", "sentry-environment=server,sentry-release=r25-02-27-1759-34126301,sentry-public_key=f9d0a0afb8d5420bb353a190580ae049,sentry-trace_id=432489b46f02c9afda73a7ec05cc6b61")
	req.Header.Set("Client-Token", "7ba97b6f4049436dab90c789f946ee2f")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close() 

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var prettyJSON map[string]interface{}
	if err := json.Unmarshal(body, &prettyJSON); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	formattedJSON, err := json.MarshalIndent(prettyJSON, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при форматировании JSON:", err)
		return
	}

	err = ioutil.WriteFile("response.txt", formattedJSON, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Статус ответа:", resp.Status)
	fmt.Println("json сохранен в файл response.txt")
	
	// Пример получения продукта и цены
	if departments, ok := prettyJSON["departments"].([]interface{}); ok {
		for _, dep := range departments {
			if department, ok := dep.(map[string]interface{}); ok {
				fmt.Println("Department Name:", department["name"])
				if products, ok := department["products"].([]interface{}); ok {
					for _, prod := range products {
						if product, ok := prod.(map[string]interface{}); ok {
							fmt.Println("Product name and price:", product["name"], product["price"], product["canonical_url"])
						}
					}
				}
			}
		}
	}
}
