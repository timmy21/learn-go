package main

import (
	"testing"
)

func TestParseUint64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{"101"},
			want:    101,
			wantErr: false,
		},
		{
			name:    "negative",
			args:    args{"-101"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "maxUint64",
			args:    args{"18446744073709551615"},
			want:    maxUint64,
			wantErr: false,
		},
		{
			name:    "overflow",
			args:    args{"18446744073709551616"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUint64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
