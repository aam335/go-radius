package radius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type dictEntryT struct {
	Vendor uint32
	Type   byte
	Name   string
}

func TestRegister(t *testing.T) {
	d := Dictionary{}
	require.NoError(t, d.Register("Test-Attr", 100, AttributeString), "Simple regiter to main")
	require.Error(t, d.Register("Test-Attr", 101, AttributeString), "Differ name, same type")
	require.Error(t, d.Register("Test-AttrDupType", 100, AttributeString), "Differ name, same type")
	require.Error(t, d.VsaRegister(555, "Test-Attr", 100, AttributeString), "Register to VSA w/duplicate name")
	require.NoError(t, d.VsaRegister(555, "Test-Attr-VSA", 100, AttributeString), "Register to VSA")
	require.NoError(t, d.VsaRegister(555, "Test-Attr-VSA1", 101, AttributeString), "Register to VSA")

	a, ok := d.Type("Test-Attr")
	require.True(t, ok, "got attr by name")
	require.Equal(t, uint8(100), a, "type to name")

	b, ok := d.Name(100)
	require.True(t, ok, "got name by attr")
	require.Equal(t, "Test-Attr", b, "name to type")

	// old interface shouldn't use VSA attr as native
	_, ok = d.Type("Test-Attr-VSA")
	require.False(t, ok, "VSA Name as native")
	_, ok = d.Name(101)
	require.False(t, ok, "VSA Type as native")

	//
}
