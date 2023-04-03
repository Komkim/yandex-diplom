package updateorder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"yandex-diplom/config"
	storage "yandex-diplom/storage/repository"
)

const (
	SCHEME = "http"
	PATH   = "api/orders/"
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

func (c *MyClient) GetAccrual(number int64) (*storage.Order, error) {
	u := &url.URL{
		Scheme: SCHEME,
		Host:   c.config.AccrualAddress,
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

	//req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	var order storage.Order
	if err = json.NewDecoder(resp.Body).Decode(&order); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return &order, nil
}
