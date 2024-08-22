package core

import "net/url"

func RemoveScheme(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		panic(err)
	}

	// Clear the scheme (e.g., http, https)
	parsedURL.Scheme = ""

	// Reconstruct the URL without the scheme
	newURL := parsedURL.String()

	// Remove the leading "//" if the scheme is empty
	if len(newURL) > 2 && newURL[:2] == "//" {
		newURL = newURL[2:]
	}

	return newURL
}
