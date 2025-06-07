package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Watchdog running..")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_urls := os.Getenv("URLS")
	urls := strings.Split(_urls, ",")
	for _, url := range urls {
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get(url)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Printf("%v - %s\n", res.StatusCode, url)
	}
}
