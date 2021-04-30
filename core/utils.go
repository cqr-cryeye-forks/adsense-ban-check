package core

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Result struct {
	Target string `json:"target"`
	Data string `json:"data"`
}

func setHeaders(r *http.Request) {
	r.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:88.0) Gecko/20100101 Firefox/88.0")
	r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	r.Header.Set("Accept-Language", "en-US,en;q=0.5")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Origin", "https://adsensechecker.com")
	r.Header.Set("DNT", "1")
	r.Header.Set("Connection", "keep-alive")
	r.Header.Set("Referer", "https://adsensechecker.com/")
}

func getResponse(target *string) *http.Response {
	client := &http.Client{}
	var d = `url=` + *target + `&submit=CHECK`
	var data = strings.NewReader(d)
	req, err := http.NewRequest("POST", "https://adsensechecker.com/", data)

	if err != nil {
		log.Fatal(err)
	}

	setHeaders(req)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func parseResponse(r *http.Response, result *string) bool {
	bodyText, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	str := strings.ToLower(string(bodyText))

	if strings.Contains(str, "not a valid url") {
		*result = "Not a valid url."
		panic(*result)
	}

	return !strings.Contains(str, "is not banned by google adsense.")
}

func IsBanned(target, result *string) bool {
	resp := getResponse(target)

	if resp.StatusCode != 200 {
		log.Fatalf("Unexpected service response. Status code: %d", resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	return parseResponse(resp, result)
}

func WriteResult(target, result *string) {
	file, err := json.MarshalIndent(Result{*target, *result}, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("output.json", file, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
