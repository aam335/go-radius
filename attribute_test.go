package radius

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttrTransform(t *testing.T) {
	d := Dictionary{}
	assert.NoError(t, d.RegisterN(td{}), "Normal Attrs")

	tests := []struct {
		n    string
		v    interface{}
		exp  Attribute
		text string
	}{
		{n: "Attr-Int", v: uint32(5555), exp: Attribute{Tag: 0, Tagged: false, Type: 1, Vendor: 0, Value: uint32(5555)}, text: "uint32 to uint32"},
		{n: "Attr-Int", v: "5555", exp: Attribute{Tag: 0, Tagged: false, Type: 1, Vendor: 0, Value: uint32(5555)}, text: "string to uint32"},
	}
	for _, tst := range tests {
		act, err := d.Attr(tst.n, tst.v)
		require.NoError(t, err, tst.text)
		require.NoError(t, attrCmp("Attr-Int-Tag", &tst.exp, act), tst.text)
	}
}
