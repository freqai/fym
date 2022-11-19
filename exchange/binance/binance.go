package binance

import (
	binance "github.com/adshao/go-binance/v2"
	"github.com/dqner/go-binance/v2/futures"
)

const (
	DefaultHost = "api.hbdm.com"
)

type Client struct {
	Label     string
	AccessKey string
	SecretKey string
	Host      string

	SpotAccountId int64
}

func New(label, accessKey, secretKey, host string) (*Client, error) {
	c := &Client{
		Label:     label,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
	if host != "" {
		c.Host = host
	} else {
		c.Host = DefaultHost
	}
	accountId, err := c.GetSpotAccountId()
	if err == nil {
		c.SpotAccountId = accountId
	} else {
		return nil, err
	}
	return c, nil
}

func (c *Client) Name() string {
	return "huobi"
}

func (c *Client) GetAccountInfo() ([]account.AccountInfo, error) {
	hb := new(client.AccountClient).Init(c.AccessKey, c.SecretKey, c.Host)
	return hb.GetAccountInfo()
}

func (c *Client) GetSpotAccountId() (int64, error) {
	accounts, err := c.GetAccountInfo()
	if err != nil {
		return 0, err
	}
	for _, a := range accounts {
		if a.Type == "spot" {
			return a.Id, nil
		}
	}
	return 0, nil
}
