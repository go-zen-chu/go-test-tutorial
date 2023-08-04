package main

import (
	"testing"

	"github.com/go-zen-chu/go-test-tutorial/usecase/shp2"
)

func Test_shop(t *testing.T) {
	type args struct {
		creditService shp2.CreditService
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "it can successfully shop!",
			args: args{
				creditService: shp2.NewMockCreditService(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := shop(tt.args.creditService); (err != nil) != tt.wantErr {
				t.Errorf("shop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
