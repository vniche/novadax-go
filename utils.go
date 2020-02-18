package novadax

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (client *Client) signRequest(req *http.Request, method string, path string, encodedData string, now string) {

	// initializes hmac+sha256 hash with configured private key
	hash := hmac.New(sha256.New, []byte(client.Config.PrivateKey))

	// writes signature to hmac+sha256 hash
	hash.Write([]byte(method + "\\n" + path + "\\n" + encodedData + "\\n" + now))

	// encodes hash to hex
	signedHash := hex.EncodeToString(hash.Sum(nil))

	req.Header.Set("X-Nova-Signature", signedHash)
	req.Header.Set("X-Nova-Access-Key", client.Config.AccessKey)
	req.Header.Set("X-Nova-Timestamp", now)
}

func (client *Client) buildRequest(method, path string, body interface{}, secure bool) (*http.Request, error) {
	// parses request path
	splittedPath := strings.Split(path, "?")
	apiPath := splittedPath[0]

	// assembles request info
	rel := &url.URL{Path: apiPath}
	apiURL := client.BaseURL.ResolveReference(rel)
	var jsonBytes []byte
	var err error
	if body != nil {
		jsonBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, apiURL.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	now := strconv.FormatInt(time.Now().Unix()*1000, 10)
	if method == "GET" {
		var params url.Values
		if len(splittedPath) > 1 {
			params, err = url.ParseQuery(splittedPath[1])
			if err != nil {
				return nil, err
			}
		}

		if params.Encode() != "" {
			req.URL.RawQuery = params.Encode()
		}

		// if requesting for identity-required resources
		if secure {
			client.signRequest(req, method, apiPath, params.Encode(), now)
		}
	} else {
		if body != nil && secure {
			bodyJSON, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}

			jsonString := string(bodyJSON)

			log.Printf("%s", jsonString)

			hash := MD5Digest(&jsonString)
			client.signRequest(req, method, apiPath, hash, now)
		}
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", client.UserAgent)
	return req, nil
}

// StatusCode stands for the resp.Status code index
const StatusCode = 0

func (client *Client) do(req *http.Request, body interface{}) (*http.Response, error) {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	statusInfo := strings.Split(resp.Status, " ")

	// test for response status code
	status, err := strconv.Atoi(statusInfo[StatusCode])
	if err != nil {
		return nil, err
	} else if status < 200 || status > 299 {
		return nil, fmt.Errorf("request failed with status %d", status)
	}

	err = json.Unmarshal(bodyBytes, body)
	return resp, err
}

// structToURLValues transforms a go struct into a url.Values for API requests
func structToURLValues(any interface{}) (values url.Values) {
	values = url.Values{}
	fieldValue := reflect.ValueOf(any).Elem()
	fieldType := fieldValue.Type()
	for index := 0; index < fieldValue.NumField(); index++ {
		var (
			value string
			field = fieldValue.Field(index)
		)

		switch field.Interface().(type) {
		case int, int8, int16, int32, int64:
			if field.Int() == 0 {
				break
			}

			value = strconv.FormatInt(field.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			if field.Uint() == 0 {
				break
			}

			value = strconv.FormatUint(field.Uint(), 10)
		case float32:
			if field.Float() == 0 {
				break
			}

			value = strconv.FormatFloat(field.Float(), 'f', 4, 32)
		case float64:
			if field.Float() == 0 {
				break
			}

			value = strconv.FormatFloat(field.Float(), 'f', 4, 64)
		case []byte:
			if len(field.Bytes()) == 0 {
				break
			}

			value = string(field.Bytes())
		case string:
			if field.String() == "" {
				break
			}

			value = field.String()
		}

		if value != "" {
			values.Set(fieldType.Field(index).Tag.Get("json"), value)
		}
	}
	return values
}

// MD5Digest returns a string digested MD5 hash
func MD5Digest(s *string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(*s)))
}
