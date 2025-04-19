package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL parsing:")

	// A sample URL string
	myurl := "https://www.youtube.com/"
	fmt.Printf("Type of url: %T\n", myurl) // string

	// Parsing the URL
	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Can't parse URL:", err)
		return
	}
	fmt.Printf("Type of parsedURL: %T\n", parsedURL) // *url.URL (a struct pointer)

	// Accessing parts of the URL
	fmt.Println("Scheme:   ", parsedURL.Scheme)   // https
	fmt.Println("Host:     ", parsedURL.Host)     // www.youtube.com
	fmt.Println("Path:     ", parsedURL.Path)     // / (empty path here)
	fmt.Println("RawQuery: ", parsedURL.RawQuery) // empty here

	// ===== Modify URL components =====
	parsedURL.Path = "results"
	queryParams := url.Values{}
	queryParams.Add("search_query", "golang tutorials")
	queryParams.Add("sort", "relevance")
	parsedURL.RawQuery = queryParams.Encode()

	// Construct the URL string back from the modified URL object
	newUrl := parsedURL.String()
	fmt.Println("Modified URL:", newUrl)
}

/*
======= Additional Notes =======
- `url.Parse()` returns a pointer to a URL struct.
- You can modify individual components like `Path`, `Host`, or `RawQuery`.
- Use `url.Values{}` to build query parameters easily.
- `queryParams.Encode()` converts the map into a URL-encoded string.
- Query parameters are automatically encoded (e.g., space becomes + or %20).

Example output:
Modified URL: https://www.youtube.com/results?search_query=golang+tutorials&sort=relevance

======= Common Methods =======
- `.Scheme`         → protocol (http/https)
- `.Host`           → domain name
- `.Path`           → path on the server
- `.RawQuery`       → query string (can be set manually or via `url.Values`)
- `url.Values{}`    → map-like structure to manage query params
- `.String()`       → reconstructs full URL
*/
