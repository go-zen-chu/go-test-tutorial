package shp2

import "log"

type mockCreditService struct{}

func NewMockCreditService() CreditService {
	return &mockCreditService{}
}

func (c *mockCreditService) DoTransaction(priceYen int) error {
	log.Printf("[test] be relieved, this is test: %d", priceYen)
	return nil
}
