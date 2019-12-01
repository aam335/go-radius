package radius

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type td struct{}

func (td) Dict() (uint32, []DictionaryAttr) {
	return 0, []DictionaryAttr{
		{Type: 1, Name: "Attr-Int", Tagged: false, Codec: AttributeInteger},
		{Type: 2, Name: "Attr-Str", Tagged: false, Codec: AttributeString},
		{Type: 3, Name: "Attr-Time", Tagged: false, Codec: AttributeTime},
		{Type: 4, Name: "Attr-Text", Tagged: false, Codec: AttributeText},
		{Type: 5, Name: "Attr-Addr", Tagged: false, Codec: AttributeAddress},
		{Type: 11, Name: "Attr-Int-Tag", Tagged: true, Codec: AttributeInteger},
		{Type: 12, Name: "Attr-Str-Tag", Tagged: true, Codec: AttributeString},
		{Type: 13, Name: "Attr-Time-Tag", Tagged: true, Codec: AttributeTime},
		{Type: 14, Name: "Attr-Text-Tag", Tagged: true, Codec: AttributeText},
		{Type: 15, Name: "Attr-Addr-Tag", Tagged: true, Codec: AttributeAddress},
	}
}

func TestRegisterN(t *testing.T) {
	d := Dictionary{}
	assert.NoError(t, d.RegisterN(td{}), "Normal Attrs")
	_, ok := d.Name(1)
	require.True(t, ok, "unable to find attr Type=1")
}
