package vendor

import r "github.com/aam335/go-radius"

type dict struct {
	vendorID   uint32
	dictionary []r.DictionaryAttr
}

// Dict returns VendorID && Attributes
func (d dict) Dict() (uint32, []r.DictionaryAttr) {
	return d.vendorID, d.dictionary
}
