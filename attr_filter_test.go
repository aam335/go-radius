package radius

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var attrs = []string{"Attr-Int", "Attr-Time", "VSA-Attr-Time"}

func TestNewAttrFilter(t *testing.T) {
	d := Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})

	nf, err := d.NewAttrFilter(attrs)
	require.NoError(t, err)

	for _, a := range attrs {
		de := d.getDictEntryByName(a)
		key := uint64(de.Vendor)<<8 + uint64(de.Type)
		require.NotNil(t, nf.dictReduced[key])
		require.Equal(t, a, de.Name)
	}

	neattrs := []string{"Attr-Not-Exists"}
	_, err = d.NewAttrFilter(neattrs)
	require.Error(t, err)

}

func makeTestPacket(d *Dictionary) (*Packet, error) {
	p := New(CodeAccessAccept, secret)
	p.Dictionary = d
	for _, x := range d.attributesByName {
		var val interface{}
		switch {
		case strings.Index(x.Name, "Int") != -1:
			val = "12345"
		case strings.Index(x.Name, "Str") != -1:
			val = "test value str"
		case strings.Index(x.Name, "Text") != -1:
			val = "test value text"
		case strings.Index(x.Name, "Time") != -1:
			val = time.Unix(0, 0)
		case strings.Index(x.Name, "Addr") != -1:
			val = "172.17.18.19"
		}
		if x.Tagged {
			if err := p.Add(x.Name, ValueTagged{Tag: 10, Value: val}); err != nil {
				return nil, err
			}
			continue
		}
		if err := p.Add(x.Name, val); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func TestFilter(t *testing.T) {
	d := &Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})
	p, err := makeTestPacket(d)
	require.NoError(t, err)
	nf, err := d.NewAttrFilter(attrs)
	require.NoError(t, err)
	filteredAssoc := nf.Filter(p)

	for _, attrName := range attrs {
		assert.NotNilf(t, filteredAssoc[attrName], "%v not in reply", attrName)
	}

	attrName := "Attr-Int"
	copys := 10
	for i := 0; i < copys; i++ {
		assert.NoError(t, p.Add(attrName, uint32(555))) // duplicate assoc
	}

	filteredAssoc = nf.Filter(p)

	assert.NotNilf(t, filteredAssoc[attrName], "%v not in reply", attrName)
	for i := 0; i < copys; i++ {
		serialName := fmt.Sprint(attrName, ".", i)
		assert.NotNilf(t, filteredAssoc[serialName], "%v not in reply", serialName)
	}

}

var keysAttrs = []string{"Attr-Int", "Attr-Time", "VSA-Attr-Time", "Attr-Text"}

func TestAttrFilter_SetKeys(t *testing.T) {
	d := &Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})
	nf, err := d.NewAttrFilter(keysAttrs)
	require.NoError(t, err)

	type fields struct {
		dictReduced map[uint64]*dictEntry
		keys        []key
	}
	type args struct {
		keys []OneKey
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "add not exist on empty filter", fields: fields{}, args: args{keys: []OneKey{{Name: "not exists"}}}, wantErr: true},
		{name: "add not exist", fields: fields{dictReduced: nf.dictReduced}, args: args{keys: []OneKey{{Name: "not exists"}}}, wantErr: true},
		{name: "add exist", fields: fields{dictReduced: nf.dictReduced}, args: args{keys: []OneKey{{Name: "Attr-Text"}}}, wantErr: false},
		{name: "add exists regexp", fields: fields{dictReduced: nf.dictReduced}, args: args{keys: []OneKey{{Name: "Attr-Text", Regexp: `(\w+)`, Fields: []int{0, 2}}}}, wantErr: false},
		{name: "add exists regexp errored", fields: fields{dictReduced: nf.dictReduced}, args: args{keys: []OneKey{{Name: "Attr-Text", Regexp: `(\w+`, Fields: []int{0, 2}}}}, wantErr: true},
		{name: "add exists dual", fields: fields{dictReduced: nf.dictReduced}, args: args{keys: []OneKey{{Name: "Attr-Text"}, {Name: "Attr-Text", Regexp: `(\w+)`, Fields: []int{0, 2}}}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AttrFilter{
				dictReduced: tt.fields.dictReduced,
				keys:        tt.fields.keys,
			}
			if err := a.SetKeys(tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("AttrFilter.SetKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// p, err := makeTestPacket(d)
	// require.NoError(t, err)

}

func TestAttrFilter_keygen(t *testing.T) {
	d := &Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})
	nf, err := d.NewAttrFilter(keysAttrs)
	require.NoError(t, err)
	p, err := makeTestPacket(d)
	require.NoError(t, err)

	filtered := nf.Filter(p)
	require.NoError(t, nf.SetKeys([]OneKey{{Name: "Attr-Text"}, {Name: "Attr-Int"}, {Name: "Attr-Text", Regexp: `(\w+)\s+\w+\s+(\w+)`, Fields: []int{1, 2}}}))

	type fields struct {
		dictReduced map[uint64]*dictEntry
		keys        []key
	}
	type args struct {
		attrMap map[string]*Attribute
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{name: "nils", want: ""},
		{name: "nil map", fields: fields{keys: []key{{attrName: "Attr-Text"}}}, want: ""},
		{name: "text", fields: fields{keys: []key{{attrName: "Attr-Text"}}}, args: args{filtered}, want: `test_value_text`},
		{name: "text", fields: fields{keys: nf.keys}, args: args{filtered}, want: `test_value_text$12345$test$text`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AttrFilter{
				dictReduced: tt.fields.dictReduced,
				keys:        tt.fields.keys,
			}
			if got := a.Keygen(tt.args.attrMap); got != tt.want {
				t.Errorf("AttrFilter.keygen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttrFilter_FilterStrings(t *testing.T) {
	d := &Dictionary{}
	d.RegisterDC(td{})
	d.RegisterDC(tdVS{})
	p, err := makeTestPacket(d)
	require.NoError(t, err)
	nf, err := d.NewAttrFilter(keysAttrs)
	require.NoError(t, err)
	err = nf.SetKeys([]OneKey{{Name: "Attr-Text"}, {Name: "Attr-Int"}})
	require.NoError(t, err)

	type args struct {
		p *Packet
	}
	tests := []struct {
		name         string
		args         args
		wantKey      string
		wantStrAttrs map[string]string
	}{
		// TODO: Add test cases.
		{name: "1st", args: args{p}, wantKey: `test_value_text$12345`,
			wantStrAttrs: map[string]string{"Attr-Int": "12345", "Attr-Text": "test value text", "Attr-Time": "1970-01-01 03:00:00 +0300 MSK", "VSA-Attr-Time": "1970-01-01 03:00:00 +0300 MSK"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := nf
			gotKey, gotStrAttrs := a.FilterStrings(tt.args.p)

			if gotKey != tt.wantKey {
				t.Errorf("AttrFilter.FilterStrings() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if !reflect.DeepEqual(gotStrAttrs, tt.wantStrAttrs) {
				t.Errorf("AttrFilter.FilterStrings() gotStrAttrs = %v, want %v", gotStrAttrs, tt.wantStrAttrs)
			}
		})
	}
}
