package centrifugoKit

import (
	"github.com/centrifugal/gocent/v3"
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/httpClientKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"net/http"
)

// NewClient
/*
@param httpClient 可以为nil（将使用默认值 httpClientKit.DefaultHttpClient）
*/
func NewClient(addrs []string, apiKey string, httpClient *http.Client) (*gocent.Client, error) {
	// addrs
	addrs = sliceKit.Uniq(addrs)
	addrs = sliceKit.RemoveEmpty(addrs)
	if err := sliceKit.AssertNotEmpty(addrs, "addrs"); err != nil {
		return nil, err
	}
	// apiKey
	if err := strKit.AssertNotEmpty(apiKey, "apiKey"); err != nil {
		return nil, err
	}
	// httpClient
	if httpClient == nil {
		httpClient = httpClientKit.DefaultHttpClient
	}

	/* (1) 一个地址 */
	if len(addrs) == 1 {
		config := gocent.Config{
			Addr:       addrs[0],
			Key:        apiKey,
			HTTPClient: httpClient,
		}
		return gocent.New(config), nil
	}
	/* (2) 多个地址 */
	getAddr := func() (string, error) {
		// TODO:待完善
		return "", nil
	}
	config := gocent.Config{
		GetAddr:    getAddr,
		Key:        apiKey,
		HTTPClient: httpClient,
	}
	return gocent.New(config), nil
}