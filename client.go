// Package eveapi implements access to EVE Onlines XML APi
package eveapi

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	Tranquility      = "https://api.eveonline.com"
	Singularity      = "https://api.testeveonline.com"
	dateFormat       = "2006-01-02 15:04:05"
	defaultUserAgent = "Go API Wrapper"
)

type Key struct {
	ID    int64
	VCode string
}

type API struct {
	Server    string
	APIKey    Key
	UserAgent string
	Debug     bool
}

func Simple(key Key) *API {
	return &API{Tranquility, key, defaultUserAgent, false}
}

type APIResult struct {
	Version     int64     `xml:"version,attr"`
	CurrentTime eveTime   `xml:"currentTime"`
	Error       *APIError `xml:"error,omitempty"`
	CachedUntil eveTime   `xml:"cachedUntil"`
}
type APIError struct {
	Code    int64  `xml:"code,attr"`
	Message string `xml:",chardata"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Error! %v (code:%v)", e.Message, e.Code)
}
func (api API) Call(path string, args url.Values, output interface{}) error {
	uri := api.Server + path
	if args == nil {
		args = url.Values{}
	}
	args.Set("keyID", strconv.FormatInt(api.APIKey.ID, 10))
	args.Set("vCode", api.APIKey.VCode)
	resp, err := http.PostForm(uri, args)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if api.Debug {
		io.Copy(os.Stdout, resp.Body)
	}
	//TODO: LimitReader if it explodes?
	err = xml.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		return err
	}
	return nil
}
