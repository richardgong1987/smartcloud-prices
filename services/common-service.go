package services

import (
	client "github.com/richardgong1987/smartcloud-prices/clients"
	model "github.com/richardgong1987/smartcloud-prices/models"
)

type SmartcloudPriceService struct {
}

var request = client.NewSmartcloudClient("http://localhost:8080")

func (s *SmartcloudPriceService) GetPrices(kind string) ([]model.Price, error) {
	prices, err := request.FetchPrices(kind)
	if err != nil {
		// Handle errors gracefully
		return nil, err
	}
	return prices, nil
}
