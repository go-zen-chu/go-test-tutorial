package shp1

import (
	"fmt"
	"log"

	"github.com/go-zen-chu/go-test-tutorial/internal/bug"
)

type CreditService struct {
}

func NewCreditService() *CreditService {
	return &CreditService{}
}

func (c *CreditService) DoTransaction(priceYen int) error {
	if err := bug.Bug3(); err != nil {
		return fmt.Errorf("something went wrong: %w", err)
	}
	log.Printf("Thank you for buying!!! Price: %d", priceYen)
	return nil
}
