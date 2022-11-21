package iso3166

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestIso3166ByAlpha3Code(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Iso3166
	}{
		{"004 Test", args{s: "004"}, nil},
		{"AO Test", args{s: "AO"}, nil},
		{"ABW Test", args{s: "ABW"}, &Iso3166{nameEn: "Aruba", nameFr: "Aruba", alpha2Code: "AW", alpha3Code: "ABW", numericCode: "533"}},
		{"OVO Test", args{s: "OVO"}, nil},
		{"BRA Test", args{s: "BRA"}, &Iso3166{nameEn: "Brazil", nameFr: "Brésil (le)", alpha2Code: "BR", alpha3Code: "BRA", numericCode: "076"}},
		{"ARG Test", args{s: "ARG"}, &Iso3166{nameEn: "Argentina", nameFr: "Argentine (l')", alpha2Code: "AR", alpha3Code: "ARG", numericCode: "032"}},
		{"TEST Test", args{s: "TEST"}, nil},
		{"GBR Test", args{s: "GBR"}, &Iso3166{nameEn: "United Kingdom of Great Britain and Northern Ireland (the)", nameFr: "Royaume-Uni de Grande-Bretagne et d'Irlande du Nord (le)", alpha2Code: "GB", alpha3Code: "GBR", numericCode: "826"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iso3166ByAlpha3Code(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIso3166ByAlpha3Code() = %#v, want %#v", got, tt.want)
				if tt.want != nil {
					require.Equal(t, got.alpha2Code, tt.want.alpha2Code)
					require.Equal(t, got.alpha3Code, tt.want.alpha3Code)
					require.Equal(t, got.numericCode, tt.want.numericCode)
					require.Equal(t, got.nameEn, tt.want.nameEn)
					require.Equal(t, got.nameFr, tt.want.nameFr)
				}
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

func TestIso3166ByAlpha2Code(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Iso3166
	}{
		{"004 Test", args{s: "004"}, nil},
		{"AO Test", args{s: "AO"}, nil},
		{"ABW Test", args{s: "SD"}, &Iso3166{nameEn: "Sudan (the)", nameFr: "Soudan (le)", alpha2Code: "SD", alpha3Code: "SDN", numericCode: "729"}},
		{"OVO Test", args{s: "OVO"}, nil},
		{"BRA Test", args{s: "BR"}, &Iso3166{nameEn: "Brazil", nameFr: "Brésil (le)", alpha2Code: "BR", alpha3Code: "BRA", numericCode: "076"}},
		{"ARG Test", args{s: "ARG"}, nil},
		{"TEST Test", args{s: "TEST"}, nil},
		{"GBR Test", args{s: "US"}, &Iso3166{nameEn: "United States of America (the)", nameFr: "États-Unis d'Amérique (les)", alpha2Code: "US", alpha3Code: "USA", numericCode: "840"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iso3166ByAlpha2Code(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIso3166ByAlpha2Code() = %#v, want %#v", got, tt.want)
				if tt.want != nil {
					require.Equal(t, got.alpha2Code, tt.want.alpha2Code)
					require.Equal(t, got.alpha3Code, tt.want.alpha3Code)
					require.Equal(t, got.numericCode, tt.want.numericCode)
					require.Equal(t, got.nameEn, tt.want.nameEn)
					require.Equal(t, got.nameFr, tt.want.nameFr)
				}
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

func TestIso3166ByNumericCode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Iso3166
	}{
		{"004 Test", args{s: "004"}, &Iso3166{nameEn: "Afghanistan", nameFr: "Afghanistan (l')", alpha2Code: "AF", alpha3Code: "AFG", numericCode: "004"}},
		{"AO Test", args{s: "AO"}, nil},
		{"ABW Test", args{s: "729"}, &Iso3166{nameEn: "Sudan (the)", nameFr: "Soudan (le)", alpha2Code: "SD", alpha3Code: "SDN", numericCode: "729"}},
		{"OVO Test", args{s: "OVO"}, nil},
		{"BRA Test", args{s: "076"}, &Iso3166{nameEn: "Brazil", nameFr: "Brésil (le)", alpha2Code: "BR", alpha3Code: "BRA", numericCode: "076"}},
		{"ARG Test", args{s: "ARG"}, nil},
		{"TEST Test", args{s: "TEST"}, nil},
		{"GBR Test", args{s: "840"}, &Iso3166{nameEn: "United States of America (the)", nameFr: "États-Unis d'Amérique (les)", alpha2Code: "US", alpha3Code: "USA", numericCode: "840"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iso3166ByNumericCode(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestIso3166ByNumericCode() = %#v, want %#v", got, tt.want)
				if tt.want != nil {
					require.Equal(t, got.alpha2Code, tt.want.alpha2Code)
					require.Equal(t, got.alpha3Code, tt.want.alpha3Code)
					require.Equal(t, got.numericCode, tt.want.numericCode)
					require.Equal(t, got.nameEn, tt.want.nameEn)
					require.Equal(t, got.nameFr, tt.want.nameFr)
				}
			}
		})
	}
}
