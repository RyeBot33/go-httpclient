package examples

import (
	"github.com/bsi/go-httpclient/gohttp"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).Build()

	return client
}
