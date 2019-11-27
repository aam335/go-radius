package radius

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type dictEntryT struct {
	Vendor uint32
	Type   byte
	Name   string
}

func TestEncodings(t *testing.T) {
	d := Dictionary{}
	require.NoError(t, d.Register("Test-Attr-Int", 100, AttributeInteger), "Simple regiter to main")
	// require.NoError(t, d.Register("Test-Attr-String", 100, AttributeInteger), "Simple regiter to main")
	// require.NoError(t, d.Register("Test-Attr", 100, AttributeInteger), "Simple regiter to main")
	// require.NoError(t, d.Register("Test-Attr", 100, AttributeInteger), "Simple regiter to main")
	// require.NoError(t, d.Register("Test-Attr", 100, AttributeInteger), "Simple regiter to main")
	_, err := d.Attr("Test-Attr-Int", "123")
	require.NoError(t, err, "normal flow")

}

func TestRegister(t *testing.T) {
	d := Dictionary{}
	require.NoError(t, d.Register("Test-Attr", 100, AttributeString), "Simple regiter to main")
	require.Error(t, d.Register("Test-Attr", 101, AttributeString), "Differ name, same type")
	require.Error(t, d.Register("Test-AttrDupType", 100, AttributeString), "Differ name, same type")
	require.Error(t, d.VsaRegister(555, "Test-Attr", 100, AttributeString), "Register to VSA w/duplicate name")
	require.NoError(t, d.VsaRegister(555, "Test-Attr-VSA", 100, AttributeString), "Register to VSA")
	require.NoError(t, d.VsaRegister(555, "Test-Attr-VSA1", 101, AttributeString), "Register to VSA")
	require.Error(t, d.VsaRegisterTag(555, "Test-Attr-VSATAG", 101, AttributeString), "Register to VSA w/tag")
	require.NoError(t, d.VsaRegisterTag(555, "Test-Attr-VSATAG", 102, AttributeString), "Register to VSA w/tag")
	require.NoError(t, d.VsaRegisterTagFlag(555, "Test-Attr-VSATAG1", 103, true, AttributeString), "Register to VSA w/tag")

	// old interface methods
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

	// test untagged Attributes

	// Attr exists
	act, err := d.Attr("Test-Attr", "testing")
	require.NoError(t, err, "existing attr")
	require.NoError(t, attrCmp("Test-Attr", &Attribute{Tag: 0, Tagged: false, Type: 100, Vendor: 0, Value: "testing"}, act), "Register to VSA w/tag")

	// Attr not exists
	_, err = d.Attr("Test-Attr2", "testing")
	require.Error(t, err, "not existing attr")

	// Attr exists and is VSA
	act, err = d.Attr("Test-Attr-VSA", "testingVSA")
	require.NoError(t, err, "existing attr w/VSA")
	require.NoError(t, attrCmp("Test-Attr-VSA", &Attribute{Tag: 0, Tagged: false, Type: 100, Vendor: 555, Value: "testingVSA"}, act), "Register to VSA w/tag")

	// Attr untagged method to tagged Attr
	act, err = d.Attr("Test-Attr-VSATAG", "testing")
	require.Error(t, err, "untagged value to tagged attr")

	//
	act, err = d.AttrTagged("Test-Attr-VSATAG", 99, "testingVSA")
	require.NoError(t, err, "existing attr w/VSA")
	require.NoError(t, attrCmp("Test-Attr-VSA", &Attribute{Tag: 99, Tagged: true, Type: 102, Vendor: 555, Value: "testingVSA"}, act), "Register to VSA w/tag")

}

func attrCmp(name string, exp, act *Attribute) error {
	if exp.Tag != act.Tag {
		return fmt.Errorf("Tag not equal exp: %v act:%v", exp.Tag, act.Tag)
	}
	if exp.Tagged != act.Tagged {
		return fmt.Errorf("Tagged not equal exp: %v act:%v", exp.Tagged, act.Tagged)
	}
	if exp.Vendor != act.Vendor {
		return fmt.Errorf("Vendor not equal exp: %v act:%v", exp.Vendor, act.Vendor)
	}
	if exp.Type != act.Type {
		return fmt.Errorf("Type not equal exp: %v act:%v", exp.Type, act.Type)
	}

	if !reflect.DeepEqual(exp, act) {
		return fmt.Errorf("Value not equal")
	}
	return nil
}
