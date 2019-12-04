package radius

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dictEntryT struct {
	Vendor uint32
	Type   byte
	Name   string
}

func TestValueTagged(t *testing.T) {
	d := Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})
	intT := ValueTagged{Value: uint32(12345), Tag: 10}
	intStrT := ValueTagged{Value: "12345", Tag: 10}
	// Tagged val, tagged attr
	act, err := d.Attr("VSA-Attr-Int-Tag", intT)
	require.NoError(t, err, "Tagged val, tagged attr")
	require.NoError(t, attrCmp("Attr-Int-Tag", &Attribute{Tag: 10, Tagged: true, Type: 11, Vendor: 100, Value: uint32(12345)}, act), "Tagged val, tagged attr")

	// Tagged val as string, tagged attr
	act, err = d.Attr("VSA-Attr-Int-Tag", intStrT)
	require.NoError(t, err, "Tagged val, tagged attr, string")
	require.NoError(t, attrCmp("VSA-Attr-Int-Tag", &Attribute{Tag: 10, Tagged: true, Type: 11, Vendor: 100, Value: uint32(12345)}, act), "Tagged val, tagged attr")

	// tagged val, not tagged attr
	_, err = d.Attr("Attr-Int", intT)
	require.Error(t, err, "Tagged val, not tagged attr")
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
	_, err = d.Attr("Test-Attr-VSATAG", "testing")
	require.Error(t, err, "untagged value to tagged attr")
	//
	act, err = d.AttrTagged("Test-Attr-VSATAG", 99, "testingVSA")
	require.NoError(t, err, "existing attr w/VSA")
	require.NoError(t, attrCmp("Test-Attr-VSA", &Attribute{Tag: 99, Tagged: true, Type: 102, Vendor: 555, Value: "testingVSA"}, act), "Register to VSA w/tag")
}

func attrCmp(name string, exp, act *Attribute) error {
	if exp.Tag != act.Tag {
		return fmt.Errorf("Tag not equal %v exp: %v act:%v", name, exp.Tag, act.Tag)
	}
	if exp.Tagged != act.Tagged {
		return fmt.Errorf("Tagged not equal %v exp: %v act:%v", name, exp.Tagged, act.Tagged)
	}
	if exp.Vendor != act.Vendor {
		return fmt.Errorf("Vendor not equal %v exp: %v act:%v", name, exp.Vendor, act.Vendor)
	}
	if exp.Type != act.Type {
		return fmt.Errorf("Type not equal %v exp: %v act:%v", name, exp.Type, act.Type)
	}

	if !reflect.DeepEqual(exp.Value, act.Value) {
		return fmt.Errorf("Value not equal %v exp:%v(%v) act:%v(%v)", name, reflect.TypeOf(exp.Value), exp.Value, reflect.TypeOf(act.Value), act.Value)
	}
	return nil
}

func TestGetDict(t *testing.T) {
	d := Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})

	assert.NotNil(t, d.getDictEntryByName("Attr-Int"))
	assert.Nil(t, d.getDictEntryByName("Attr-Int-NotExists"))
	assert.NotNil(t, d.getDictEntryByName("VSA-Attr-Int-Tag"))
}
