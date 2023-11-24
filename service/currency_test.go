//go:build unit

package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertCurrency(t *testing.T) {
	type args struct {
		params CurrencyDTO
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "#1 normal case",
			args: args{
				CurrencyDTO{
					From:   TWD,
					To:     JPY,
					Amount: float64(3),
				},
			},
			want:    "11.01",
			wantErr: nil,
		},
		{
			name: "#2 unsupported currency",
			args: args{
				CurrencyDTO{
					From: 100,
					To:   101,
				},
			},
			wantErr: fmt.Errorf("unsupported currency conversion: 100 to 101"),
		},
		{
			name: "#3 input with 0",
			args: args{
				CurrencyDTO{
					From:   USD,
					To:     TWD,
					Amount: float64(0),
				},
			},
			want:    "0.00",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertCurrency(tt.args.params)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
