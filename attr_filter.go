package radius

import (
	"fmt"
	"regexp"
	"strings"
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
func (a *AttrFilter) Filter(p *Packet) map[string]*Attribute {
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
	return filtered
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

// Keygen generates unique requester identifier for pub/sub and smart cache
// Keys MUST be registered by SetKeys. attrMap is a output of Filter
func (a *AttrFilter) Keygen(attrMap map[string]*Attribute) string {
	out := []string{}
	for _, key := range a.keys {
		if attr, ok := attrMap[key.attrName]; ok {
			val := fmt.Sprint(attr.Value)
			if key.re == nil {
				out = append(out, val)
				continue
			}
			fields := key.re.FindStringSubmatch(val)
			for _, fldNo := range key.fields {
				if len(fields) > fldNo {
					out = append(out, fields[fldNo])
				}
			}
		}
	}
	//b, _ := json.Marshal(out)
	b := strings.Join(out, "$")
	b = strings.ReplaceAll(b, " ", "_")
	return string(b)
}

// FilterStrings returns transport key and attributes in map[string]string
func (a *AttrFilter) FilterStrings(p *Packet) (key string, strAttrs map[string]string) {
	attrs := a.Filter(p)
	if len(a.keys) > 0 {
		key = a.Keygen(attrs)
	}
	strAttrs = make(map[string]string)
	for name, attr := range attrs {
		val := fmt.Sprint(attr.Value)
		strAttrs[name] = val
	}
	return
}
