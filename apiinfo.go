package eveapi

import "time"

const (
	// APIInfoURL is the URL to fetch account info
	APIInfoURL = "/account/APIKeyInfo.xml.aspx"
)

// APIInfo contains information about the current API key
type APIInfo struct {
	AccessMask string   `xml:"accessMask,attr"`
	Type       string   `xml:"type,attr"`
	Expires    *eveTime `xml:"expires,attr"`
}

// APIInfoResult is the result wrapper for APIInfo
type APIInfoResult struct {
	APIResult
	Info APIInfo `xml:"result>key"`
}

// Info calls the API
// Returns a APIInfo struct
func (api API) Info() (*APIInfo, error) {
	output := APIInfoResult{}
	err := api.Call(APIInfoURL, nil, &output)
	if err != nil {
		return nil, err
	}
	if output.Error != nil {
		return nil, output.Error
	}

	if output.Info.Expires.Equal(time.Unix(0, 0)) {
		output.Info.Expires = nil
	}

	return &output.Info, nil
}
