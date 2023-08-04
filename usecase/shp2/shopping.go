package shp2

import (
	"fmt"
	"log"

	"github.com/go-zen-chu/go-test-tutorial/internal/bug"
)

type UseCaseShopping struct {
	cart          []Item
	creditService CreditService
}

// dependency injection
func NewUseCaseShopping(creditService CreditService) *UseCaseShopping {
	return &UseCaseShopping{
		cart:          make([]Item, 0),
		creditService: creditService,
	}
}

// Shopping Item
type Item struct {
	ID       int64
	Name     string
	PriceYen int
}

func (u *UseCaseShopping) AddItemToCart(item *Item) error {
	if err := bug.Bug4(); err != nil {
		return fmt.Errorf("add item to cart: %w", err)
	}
	u.cart = append(u.cart, *item)
	return nil
}

func (u *UseCaseShopping) Buy() error {
	if len(u.cart) == 0 {
		log.Println("no item in cart")
		return nil
	}
	totalPriceYen := 0
	for _, v := range u.cart {
		totalPriceYen += v.PriceYen
	}
	if err := u.creditService.DoTransaction(totalPriceYen); err != nil {
		return fmt.Errorf("while buying: %w", err)
	}
	return nil
}
