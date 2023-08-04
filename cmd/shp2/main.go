package main

import (
	"github.com/go-zen-chu/go-test-tutorial/usecase/shp2"
)

func shop(creditService shp2.CreditService) error {
	sh := shp2.NewUseCaseShopping(creditService)
	if err := sh.AddItemToCart(&shp2.Item{
		ID:       1,
		Name:     "Super Expensive Book",
		PriceYen: 500000,
	}); err != nil {
		return err
	}
	if err := sh.AddItemToCart(&shp2.Item{
		ID:       2,
		Name:     "Super Expensive Game",
		PriceYen: 800000,
	}); err != nil {
		return err
	}
	if err := sh.Buy(); err != nil {
		return err
	}
	return nil
}

func main() {
	cs := shp2.NewCreditService()
	if err := shop(cs); err != nil {
		panic(err)
	}
}
