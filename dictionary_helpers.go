package radius

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

// MustRegisterDC is a helper for RegisterN that panics if it returns an error.
func (d *Dictionary) MustRegisterDC(dict DictionaryContainer) {
	if err := d.RegisterDC(dict); err != nil {
		panic(err)
	}
}
