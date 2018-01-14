package bitstamp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type client struct {
	httpClient *http.Client
}

func NewClient() (c *client) {
	return &client{&http.Client{}}
}

func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	select {
		case r := <-done:
			return r.resp, r.err
		case <-timer.C:
			return nil, errors.New("Timeout on reading from Bistamp API")
	}
}

func (c *client) do(method string, resource string) (response []byte, err error) {
	connectTimer := time.NewTimer(DEFAULT_HTTPCLIENT_TIMEOUT * time.Second)

	rawurl := fmt.Sprintf("%s%s/%s", API_BASE, API_VERSION, resource)

	req, err := http.NewRequest(method, rawurl, strings.NewReader(""))
	if err != nil {
		return
	}
	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
