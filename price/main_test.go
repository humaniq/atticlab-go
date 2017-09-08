package price_test

import (
	"testing"

	"github.com/humaniq/atticlab-go/price"
	"github.com/humaniq/atticlab-go/xdr"
)

var Tests = []struct {
	S string
	P xdr.Price
}{
	{"0.1", xdr.Price{1, 10}},
	{"0.01", xdr.Price{1, 100}},
	{"0.001", xdr.Price{1, 1000}},
	{"543.017930", xdr.Price{54301793, 100000}},
	{"319.69983", xdr.Price{31969983, 100000}},
	{"0.93", xdr.Price{93, 100}},
	{"0.5", xdr.Price{1, 2}},
	{"1.730", xdr.Price{173, 100}},
	{"0.85334384", xdr.Price{5333399, 6250000}},
	{"5.5", xdr.Price{11, 2}},
	{"2.72783", xdr.Price{272783, 100000}},
	{"638082.0", xdr.Price{638082, 1}},
	{"2.93850088", xdr.Price{36731261, 12500000}},
	{"58.04", xdr.Price{1451, 25}},
	{"41.265", xdr.Price{8253, 200}},
	{"5.1476", xdr.Price{12869, 2500}},
	{"95.14", xdr.Price{4757, 50}},
	{"0.74580", xdr.Price{3729, 5000}},
	{"4119.0", xdr.Price{4119, 1}},
}

func TestParse(t *testing.T) {
	for _, v := range Tests {
		o, err := price.Parse(v.S)
		if err != nil {
			t.Errorf("Couldn't parse %s: %v+", v.S, err)
			continue
		}

		if o.N != v.P.N || o.D != v.P.D {
			t.Errorf("%s parsed to %d, not %d", v.S, o, v.P)
		}
	}

	_, err := price.Parse("0.0000000003")
	if err == nil {
		t.Error("Expected error")
	}

	_, err = price.Parse("2147483649")
	if err == nil {
		t.Error("Expected error")
	}
}
