package radius

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttrTransform(t *testing.T) {
	d := Dictionary{}
	assert.NoError(t, d.RegisterDC(td{}), "Normal Attrs")

	tests := []struct {
		n    string
		v    interface{}
		exp  interface{}
		text string
	}{
		{n: "Attr-Int", v: uint32(5555), exp: uint32(5555), text: "uint32 to uint32"},
		{n: "Attr-Int", v: "5555", exp: uint32(5555), text: "string to uint32"},
		{n: "Attr-Time", v: time.Unix(555555, 0), exp: time.Unix(555555, 0), text: "string to uint32"},
		{n: "Attr-Time", v: "555555", exp: time.Unix(555555, 0), text: "string to uint32"},
		{n: "Attr-Addr", v: net.ParseIP("10.11.12.13"), exp: net.IP{10, 11, 12, 13}, text: "string to uint32"},
		{n: "Attr-Addr", v: "10.11.12.13", exp: net.IP{10, 11, 12, 13}, text: "string to uint32"},
	}
	for _, tst := range tests {
		act, err := d.Attr(tst.n, tst.v)
		require.NoError(t, err, tst.text)
		require.Equal(t, tst.exp, act.Value, tst.text)
	}
}
