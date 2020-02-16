package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func (c *Client) signRequest(method, path string, body interface{}, secure bool) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// if requesting for identity-required resources
	if secure {
		now := time.Now()
		if method == "GET" {
			splittedPath := strings.Split(path, "?")
			apiPath := splittedPath[0]
			params, err := url.ParseQuery(splittedPath[1])
			if err != nil {
				log.Fatalf("%s", err.Error())
			}

			h := hmac.New(sha256.New, []byte(os.Getenv("NOVADAX_SECRET_KEY")))
			h.Write([]byte(method + "\n" + apiPath + "\n" + params.Encode() + "\n" + string(now.Unix())))
			sha := hex.EncodeToString(h.Sum(nil))

			req.Header.Set("X-Nova-Signature", sha)
		}
		req.Header.Set("X-Nova-Access-Key", os.Getenv("NOVADAX_ACCESS_KEY"))
		req.Header.Set("X-Nova-Timestamp", now.String())
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
