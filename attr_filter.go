package radius

import (
	"fmt"
	"regexp"
)

// AttrFilter is used for smart decoding of attributes from a package.
type AttrFilter struct {
	dictReduced map[uint64]*dictEntry
	keys        []key
}

// OneKey expands one part of the key attribute
type OneKey struct {
	Name   string
	Regexp string
	Fields []int
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

type key struct {
	attrName string
	re       *regexp.Regexp
	fields   []int
}

// SetKeys sets keys into filter. Generated key used for querying sql backend
// and smart caching
func (a *AttrFilter) SetKeys(keys []OneKey) (err error) {
	a.keys = a.keys[:0]
	attrs := make(map[string]bool)
	for _, attr := range a.dictReduced {
		attrs[attr.Name] = true
	}
	for _, k := range keys {
		if _, ok := attrs[k.Name]; !ok {
			return fmt.Errorf("key '%v' not in filtered attrs", k.Name)
		}
		key := key{attrName: k.Name}
		if k.Regexp != "" {
			if key.re, err = regexp.Compile(k.Regexp); err != nil {
				return
			}
			key.fields = make([]int, len(k.Fields))
			copy(key.fields, k.Fields)
		}
		a.keys = append(a.keys, key)
	}
	return nil
}
