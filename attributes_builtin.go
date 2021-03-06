package radius

import (
	"encoding/binary"
	"errors"
	"net"
	"strconv"
	"time"
	"unicode/utf8"
)

// The base attribute value formats that are defined in RFC 2865.
var (
	// string
	AttributeText AttributeCodec
	// []byte
	AttributeString AttributeCodec
	// net.IP
	AttributeAddress AttributeCodec
	// uint32
	AttributeInteger AttributeCodec
	// uint64
	AttributeInteger64 AttributeCodec
	// time.Time
	AttributeTime AttributeCodec
	// []byte
	AttributeUnknown AttributeCodec
)

func init() {
	AttributeText = attributeText{}
	AttributeString = attributeString{}
	AttributeAddress = attributeAddress{}
	AttributeInteger = attributeInteger{}
	AttributeInteger64 = attributeInteger64{}
	AttributeTime = attributeTime{}
	AttributeUnknown = attributeString{}
}

type attributeText struct{}

func (attributeText) Decode(packet *Packet, value []byte) (interface{}, error) {
	if !utf8.Valid(value) {
		return nil, errors.New("radius: text attribute is not valid UTF-8")
	}
	return string(value), nil
}

func (attributeText) Encode(packet *Packet, value interface{}) ([]byte, error) {
	str, ok := value.(string)
	if ok {
		return []byte(str), nil
	}
	raw, ok := value.([]byte)
	if ok {
		return raw, nil
	}
	return nil, errors.New("radius: text attribute must be string or []byte")
}

type attributeString struct{}

// may be transform from-to base64 needed there
func (attributeString) Decode(packet *Packet, value []byte) (interface{}, error) {
	v := make([]byte, len(value))
	copy(v, value)
	return v, nil
}

func (attributeString) Encode(packet *Packet, value interface{}) ([]byte, error) {
	raw, ok := value.([]byte)
	if ok {
		return raw, nil
	}
	str, ok := value.(string)
	if ok {
		return []byte(str), nil
	}
	return nil, errors.New("radius: string attribute must be []byte or string")
}

type attributeAddress struct{}

func (attributeAddress) Decode(packet *Packet, value []byte) (interface{}, error) {
	if len(value) != net.IPv4len {
		return nil, errors.New("radius: address attribute has invalid size")
	}
	v := make([]byte, len(value))
	copy(v, value)
	return net.IP(v), nil
}

func (attributeAddress) Encode(packet *Packet, value interface{}) ([]byte, error) {
	ip, ok := value.(net.IP)
	if !ok {
		return nil, errors.New("radius: address attribute must be net.IP")
	}
	ip = ip.To4()
	if ip == nil {
		return nil, errors.New("radius: address attribute must be an IPv4 net.IP")
	}
	return []byte(ip), nil
}

func (attributeAddress) String(value interface{}) (interface{}, error) {
	if val, ok := value.(net.IP); ok {
		return val.String(), nil
	}
	return "", errors.New("radius: not net.IP")
}

func (attributeAddress) Transform(invalue interface{}) (interface{}, error) {
	if val, ok := invalue.(net.IP); ok {
		return val.To4(), nil
	}
	if str, ok := invalue.(string); ok {
		if ip := net.ParseIP(str).To4(); ip != nil {
			return ip, nil
		}
	}
	return nil, errors.New("radius: invalid IP format")
}

type attributeInteger struct{}

func (attributeInteger) Decode(packet *Packet, value []byte) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("radius: integer attribute has invalid size")
	}
	return binary.BigEndian.Uint32(value), nil
}

func (attributeInteger) Encode(packet *Packet, value interface{}) ([]byte, error) {
	integer, ok := value.(uint32)
	if !ok {
		return nil, errors.New("radius: integer attribute must be uint32")
	}
	raw := make([]byte, 4)
	binary.BigEndian.PutUint32(raw, integer)
	return raw, nil
}

func (attributeInteger) String(value interface{}) (interface{}, error) {
	if val, ok := value.(uint32); ok {
		return strconv.FormatUint(uint64(val), 10), nil
	}
	return "", errors.New("radius: not integer")
}

func (attributeInteger) Transform(invalue interface{}) (interface{}, error) {
	if val, ok := invalue.(uint32); ok {
		return val, nil
	}
	if _, ok := invalue.(string); ok {
		u, err := strconv.ParseUint(invalue.(string), 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(u), nil
	}
	return nil, errors.New("radius: invalid input type")
}

type attributeTime struct{}

func (attributeTime) Decode(packet *Packet, value []byte) (interface{}, error) {
	if len(value) != 4 {
		return nil, errors.New("radius: time attribute has invalid size")
	}
	return time.Unix(int64(binary.BigEndian.Uint32(value)), 0), nil
}

func (attributeTime) Encode(packet *Packet, value interface{}) ([]byte, error) {
	timestamp, ok := value.(time.Time)
	if !ok {
		return nil, errors.New("radius: time attribute must be time.Time")
	}
	raw := make([]byte, 4)
	binary.BigEndian.PutUint32(raw, uint32(timestamp.Unix()))
	return raw, nil
}

func (attributeTime) String(value interface{}) (interface{}, error) {
	if val, ok := value.(time.Time); ok {
		return strconv.FormatUint(uint64(val.Unix()), 10), nil
	}
	return "", errors.New("radius: not integer")
}

func (attributeTime) Transform(invalue interface{}) (interface{}, error) {
	if val, ok := invalue.(time.Time); ok {
		return val, nil
	}
	if _, ok := invalue.(string); ok {
		u, err := strconv.ParseInt(invalue.(string), 10, 32)
		if err != nil {
			return nil, err
		}
		t := time.Unix(u, 0)
		return t, nil
	}
	return nil, errors.New("radius: invalid input type")
}

type attributeInteger64 struct{}

func (attributeInteger64) Decode(packet *Packet, value []byte) (interface{}, error) {
	if len(value) != 8 {
		return nil, errors.New("radius: integer attribute has invalid size")
	}
	return binary.BigEndian.Uint64(value), nil
}

func (attributeInteger64) Encode(packet *Packet, value interface{}) ([]byte, error) {
	integer, ok := value.(uint64)
	if !ok {
		return nil, errors.New("radius: integer attribute must be uint32")
	}
	raw := make([]byte, 8)
	binary.BigEndian.PutUint64(raw, integer)
	return raw, nil
}

func (attributeInteger64) String(value interface{}) (interface{}, error) {
	if val, ok := value.(uint64); ok {
		return strconv.FormatUint(uint64(val), 10), nil
	}
	return "", errors.New("radius: not integer")
}

func (attributeInteger64) Transform(invalue interface{}) (interface{}, error) {
	if val, ok := invalue.(uint64); ok {
		return val, nil
	}
	if _, ok := invalue.(string); ok {
		u, err := strconv.ParseUint(invalue.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		return uint64(u), nil
	}
	return nil, errors.New("radius: invalid input type")
}
