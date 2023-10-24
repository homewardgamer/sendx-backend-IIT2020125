package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Page struct {
	Content     string
	LastFetched time.Time
}

var (
	mu       sync.RWMutex
	pageData = make(map[string]Page)
)

func urlToFileName(url string) (string, error) {
	// Remove "http://" and "https://" from the URL
	fileName := strings.Replace(url, "http://", "", -1)
	fileName = strings.Replace(fileName, "https://", "", -1)
	// Remove any remaining slashes and append ".html"
	fileName = strings.Replace(fileName, "/", "", -1) + ".html"

	// Ensure the "files" directory exists
	dir := "./files"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}

	// Create the full file path
	filePath := filepath.Join(dir, fileName)

	return filePath, nil
}

func writeToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func SavePage(url, content string) error {
	mu.Lock()
	defer mu.Unlock()

	fileName, _ := urlToFileName(url)

	err := writeToFile(fileName, content)
	if err != nil {
		return err
	}

	pageData[url] = Page{
		Content:     fileName,
		LastFetched: time.Now(),
	}

	fmt.Printf("Saved page content to file %s for URL %s\n", fileName, url)

	return nil
}

func GetPage(url string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	page, exists := pageData[url]

	if !exists {
		return "", errors.New("Page not found")
	}

	if time.Since(page.LastFetched) > 60*time.Minute {
		return "", errors.New("Page too old")
	}

	return page.Content, nil
}
