package radius

import (
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {
	attr := attributeInteger{}
	expect := uint32(100)
	x, err := attr.Encode(nil, expect)
	assert.NoError(t, err)
	v, err := attr.Decode(nil, x)
	assert.NoError(t, err)
	assert.Equal(t, expect, v)

	str := strconv.FormatUint(uint64(expect), 10)
	tr, err := attr.Transform(str)
	assert.NoError(t, err)
	assert.Equal(t, expect, v)
	v, err = attr.String(tr)
	assert.NoError(t, err)
	assert.Equal(t, str, v)
}

func TestIp4(t *testing.T) {
	attr := attributeAddress{}
	strexpect := "127.0.0.1"
	expect := net.ParseIP(strexpect).To4()
	x, err := attr.Encode(nil, expect)
	assert.NoError(t, err)
	v, err := attr.Decode(nil, x)
	assert.NoError(t, err)
	assert.Equal(t, expect, v)

	tr, err := attr.Transform(strexpect)
	assert.NoError(t, err)
	assert.Equal(t, expect, v)
	v, err = attr.String(tr)
	assert.NoError(t, err)
	assert.Equal(t, strexpect, v)
}

func TestTime(t *testing.T) {
	attr := attributeTime{}
	exptimeUnix := time.Now().Unix()
	exptime := time.Unix(exptimeUnix, 0)

	xx, _ := exptime.MarshalText()
	log.Print(string(xx))
	nt := time.Time{}
	err := nt.UnmarshalText(xx)
	log.Print(exptime, nt)

	expectu32 := uint32(exptime.Unix())
	strexpect := strconv.FormatUint(uint64(expectu32), 10)

	x, err := attr.Encode(nil, exptime)
	assert.NoError(t, err)
	v, err := attr.Decode(nil, x)
	assert.NoError(t, err)
	assert.Equal(t, exptime, v)

	tr, err := attr.Transform(strexpect)
	assert.NoError(t, err)
	assert.Equal(t, exptime, v)
	v, err = attr.String(tr)
	assert.NoError(t, err)
	assert.Equal(t, strexpect, v)
}
