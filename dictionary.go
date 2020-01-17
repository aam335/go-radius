package radius

import (
	"fmt"
	"sync"
)

var builtinOnce sync.Once

// NotVSID VSA ID for non-VSA (RFC) attributes
const NotVSID = 0

// Builtin is the built-in dictionary. It is initially loaded with the
// attributes defined in RFC 2865 and RFC 2866.
var Builtin *Dictionary

func initDictionary() {
	Builtin = &Dictionary{}
}

type dictEntry struct {
	Vendor uint32
	Type   byte
	Tagged bool
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

// VsaRegisterTagFlag registers the AttributeCodec for the given attribute name and type, value with or without tag attr.
func (d *Dictionary) VsaRegisterTagFlag(vendorID uint32, name string, t byte, hasTag bool, codec AttributeCodec) error {
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
		Tagged: hasTag,
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

// VsaRegister registers the AttributeCodec for the given attribute name and type, without tag attr.
func (d *Dictionary) VsaRegister(vendorID uint32, name string, t byte, codec AttributeCodec) error {
	return d.VsaRegisterTagFlag(vendorID, name, t, false, codec)
}

// VsaRegisterTag registers the AttributeCodec for the given attribute name and type, without tag attr.
func (d *Dictionary) VsaRegisterTag(vendorID uint32, name string, t byte, codec AttributeCodec) error {
	return d.VsaRegisterTagFlag(vendorID, name, t, true, codec)
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

// VsaMustRegisterTag is a helper for Register that panics if it returns an error.
func (d *Dictionary) VsaMustRegisterTag(vendorID uint32, name string, t byte, codec AttributeCodec) {
	if err := d.VsaRegisterTagFlag(vendorID, name, t, true, codec); err != nil {
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
// name. For tagged attributes use ValueTagged type w/same value types as untagged
//
// If name is not registered, nil and an error is returned.
//
// If the attribute's codec implements AttributeTransformer, the value is
// first transformed before being stored in *Attribute. If the transform
// function returns an error, nil and the error is returned.
func (d *Dictionary) Attr(name string, value interface{}) (*Attribute, error) {
	if tagVal, ok := value.(ValueTagged); ok {
		return d.coreAttrTag(name, true, tagVal.Tag, tagVal.Value)
	}
	return d.coreAttrTag(name, false, 0, value)
}

// AttrTagged returns a new *Attribute whose type is registered under the given
// name.
//
// If name is not registered, nil and an error is returned.
//
// If the attribute's codec implements AttributeTransformer, the value is
// first transformed before being stored in *Attribute. If the transform
// function returns an error, nil and the error is returned.
func (d *Dictionary) AttrTagged(name string, tag byte, value interface{}) (*Attribute, error) {
	return d.coreAttrTag(name, true, tag, value)
}

func (d *Dictionary) coreAttrTag(name string, tagged bool, tag byte, value interface{}) (*Attribute, error) {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil {
		return nil, fmt.Errorf("radius: attribute name '%v' not registered", name)
	}
	if entry.Tagged != tagged {
		return nil, fmt.Errorf("Attribute %v has Tagged=%v, but sets as Tagged=%v", name, entry.Tagged, tagged)
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
		Tagged: tagged,
		Tag:    tag,
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
	return d.NameVID(0, t)
}

// Type returns the registered type for the given attribute name. ok is false
// if the given name is not registered.
func (d *Dictionary) Type(name string) (t byte, ok bool) {
	d.mu.RLock()
	entry := d.attributesByName[name]
	d.mu.RUnlock()
	if entry == nil || entry.Vendor != NotVSID {
		return
	}
	t = entry.Type
	ok = true
	return
}

// Codec returns the AttributeCodec for the given registered type. nil is
// returned if the given type is not registered.
func (d *Dictionary) Codec(t byte) AttributeCodec {
	entry := d.getDictEntry(NotVSID, t)
	if entry == nil {
		return nil
	}
	return entry.Codec
}

// NameVID returns the registered name for the given attribute VID and type. ok is false
// if the given type is not registered.
func (d *Dictionary) NameVID(vendorID uint32, t byte) (name string, ok bool) {
	entry := d.getDictEntry(vendorID, t)
	if entry == nil {
		return
	}
	name = entry.Name
	ok = true
	return
}

// CodecVID returns the AttributeCodec for the given registered type. nil is
// returned if the given type is not registered.
func (d *Dictionary) CodecVID(vendorID uint32, t byte) AttributeCodec {
	entry := d.getDictEntry(vendorID, t)
	if entry == nil {
		return nil
	}
	return entry.Codec
}

// getDictEntry ...
func (d *Dictionary) getDictEntry(vendorID uint32, t byte) (entry *dictEntry) {
	//	var entry *dictEntry = nil
	d.mu.RLock()
	if d.attributesByType[vendorID] != nil {
		entry = d.attributesByType[vendorID][t]
	}
	d.mu.RUnlock()
	return
}

// getDictEntryByName ...
func (d *Dictionary) getDictEntryByName(name string) (entry *dictEntry) {
	//	var entry *dictEntry = nil
	d.mu.RLock()
	entry = d.attributesByName[name]
	d.mu.RUnlock()
	return
}
