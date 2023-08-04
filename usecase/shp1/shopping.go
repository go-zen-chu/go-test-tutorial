package shp1

import (
	"fmt"
	"log"

	"github.com/go-zen-chu/go-test-tutorial/internal/bug"
)

type UseCaseShopping struct {
	Cart []Item
}

func NewUseCaseShopping() *UseCaseShopping {
	return &UseCaseShopping{
		Cart: make([]Item, 0),
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
	u.Cart = append(u.Cart, *item)
	return nil
}

func (u *UseCaseShopping) Buy() error {
	if len(u.Cart) == 0 {
		log.Println("no item in cart")
		return nil
	}
	// create credit service when cart has items
	cs := NewCreditService()
	totalPriceYen := 0
	for _, v := range u.Cart {
		totalPriceYen += v.PriceYen
	}
	if err := cs.DoTransaction(totalPriceYen); err != nil {
		return fmt.Errorf("while buying: %w", err)
	}
	return nil
}
