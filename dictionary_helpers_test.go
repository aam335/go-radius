package radius

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type td struct{}

func (td) Dict() (uint32, []DictionaryAttr) {
	return NotVSID, []DictionaryAttr{
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

type tdVS struct{}

const defTestVendor = 100 // vendorID for testing

func (tdVS) Dict() (uint32, []DictionaryAttr) {
	return defTestVendor, []DictionaryAttr{
		{Type: 1, Name: "VSA-Attr-Int", Tagged: false, Codec: AttributeInteger},
		{Type: 2, Name: "VSA-Attr-Str", Tagged: false, Codec: AttributeString},
		{Type: 3, Name: "VSA-Attr-Time", Tagged: false, Codec: AttributeTime},
		{Type: 4, Name: "VSA-Attr-Text", Tagged: false, Codec: AttributeText},
		{Type: 5, Name: "VSA-Attr-Addr", Tagged: false, Codec: AttributeAddress},
		{Type: 11, Name: "VSA-Attr-Int-Tag", Tagged: true, Codec: AttributeInteger},
		{Type: 12, Name: "VSA-Attr-Str-Tag", Tagged: true, Codec: AttributeString},
		{Type: 13, Name: "VSA-Attr-Time-Tag", Tagged: true, Codec: AttributeTime},
		{Type: 14, Name: "VSA-Attr-Text-Tag", Tagged: true, Codec: AttributeText},
		{Type: 15, Name: "VSA-Attr-Addr-Tag", Tagged: true, Codec: AttributeAddress},
	}
}

func TestRegisterN(t *testing.T) {
	d := Dictionary{}
	assert.NoError(t, d.RegisterN(td{}), "Normal Attrs")
	_, ok := d.Name(1)
	require.True(t, ok, "unable to find attr Type=1")
}
