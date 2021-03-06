package swapi

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

type Client struct {
	base string
	http *http.Client
}

// NewClient ...
func NewClient(c *http.Client, base string) *Client {
	if c == nil {
		c = http.DefaultClient
	}

	return &Client{base: base, http: c}
}

func (c *Client) NewRequest(ctx context.Context, url string, method string, body io.Reader) (*http.Request, error) {
	if len(url) == 0 {
		return nil, errors.New("invalid empty-string url")
	}

	if url[0] == '/' { // Assume the user has given a relative path.
		url = c.base + url
	}
	var (
		r   *http.Request
		err error
	)
	if method == "GET" {
		r, err = http.NewRequest("GET", url, nil)
	} else {
		r, err = http.NewRequest("POST", url, body)
	}

	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("X-Vaccount-Id", "1")
	r.Header.Add("X-Portal-Id", "1")
	r.Header.Add("vaccount-id", "1")
	r.Header.Add("portal", "1")
	r.Header.Add("SkipCsrfCheck", b64.StdEncoding.EncodeToString([]byte("VoonikFramework")))
	r.Header.Add("VServiceCheck", b64.StdEncoding.EncodeToString([]byte("VNKSRVC")))

	var result httpstat.Result
	ctx = httpstat.WithHTTPStat(r.Context(), &result)
	ctx = context.WithValue(ctx, "requestTrace", &result)

	return r.WithContext(ctx), nil
}

// Do the request.
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	fmt.Printf("req %+v", r)
	resp, err := c.http.Do(r)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	// respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody))
	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, fmt.Errorf("unable to parse JSON [%s %s]: %v", r.Method, r.URL.RequestURI(), err)
		}
	}
	// Show the results
	ctx := r.Context()
	result := ctx.Value("requestTrace").(*httpstat.Result)
	log.Printf("DNS lookup: %d ms", int(result.DNSLookup/time.Millisecond))
	log.Printf("TCP connection: %d ms", int(result.TCPConnection/time.Millisecond))
	log.Printf("TLS handshake: %d ms", int(result.TLSHandshake/time.Millisecond))
	log.Printf("Server processing: %d ms", int(result.ServerProcessing/time.Millisecond))
	log.Printf("Content transfer: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))

	return resp, nil
}
