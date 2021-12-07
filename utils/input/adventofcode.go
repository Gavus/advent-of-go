// Functionality directly related to advent of code webpage.
package input

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	domain         = "adventofcode.com"
	inputUrlFormat = "https://%s/%d/day/%d/input"
	sessionIdPath  = "./sessionid"
)

// Get input directly from Advent of code webpage.
func DownloadInput(year int, day int) ([]byte, error) {
	input := make([]byte, 0)
	url := fmt.Sprintf(inputUrlFormat, domain, year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return input, err
	}

	cookie, err := getCookie()
	if err != nil {
		return input, err
	}
	req.AddCookie(cookie)

	// Send the http request and receive the response.
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return input, err
	} else if resp.StatusCode != 200 {
		return input, errors.New(resp.Status)
	}

	// Save the response.
	defer resp.Body.Close()
	input, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return input, err
	}

	return input, nil
}

// Get session cookie from adventofcode.
func getCookie() (*http.Cookie, error) {
	if fileExists(sessionIdPath) {
		input, err := getInputFromFile(sessionIdPath, "\n")
		if err != nil {
			return nil, errors.New("Could not read sessionid file")
		} else if len(input) != 1 {
			return nil, errors.New("sessionid should only contain one line followed by a newline")
		} else {
			cookie := new(http.Cookie)
			cookie.Name, cookie.Value = "session", input[0]
			return cookie, nil
		}
	}

	return nil, errors.New("Could not find cookie")
}
