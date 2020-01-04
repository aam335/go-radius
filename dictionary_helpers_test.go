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
	}
}

type tdVS struct{}

const defTestVendor = 100 // vendorID for testing

func (tdVS) Dict() (uint32, []DictionaryAttr) {
	return defTestVendor, []DictionaryAttr{
		{Type: 21, Name: "VSA-Attr-Int", Tagged: false, Codec: AttributeInteger},
		{Type: 22, Name: "VSA-Attr-Str", Tagged: false, Codec: AttributeString},
		{Type: 23, Name: "VSA-Attr-Time", Tagged: false, Codec: AttributeTime},
		{Type: 24, Name: "VSA-Attr-Text", Tagged: false, Codec: AttributeText},
		{Type: 25, Name: "VSA-Attr-Addr", Tagged: false, Codec: AttributeAddress},
		{Type: 11, Name: "VSA-Attr-Int-Tag", Tagged: true, Codec: AttributeInteger},
		{Type: 12, Name: "VSA-Attr-Str-Tag", Tagged: true, Codec: AttributeString},
		{Type: 13, Name: "VSA-Attr-Time-Tag", Tagged: true, Codec: AttributeTime},
		{Type: 14, Name: "VSA-Attr-Text-Tag", Tagged: true, Codec: AttributeText},
		{Type: 15, Name: "VSA-Attr-Addr-Tag", Tagged: true, Codec: AttributeAddress},
	}
}

func TestRegisterN(t *testing.T) {
	d := Dictionary{}
	assert.NoError(t, d.RegisterDC(td{}), "Normal Attrs")
	_, ok := d.Name(1)
	require.True(t, ok, "unable to find attr Type=1")
}

func TestDictionary_StrsToAttrs(t *testing.T) {
	d := Dictionary{}
	d.MustRegisterDC(td{})

	m := map[string]string{"Attr-Int": "1", "Attr-Int.0": "1", "Attr-Str": "string"}
	m1 := map[string]string{"Attr-Int2": "1", "Attr-Str": "string"}

	type args struct {
		m map[string]string
	}
	tests := []struct {
		name      string
		args      args
		wantAttrs []*Attribute
		wantErr   bool
	}{
		{name: "all attrs in Dict", args: args{m: m}, wantErr: false},
		{name: "attr not in Dict", args: args{m: m1}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := d.StrsToAttrs(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dictionary.StrsToAttrs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotAttrs, tt.wantAttrs) {
			// 	t.Errorf("Dictionary.StrsToAttrs() = %v, want %v", gotAttrs, tt.wantAttrs)
			// }
		})
	}
}
