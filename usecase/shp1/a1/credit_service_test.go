package shp1

import "testing"

// CAUTION: THIS TEST IS ACTUALLY DO TRANSACTION
func TestCreditService_DoTransaction(t *testing.T) {
	type args struct {
		priceYen int
	}
	tests := []struct {
		name    string
		c       *CreditService
		args    args
		wantErr bool
	}{
		{
			name: "it should work",
			c:    NewCreditService(),
			args: args{
				// change price (once test is succeeded, it will cache)
				priceYen: 1000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreditService{}
			if err := c.DoTransaction(tt.args.priceYen); (err != nil) != tt.wantErr {
				t.Errorf("CreditService.DoTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
