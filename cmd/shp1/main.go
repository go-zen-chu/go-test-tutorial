package main

import (
	"github.com/go-zen-chu/go-test-tutorial/usecase/shp1"
)

func main() {
	sh := shp1.NewUseCaseShopping()
	if err := sh.AddItemToCart(&shp1.Item{
		ID:       1,
		Name:     "Super Expensive Book",
		PriceYen: 500000,
	}); err != nil {
		panic(err)
	}
	if err := sh.AddItemToCart(&shp1.Item{
		ID:       2,
		Name:     "Super Expensive Game",
		PriceYen: 800000,
	}); err != nil {
		panic(err)
	}
	if err := sh.Buy(); err != nil {
		panic(err)
	}
}
