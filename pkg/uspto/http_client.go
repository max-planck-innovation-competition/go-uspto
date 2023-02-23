package uspto

import (
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"time"
)

// NewHttpClient builds a http client
// if the env variables PROXY or HTTP_PROXY are set
// the http client uses these
func NewHttpClient() http.Client {
	c := http.Client{
		Timeout: 1 * time.Minute,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   20 * time.Second,
			ResponseHeaderTimeout: 20 * time.Second,
			ExpectContinueTimeout: 3 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}

	proxyEvnUrl := ""
	if len(os.Getenv("PROXY")) > 0 {
		log.WithField("url", os.Getenv("PROXY")).Info("use proxy")
		proxyEvnUrl = os.Getenv("PROXY")
	}
	if len(os.Getenv("HTTP_PROXY")) > 0 {
		log.WithField("url", os.Getenv("HTTP_PROXY")).Info("use proxy")
		proxyEvnUrl = os.Getenv("HTTP_PROXY")
	}
	if len(proxyEvnUrl) > 0 {
		proxyUrl, err := url.Parse(proxyEvnUrl)
		if err != nil {
			log.Fatal(err)
		}
		c = http.Client{
			Timeout: 1 * time.Minute,
			Transport: &http.Transport{
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				Proxy:                 http.ProxyURL(proxyUrl),
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
		}
		log.Info("proxy http client")
	} else {
		log.Info("default http client")
	}

	return c
}
