package shp1

import (
	"testing"
)

func TestUseCaseShopping_AddItemToCart(t *testing.T) {
	type fields struct {
		Cart []Item
	}
	type args struct {
		item *Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "it should success",
			fields: fields{
				Cart: []Item{},
			},
			args: args{
				item: &Item{
					Name:     "test",
					PriceYen: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCaseShopping{
				Cart: tt.fields.Cart,
			}
			if err := u.AddItemToCart(tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseShopping.AddItemToCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseShopping_Buy(t *testing.T) {
	type fields struct {
		Cart []Item
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "it should success",
			fields: fields{
				Cart: []Item{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCaseShopping{
				Cart: tt.fields.Cart,
			}
			if err := u.Buy(); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseShopping.Buy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
