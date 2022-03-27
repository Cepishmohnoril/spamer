package lib

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Specify Tor proxy ip and port
var torProxy string = "socks5://127.0.0.1:9050" // 9150 w/ Tor Browser

func GetTorClient() *http.Client {
	// Parse Tor proxy URL string to a URL type
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}

	return client
}

func SendTorPost(webUrl string, client *http.Client) {
	// Make request

	data := url.Values{}
	r, err := http.NewRequest("POST", webUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)

	if err != nil {
		log.Fatal("Error making POST request.", err)
	}
	defer resp.Body.Close()

	log.Println(webUrl, " Return status code:", resp.StatusCode)
}
