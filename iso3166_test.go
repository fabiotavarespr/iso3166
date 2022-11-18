package iso3166

import (
	"reflect"
	"testing"
)

func TestExistsIso3166ByAlpha3Code(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"004 Test", args{s: "004"}, false},
		{"AO Test", args{s: "AO"}, false},
		{"ABW Test", args{s: "ABW"}, true},
		{"OVO Test", args{s: "OVO"}, false},
		{"BRA Test", args{s: "BRA"}, true},
		{"ARG Test", args{s: "ARG"}, true},
		{"TEST Test", args{s: "TEST"}, false},
		{"GBR Test", args{s: "GBR"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExistsIso3166ByAlpha3Code(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExistsIso3166ByAlpha3Code() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestExistsIso3166ByAlpha2Code(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"004 Test", args{s: "004"}, false},
		{"AO Test", args{s: "AO"}, true},
		{"OVO Test", args{s: "OVO"}, false},
		{"SA Test", args{s: "SD"}, true},
		{"BR Test", args{s: "BR"}, true},
		{"TEST Test", args{s: "TEST"}, false},
		{"US Test", args{s: "US"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExistsIso3166ByAlpha2Code(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExistsIso3166ByAlpha2Code() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestExistsIso3166ByNumericCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"004 Test", args{s: "004"}, true},
		{"BRA Test", args{s: "BRA"}, false},
		{"716 Test", args{s: "716"}, true},
		{"748 Test", args{s: "748"}, true},
		{"TEST Test", args{s: "TEST"}, false},
		{"U630 Test", args{s: "630"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExistsIso3166ByNumericCode(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExistsIso3166ByNumericCode() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
