package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run webcrawler.go <starting_url>")
		os.Exit(1)
	}
	startingURL := os.Args[1]
	baseDomain = getBaseDomain(startingURL)
	crawledPages.visited = make(map[string]bool)

	root := &TreeNode{Name: startingURL} // Create the root node
	wg := &sync.WaitGroup{}
	doneChan := make(chan string)

	startTime := time.Now() // Start the timer

	wg.Add(1) // Increment the wait group count

	go RecursiveCrawl(startingURL, 0, 0, doneChan, root, wg)

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
	jsonData, err := convertTreeToJson(root)
	if err != nil {
		fmt.Printf("Error converting tree to JSON: %v\n", err)
		return
	}

	// Print the JSON data after all crawling is complete
	fmt.Printf("Tree structure in JSON:\n%s\n", jsonData)

	elapsed := time.Since(startTime)
	fmt.Printf("Total time taken: %s\n", elapsed)
}

func RecursiveCrawl(parentURL string, level int, parent int, doneChan chan string, parentNode *TreeNode, wg *sync.WaitGroup) {
	defer func() {
		doneChan <- parentURL
	}()

	//if level < 2 {
	foundLinks := Crawl(parentURL)
	fmt.Printf("\nLevel:%v, Parent: %v, %s, \n", level, parent, parentURL)
	counter := 0
	for sonURL := range foundLinks {
		counter++
		//fmt.Printf("Son %v: %s\n", counter, sonURL)
		wg.Add(1) // Increment the wait group count
		sonNode := &TreeNode{Name: sonURL}
		parentNode.Children = append(parentNode.Children, sonNode) // Add child node to parent
		go RecursiveCrawl(sonURL, level+1, counter, doneChan, sonNode, wg)
	}
	//}
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
		return foundLinks
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error fetching %s: Status Code %d\n", parentURL, response.StatusCode)
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
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())

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

func getBaseDomain(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		fmt.Println("Invalid URL")
		os.Exit(1)
	}
	return parts[0] + "//" + parts[2]
}

func convertTreeToJson(root *TreeNode) ([]byte, error) {
	jsonData, err := json.Marshal(root)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// Function to recursively print the tree
func PrintTree(node *TreeNode, level int) {
	fmt.Printf("%s%s\n", getIndent(level), node.Name)
	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}

// Helper function to get the appropriate indentation
func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "-"
	}
	return indent
}
