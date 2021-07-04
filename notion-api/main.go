package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.notion.com/v1"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, baseURL+"/pages/"+os.Getenv("ROOT_PAGE_ID"), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_API_KEY"))
	req.Header.Add("Notion-Version", "2021-05-13")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
