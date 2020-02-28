package eosapi

import (
	// "bytes"
	"context"
	// "encoding/hex"
	// "encoding/json"
	// "errors"
	// "fmt"
	// "io"
	// "log"
	"net"
	"net/http"
	// "net/http/httputil"
	"strings"
	"sync"
	"time"

	"github.com/hemlokc/goeos/crypto/ecc"
)

type API struct {
	HttpClient *http.Client
	BaseURL    string
	Signer     Signer
	Debug      bool
	Compress   CompressionType
	// Header is one or more headers to be added to all outgoing calls
	Header                  http.Header
	DefaultMaxCPUUsageMS    uint8
	DefaultMaxNetUsageWords uint32 // in 8-bytes words

	lastGetInfo      *InfoResp
	lastGetInfoStamp time.Time
	lastGetInfoLock  sync.Mutex

	customGetRequiredKeys func(ctx context.Context, tx *Transaction) ([]ecc.PublicKey, error)
}

func New(baseURL string) *API {
	api := &API{
		HttpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true, // default behavior, because of `nodeos`'s lack of support for Keep alives.
			},
		},
		BaseURL:  strings.TrimRight(baseURL, "/"),
		Compress: CompressionZlib,
		Header:   make(http.Header),
	}

	return api
}
