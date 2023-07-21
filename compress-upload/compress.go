package compress

import (
	"github.com/go-resty/resty/v2"
)

type compressUploadSdk struct {
	url        string
	apiKey     string
	clientId   string
	restClient *resty.Client
	debug      bool
}

type ICompressUploadClient interface {
	HealthCheck() error
	IsDebug() bool
}

func (o *compressUploadSdk) IsDebug() bool {
	return o.debug
}

func (o *compressUploadSdk) HealthCheck() error {
	_, err := o.restyGet(o.url, nil)
	if err != nil {
		return nil
	}
	return nil
}
