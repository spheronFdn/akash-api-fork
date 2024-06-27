package testutil

import (
	"testing"

	ptypes "github.com/spheronFdn/akash-api-fork/go/node/provider/v1beta3"
	"github.com/spheronFdn/akash-api-fork/go/testutil"
)

func Provider(t testing.TB) ptypes.Provider {
	t.Helper()

	return ptypes.Provider{
		Owner:      AccAddress(t).String(),
		HostURI:    testutil.Hostname(t),
		Attributes: Attributes(t),
		Info: ptypes.ProviderInfo{
			EMail:   "test@example.com",
			Website: ProviderHostname(t),
		},
	}
}
