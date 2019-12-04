package radius

import "fmt"

// AttrFilter is used for smart decoding of attributes from a package.
type AttrFilter struct {
	dictReduced map[uint64]*dictEntry
}

// NewAttrFilter compiles attribute names into AttrFilter
func (d *Dictionary) NewAttrFilter(names []string) (*AttrFilter, error) {
	dictReduced := make(map[uint64]*dictEntry)
	for _, name := range names {
		entry := d.getDictEntryByName(name)
		if entry == nil {
			return nil, fmt.Errorf("Attribute %v not exists in dictionary", name)
		}
		// uint64 simple index, Type size 8 bits
		key := uint64(entry.Vendor)<<8 + uint64(entry.Type)
		dictReduced[key] = entry
	}
	return &AttrFilter{dictReduced: dictReduced}, nil
}

// Filter intersect packet attributes & filter, than run
func (a *AttrFilter) Filter(p *Packet) (map[string]*Attribute, error) {
	filtered := make(map[string]*Attribute)
	for _, attr := range p.Attributes {
		key := uint64(attr.Vendor)<<8 + uint64(attr.Type)
		if dictRec, ok := a.dictReduced[key]; ok {
			name := dictRec.Name
			id := 0
			for {
				if _, ok = filtered[name]; !ok {
					filtered[name] = attr
					break
				}
				name = fmt.Sprint(dictRec.Name, ".", id)
				id++
			}
		}
	}
	return filtered, nil
}
