package vendor

import (
	"testing"

	"github.com/stretchr/testify/require"

	r "github.com/aam335/go-radius"
)

func TestVendor(t *testing.T) {
	d := r.Dictionary{}
	d.MustRegisterDC(Redback)
	_, ok := d.NameVID(VendorIDRedback, 1)
	require.Truef(t, ok, "unable to find attr Vendor: %v Type %v", VendorIDRedback, 1)
}
