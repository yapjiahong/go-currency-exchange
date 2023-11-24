//go:build unit

package common

import "testing"

func TestFormatAmount(t *testing.T) {
	type args struct {
		amount string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1 less than one thousand",
			args: args{
				amount: "321.0",
			},
			want: "321.0",
		},
		{
			name: "#2 less than ten thousand",
			args: args{
				amount: "6321.0",
			},
			want: "6,321.0",
		},
		{
			name: "#3 greater than one million",
			args: args{
				amount: "1000000.12",
			},
			want: "1,000,000.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatAmount(tt.args.amount); got != tt.want {
				t.Errorf("FormatAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
