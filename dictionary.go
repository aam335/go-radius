package radius

import (
	"errors"
	"fmt"
	"sync"
)

var builtinOnce sync.Once

// Builtin is the built-in dictionary. It is initially loaded with the
// attributes defined in RFC 2865 and RFC 2866.
var Builtin *Dictionary

func initDictionary() {
	Builtin = &Dictionary{}
}

type dictEntry struct {
	Vendor uint32
	Type   byte
	Name   string
	Codec  AttributeCodec
}

type attributes [256]*dictEntry

// Dictionary stores mappings between attribute names and types and
// AttributeCodecs.
type Dictionary struct {
	mu sync.RWMutex
	// attributesByType [256]*dictEntry
	attributesByName map[string]*dictEntry
	attributesByType map[uint32]*attributes
}

// VsaRegister registers the AttributeCodec for the given attribute name and type.
func (d *Dictionary) VsaRegister(vendorID uint32, name string, t byte, codec AttributeCodec) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.attributesByType == nil {
		d.attributesByType = make(map[uint32]*attributes)
	}
	attributesByType, ok := d.attributesByType[vendorID]
	if !ok {
		attributesByType = new(attributes)
		d.attributesByType[vendorID] = attributesByType
	}

	if attributesByType[t] != nil {

		return fmt.Errorf("radius: attribute with Type '%v' already registered", t)
	}
	entry := &dictEntry{
		Vendor: vendorID,
		Type:   t,
		Name:   name,
		Codec:  codec,
	}

	if d.attributesByName == nil {
		d.attributesByName = make(map[string]*dictEntry)
	}
	if _, ok := d.attributesByName[name]; ok {
		return fmt.Errorf("radius: attribute with Name '%v' already registered", name)
	}

	d.attributesByType[vendorID][t] = entry
	d.attributesByName[name] = entry
	return nil
}

// Register registers the AttributeCodec for the given attribute name and type.
func (d *Dictionary) Register(name string, t byte, codec AttributeCodec) error {
	return d.VsaRegister(0, name, t, codec)
}

// MustRegister is a helper for Register that panics if it returns an error.
func (d *Dictionary) MustRegister(name string, t byte, codec AttributeCodec) {
	if err := d.Register(name, t, codec); err != nil {
		panic(err)
	}
}

// VsaMustRegister is a helper for Register that panics if it returns an error.
func (d *Dictionary) VsaMustRegister(vendorID uint32, name string, t byte, codec AttributeCodec) {
	if err := d.VsaRegister(vendorID, name, t, codec); err != nil {
		panic(err)
	}
}

// func (d *Dictionary) get(name string) (t byte, codec AttributeCodec, ok bool) {
// 	d.mu.RLock()
// 	entry := d.attributesByName[name]
// 	d.mu.RUnlock()
// 	if entry == nil {
// 		return
// 	}
// 	t = entry.Type
// 	codec = entry.Codec
// 	ok = true
// 	return
// }

// Attr returns a new *Attribute whose type is registered under the given
// name.
//
// If name is not registered, nil and an error is returned.
//
// If the attribute's codec implements AttributeTransformer, the value is
// first transformed before being stored in *Attribute. If the transform
// function returns an error, nil and the error is returned.
func (d *Dictionary) Attr(name string, value interface{}) (*Attribute, error) {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil {
		return nil, errors.New("radius: attribute name not registered")
	}
	if transformer, ok := entry.Codec.(AttributeTransformer); ok {
		transformed, err := transformer.Transform(value)
		if err != nil {
			return nil, err
		}
		value = transformed
	}
	return &Attribute{
		Vendor: entry.Vendor,
		Type:   entry.Type,
		Value:  value,
	}, nil
}

// MustAttr is a helper for Attr that panics if Attr were to return an error.
func (d *Dictionary) MustAttr(name string, value interface{}) *Attribute {
	attr, err := d.Attr(name, value)
	if err != nil {
		panic(err)
	}
	return attr
}

// Name returns the registered name for the given attribute type. ok is false
// if the given type is not registered.
func (d *Dictionary) Name(t byte) (name string, ok bool) {
	d.mu.RLock()
	entry := d.attributesByType[0][t]
	d.mu.RUnlock()
	if entry == nil || entry.Vendor != 0 {
		return
	}
	name = entry.Name
	ok = true
	return
}

// Type returns the registered type for the given attribute name. ok is false
// if the given name is not registered.
func (d *Dictionary) Type(name string) (t byte, ok bool) {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil || entry.Vendor != 0 {
		return
	}
	t = entry.Type
	ok = true
	return
}

// Codec returns the AttributeCodec for the given registered type. nil is
// returned if the given type is not registered.
func (d *Dictionary) Codec(t byte) AttributeCodec {
	var entry *dictEntry = nil
	d.mu.RLock()
	if d.attributesByType[0] != nil {
		entry = d.attributesByType[0][t]
	}
	d.mu.RUnlock()
	if entry == nil {
		return AttributeUnknown
	}
	return entry.Codec
}
