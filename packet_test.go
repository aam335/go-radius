package radius

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tstattrs = []struct {
	n   string
	v   interface{}
	exp Attribute
}{
	{n: "Attr-Int", v: uint32(5555), exp: Attribute{Tag: 0, Tagged: false, Type: 1, Vendor: 0, Value: uint32(5555)}},
	{n: "Attr-Time", v: time.Unix(555555, 0), exp: Attribute{Tag: 0, Tagged: false, Type: 3, Vendor: 0, Value: time.Unix(555555, 0)}},
	{n: "Attr-Addr", v: net.IP{10, 11, 12, 13}, exp: Attribute{Tag: 0, Tagged: false, Type: 5, Vendor: 0, Value: net.IP{10, 11, 12, 13}}},

	{n: "VSA-Attr-Int", v: uint32(5555), exp: Attribute{Tag: 0, Tagged: false, Type: 21, Vendor: defTestVendor, Value: uint32(5555)}},
	{n: "VSA-Attr-Time", v: time.Unix(555555, 0), exp: Attribute{Tag: 0, Tagged: false, Type: 23, Vendor: defTestVendor, Value: time.Unix(555555, 0)}},
	{n: "VSA-Attr-Addr", v: net.IP{10, 11, 12, 13}, exp: Attribute{Tag: 0, Tagged: false, Type: 25, Vendor: defTestVendor, Value: net.IP{10, 11, 12, 13}}},

	{n: "VSA-Attr-Int-Tag", v: ValueTagged{Tag: 10, Value: uint32(5555)}, exp: Attribute{Tag: 10, Tagged: true, Type: 11, Vendor: defTestVendor, Value: uint32(5555)}},
	{n: "VSA-Attr-Time-Tag", v: ValueTagged{Tag: 10, Value: time.Unix(555555, 0)}, exp: Attribute{Tag: 10, Tagged: true, Type: 13, Vendor: defTestVendor, Value: time.Unix(555555, 0)}},
	{n: "VSA-Attr-Addr-Tag", v: ValueTagged{Tag: 10, Value: net.IP{10, 11, 12, 13}}, exp: Attribute{Tag: 10, Tagged: true, Type: 15, Vendor: defTestVendor, Value: net.IP{10, 11, 12, 13}}},
}

var secret = []byte("VerySecret")

func TestPacketEncode(t *testing.T) {
	d := Dictionary{}
	require.NoError(t, d.RegisterDC(td{}), "load attrs")
	require.NoError(t, d.RegisterDC(tdVS{}), "load attrs")

	p := New(CodeAccessRequest, secret)
	p.Dictionary = &d

	for _, a := range tstattrs {
		assert.NoErrorf(t, p.Add(a.n, a.v), "unable to add %v", a.n)
	}

	data, err := p.Encode()
	require.NoError(t, err, "encode")

	// decode & compare results
	rp, err := Parse(data, secret, &d)
	if err != nil {
		log.Fatal(err)
	}

	for _, attr := range rp.Attributes {
		attrName, ok := rp.Dictionary.NameVID(attr.Vendor, attr.Type)
		require.True(t, ok, "can't resolve attribute name")
		exp := expectValue(attrName)
		require.NotNil(t, exp, "corrupted attr")
		require.NoError(t, attrCmp(attrName, exp, attr))
	}

}

func expectValue(name string) *Attribute {
	for _, a := range tstattrs {
		if a.n == name {
			return &a.exp
		}
	}
	return nil
}
