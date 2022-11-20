package binance

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
	c.Host = ""
	return c, nil
}

func (c *Client) Name() string {
	return "binance"
}

/*
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
*/
