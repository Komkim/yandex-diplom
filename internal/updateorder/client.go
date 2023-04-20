package updateorder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"yandex-diplom/config"
	storage "yandex-diplom/storage/repository"
)

const (
	SCHEME = "http"
	PATH   = "api/orders"
)

type MyClient struct {
	client *http.Client
	config *config.Server
}

func NewClient(config *config.Server) *MyClient {
	return &MyClient{
		client: &http.Client{},
		config: config,
	}
}

func (c *MyClient) GetAccrual(number int64) (*storage.OrderAccrual, error) {
	cfgAddr := strings.Split(c.config.AccrualAddress, "://")
	var scheme string
	var host string
	if len(cfgAddr) > 1 {
		scheme = cfgAddr[0]
		host = cfgAddr[1]
	}
	if len(cfgAddr) == 1 {
		scheme = SCHEME
		host = cfgAddr[0]
	}
	u := &url.URL{
		Scheme: scheme,
		Host:   host,
	}
	p := fmt.Sprintf("%s/%d", PATH, number)
	u = u.JoinPath(p)

	req, err := http.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	var order storage.OrderAccrual
	if err = json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return &order, nil
}
