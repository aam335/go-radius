package radius

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type td struct{}

func (td) Dict() (uint32, []DictionaryAttr) {
	return 0, []DictionaryAttr{
		{1, "test", false, AttributeString},
		{2, "test1", false, AttributeString},
		{3, "test2", false, AttributeString},
	}
}

func TestRegisterN(t *testing.T) {
	d := Dictionary{}
	assert.NoError(t, d.RegisterN(td{}), "Normal Attrs")
	_, ok := d.Name(1)
	require.True(t, ok, "unable to find attr Type=1")
}
