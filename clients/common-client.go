package clients

import (
	"encoding/json"
	"errors"
	model "github.com/richardgong1987/smartcloud-prices/models"
	"net/http"
	"time"
)

type SmartcloudClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewSmartcloudClient(baseURL string) *SmartcloudClient {
	return &SmartcloudClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *SmartcloudClient) FetchPrices(kind string) ([]model.Price, error) {
	url := c.BaseURL + "/prices?kind=" + kind
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch prices")
	}

	var smartcloudResp model.SmartcloudResponse
	if err := json.NewDecoder(resp.Body).Decode(&smartcloudResp); err != nil {
		return nil, err
	}

	return smartcloudResp.Prices, nil
}
