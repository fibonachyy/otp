package kavenegar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type httpClient struct {
	BaseUrl string
	ApiKey  string
	Sender  string
	cli     *http.Client
}

// https://api.kavenegar.com/v1/{API-KEY}/Scope/MethodName.OutputFormat

func newHttpClient(url, apiKey, sender string, timeout time.Duration) *httpClient {
	return &httpClient{
		BaseUrl: url,
		cli:     &http.Client{Timeout: timeout},
		Sender:  sender,
		ApiKey:  apiKey,
	}
}
func (c *httpClient) SendOtp(receptors []string, otp string) (*SendOtpResponse, error) {
	receptorSTR := strings.Join(receptors, ",")
	u := fmt.Sprintf("%s/%s/sms/send.json?receptor=%s&sender=%s&message=%s", c.BaseUrl, c.ApiKey, receptorSTR, c.Sender, otp)
	fmt.Printf(u)
	v := &SendOtpResponse{}

	if err := c.do(http.MethodGet, u, c.authHeader(), nil, v); err != nil {
		return nil, err
	}
	return v, nil
}

func (c *httpClient) do(method, url string, header http.Header, body io.Reader, v interface{}) error {
	req, err := http.NewRequest(method, url, body)
	req.Header = header
	if err != nil {
		return err
	}
	res, err := c.cli.Do(req)
	if err != nil {
		return fmt.Errorf("request error:%v", err)
	}
	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("decode error:%v", err)
	}
	return nil
}

func (c *httpClient) authHeader() http.Header {
	header := http.Header{}
	// header.Set("Authorization", "Currency "+c.ApiKey)
	header.Add("content-type", "application/json")
	return header
}
