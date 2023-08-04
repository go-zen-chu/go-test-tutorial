package shp2

import (
	"fmt"
	"log"

	"github.com/go-zen-chu/go-test-tutorial/internal/bug"
)

type CreditService interface {
	DoTransaction(priceYen int) error
}

type creditService struct {
}

func NewCreditService() CreditService {
	return &creditService{}
}

func (c *creditService) DoTransaction(priceYen int) error {
	if err := bug.Bug3(); err != nil {
		return fmt.Errorf("something went wrong: %w", err)
	}
	log.Printf("Thank you for buying!!! Price: %d", priceYen)
	return nil
}
