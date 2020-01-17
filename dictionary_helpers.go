package radius

import (
	"strings"
)

// DictionaryAttr structure for mass Attribute import
type DictionaryAttr struct {
	Type   byte
	Name   string
	Tagged bool
	Codec  AttributeCodec
}

// DictionaryContainer additional dictionaries
type DictionaryContainer interface {
	Dict() (vendorID uint32, attrs []DictionaryAttr)
}

// RegisterDC register attributes from DictionaryContainer
func (d *Dictionary) RegisterDC(dict DictionaryContainer) error {
	vendorID, nAttr := dict.Dict()
	for _, a := range nAttr {
		if err := d.VsaRegisterTagFlag(vendorID, a.Name, a.Type, a.Tagged, a.Codec); err != nil {
			return err
		}

	}
	return nil
}

// MustRegisterDC is a helper for RegisterDC that panics if it returns an error.
func (d *Dictionary) MustRegisterDC(dict DictionaryContainer) {
	if err := d.RegisterDC(dict); err != nil {
		panic(err)
	}
}

// StrsToAttrs makes []*Attribute from map[string]string
// this suitable for reply from sql backend etc...
// returns last error
func (d *Dictionary) StrsToAttrs(m map[string]string) (attrs []*Attribute, err error) {
	attrs = []*Attribute{}
	var a *Attribute
	var lastErr error
	for name, val := range m {
		if idx := strings.Index(name, "."); idx > 0 {
			name = name[:idx]
		}
		if a, lastErr = d.Attr(name, val); lastErr == nil {
			// log.Printf("Error:%v (%v:%v)", err, name, val)
			// return
			attrs = append(attrs, a)
		} else {
			err = lastErr
			continue
		}
	}
	return
}
