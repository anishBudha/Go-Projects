package main

import (
	"fmt"
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

var urlMap = make(map[string]string)
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const domain = "short.url/"

func generateShortURL () string {
	// this will generate a random combination of letters for the prefix of our short url
	b := make([]byte, 4) // a slice of size 4
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func shortenURL(originalURL string) string {
	shortCode := generateShortURL()
	for urlMap[shortCode] != "" {
		// checks if the shortCode generated already exists with a original url, means urlMap[shortcode generated] must be empty
		shortCode = generateShortURL()
	}
	urlMap[shortCode] = originalURL
	return domain + shortCode
}

func resolveURL(short string) (string, bool) {
	if strings.HasPrefix(short, domain) {
		short = strings.TrimPrefix(short, domain)
	}
	original, exists := urlMap[short] // stores original url in urlMap[inserted short url code], if exists true otherwise false, map return value and a bool ok
	return original, exists
}

func main () {
	rand.Seed(time.Now().UnixNano()) // random generate every time 
	scanner := bufio.NewScanner(os.Stdin) // creates a scanner that reads input from the keyboard
  
	for {
		fmt.Println("\n 1.Shorten URL \n 2.Resolve URL \n 3.Quit \n 4.Print map")
		fmt.Print(" Enter a choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		
		switch choice {
		case "1":
			fmt.Print(" Enter original URL: ")
			scanner.Scan()
			originalURL := strings.TrimSpace(scanner.Text())
			short := shortenURL(originalURL)
			fmt.Println(" Short URL:", short)
		case "2":
			fmt.Print(" Enter the short URL to resolve: ")
			scanner.Scan()
			shortURL := strings.TrimSpace(scanner.Text())
			if original, ok := resolveURL(shortURL); ok {
				fmt.Println(" Original URL:", original)
			} else {
				fmt.Println(" Short URL not found")
			}
		
		case "3":
			fmt.Println(" Existing...")
			return

		case "4":
			fmt.Println("\n", urlMap)

		default:
			fmt.Println(" Invalid Choice")
		}
		
	}

}























