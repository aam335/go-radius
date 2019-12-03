package radius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var attrs = []string{"Attr-Int", "Attr-Time", "VSA-Attr-Time"}

func TestNewAttrFilter(t *testing.T) {
	d := Dictionary{}
	d.RegisterN(td{})
	d.RegisterN(tdVS{})

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

/*
func makeTestPacket(d *Dictionary) {
	for _, x := range d.attributesByName {
		var val interface{}

		switch {
		case strings.Index(x.Name, "Int") != -1:
			val = "12345"
		case strings.Index(x.Name, "Str") != -1:
			val = "test value str"
		case strings.Index(x.Name, "Text") != -1:
			val = "test value text"
		case strings.Index(x.Name, "time") != -1:
			val = time.Unix(0, 0)
		}

	}

}

func TestFilter(t *testing.T) {
	d := Dictionary{}
	d.RegisterN(td{})
	d.RegisterN(tdVS{})

}
*/
