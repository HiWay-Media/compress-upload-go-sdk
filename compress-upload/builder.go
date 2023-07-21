package compress

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func NewCompressUploadSdk(apiKey string, clientId string, debug bool) (ICompressUploadClient, error) {
	client := &compressUploadSdk{
		url:        "https://api.compressapi.com/v1",
		restClient: resty.New(),
		clientId:   clientId,
		apiKey:     apiKey,
		debug:      debug,
	}
	//
	if debug {
		client.restClient.SetDebug(true)
		client.debug = true
		log.Println("Debug mode is enabled for the compress upload client ")
	}
	return client, nil
}
