package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type TreeNode struct {
	Name     string      `json:"name"`
	Children []*TreeNode `json:"children,omitempty"`
}

type Pages struct {
	visitedMutex sync.Mutex
	visited      map[string]bool
}

var baseDomain = ""
var crawledPages Pages

func main() {

	fmt.Println(len(os.Args))

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run webcrawler.go <starting_url> <deep_level_to_retrieve>")
		os.Exit(1)
	}
	startingURL := os.Args[1]
	baseDomain = GetBaseDomain(startingURL)
	maxDeepLevel, _ := strconv.ParseInt(os.Args[2], 0, 0)

	crawledPages.visited = make(map[string]bool)

	root := &TreeNode{Name: startingURL} // Create the root node
	wg := &sync.WaitGroup{}
	doneChan := make(chan string)

	startTime := time.Now() // Start the timer

	wg.Add(1) // Increment the wait group count

	go RecursiveCrawl(startingURL, 0, int(maxDeepLevel), 0, doneChan, root, wg)

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(doneChan)
	}()

	// Read from doneChan
	for range doneChan {
		// Do nothing here, just let it finish
	}

	// Print the tree
	PrintTree(root, 0)

	// Convert the tree to JSON
	jsonData, err := ConvertTreeToJson(root)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error converting tree to JSON.", err)
		LogError(errMsg)
		return
	}

	//Get the domain in the format mydomain.com
	domain := GetShortDomain(baseDomain)

	err = SaveJSONToFileWithTimestamp(jsonData, domain)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error saving JSON to file.", err)
		LogError(errMsg)
		return
	}

	// Print the JSON data after all crawling is complete
	fmt.Printf("Tree structure in JSON:\n%s\n", jsonData)

	elapsed := time.Since(startTime)

	fmt.Printf("Total time taken: %s\n", elapsed)

}

func RecursiveCrawl(parentURL string, currentLevel int, maxLevel int, parent int, doneChan chan string, parentNode *TreeNode, wg *sync.WaitGroup) {
	defer func() {
		doneChan <- parentURL
	}()

	if currentLevel <= maxLevel {
		foundLinks := Crawl(parentURL)
		fmt.Printf("\nLevel:%v, Parent: %v, %s", currentLevel, parent, parentURL)
		counter := 0
		for sonURL := range foundLinks {
			counter++
			//fmt.Printf("Son %v: %s\n", counter, sonURL)
			wg.Add(1) // Increment the wait group count
			sonNode := &TreeNode{Name: sonURL}
			parentNode.Children = append(parentNode.Children, sonNode) // Add child node to parent
			go RecursiveCrawl(sonURL, currentLevel+1, maxLevel, counter, doneChan, sonNode, wg)
		}
	}
	wg.Done() // Decrement the parent wait group count
}

func Crawl(parentURL string) map[string]bool {
	crawledPages.visitedMutex.Lock()
	defer crawledPages.visitedMutex.Unlock()

	foundLinks := make(map[string]bool, 0)

	if crawledPages.visited[parentURL] {
		return foundLinks
	}

	crawledPages.visited[parentURL] = true

	response, err := http.Get(parentURL)

	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", parentURL, err)
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error fetching "+parentURL, err)
		LogError(errMsg)

		return foundLinks
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error fetching %s: Status Code %d\n", parentURL, response.StatusCode)
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error fetching "+parentURL+" Status Code "+fmt.Sprint(response.StatusCode), err)
		LogError(errMsg)

		return foundLinks
	}

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return foundLinks
			}

			errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error tokenizing HTML ", tokenizer.Err())
			LogError(errMsg)

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						foundLink := attr.Val
						if string(foundLink[len(foundLink)-1]) == "/" {
							foundLink = foundLink[:len(foundLink)-1]
						}
						if strings.HasPrefix(foundLink, "/") {
							foundLink = parentURL + foundLink
						}
						if strings.HasPrefix(foundLink, baseDomain) && foundLink != parentURL {
							foundLinks[foundLink] = true
						}
					}
				}
			}
		}
	}
	//return foundLinks
}

func GetBaseDomain(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {

		fmt.Println("Invalid URL")

		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Invalid URL ", url)
		LogError(errMsg)

		os.Exit(1)
	}
	return parts[0] + "//" + parts[2]
}

func ConvertTreeToJson(root *TreeNode) ([]byte, error) {
	jsonData, err := json.Marshal(root)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// Function to recursively print the tree
func PrintTree(node *TreeNode, level int) {
	fmt.Printf("%s%s\n", GetIndent(level), node.Name)
	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}

// Helper function to get the appropriate indentation
func GetIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "-"
	}
	return indent
}

// Saves JSON data to a file with a timestamp in the filename.
func SaveJSONToFileWithTimestamp(jsonData []byte, prefix string) error {
	// Create the "sitemaps" subfolder if it doesn't exist
	sitemapsFolder := "sitemaps"
	if _, err := os.Stat(sitemapsFolder); os.IsNotExist(err) {
		os.Mkdir(sitemapsFolder, os.ModePerm)
	}

	// Set the format expected for output
	timestamp := time.Now().Format("2006-01-02_15-04-05") // Format: YYYY-MM-DD_HH-MM-SS

	// Create the filename within the "sitemaps" subfolder
	filename := fmt.Sprintf("%s/%s_%s.json", sitemapsFolder, prefix, timestamp)

	// Open the file on disk
	file, err := os.Create(filename)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error opening the file", filename)
		LogError(errMsg)
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Error writing to the file", filename)
		LogError(errMsg)
		return err
	}

	fmt.Printf("JSON data saved to %s\n", filename)
	return nil
}

// Return an URL in the format mydomain_com
func GetShortDomain(inputURL string) string {
	// Parse the input URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {

		fmt.Println("Invalid URL:", err)

		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", baseDomain, "Invalid URL ", err)
		LogError(errMsg)

		return ""
	}

	// Remove "www." from the host if present
	host := strings.TrimPrefix(parsedURL.Hostname(), "www.")

	// Replace periods with underscores to get the base domain
	baseDomain := strings.Replace(host, ".", "_", -1)

	return baseDomain
}

func LogError(errorMessage string) {
	// Create a "logs" subfolder if it doesn't exist
	logsFolder := "logs"
	if _, err := os.Stat(logsFolder); os.IsNotExist(err) {
		os.Mkdir(logsFolder, os.ModePerm)
	}

	// Get the current date and time
	currentTime := time.Now()
	// Format the date in YYYY-MM-DD format
	date := currentTime.Format("2006-01-02")

	// Create or open the log file in the "logs" subfolder with the current date as the filename
	logFileName := fmt.Sprintf("%s/%s.log", logsFolder, date)
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// If there is an error opening the log file, print the error and exit
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// Create a new logger that writes to the log file
	logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Log the error message with the current date and time
	logger.Printf("%s - %s", currentTime.Format("2006-01-02 15:04:05"), errorMessage)
}
