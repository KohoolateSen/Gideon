package httpproxy

import (
	. "Gideon/internal/config"
	"net/http"
	"net/url"
	"time"
)

var GlobalHTTPClient *http.Client

func SetupHTTPClient() error {
	use_proxy := AppConfig.Proxy.Enabled
	raw_proxy_url := AppConfig.Proxy.Url

	if use_proxy {
		proxy_url, err := url.Parse(raw_proxy_url)

		if err != nil {
			return err
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy_url),
		}

		GlobalHTTPClient = &http.Client{
			Transport: transport,
			Timeout:   60 * time.Second,
		}

		return nil
	}

	GlobalHTTPClient = &http.Client{}
	return nil
}
