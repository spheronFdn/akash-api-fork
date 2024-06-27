package v1beta3

import (
	etypes "github.com/spheronFdn/akash-api-fork/go/node/escrow/v1beta3"
)

const (
	EscrowScope = "deployment"
)

func EscrowAccountForDeployment(id DeploymentID) etypes.AccountID {
	return etypes.AccountID{
		Scope: EscrowScope,
		XID:   id.String(),
	}
}

func DeploymentIDFromEscrowAccount(id etypes.AccountID) (DeploymentID, bool) {
	if id.Scope != EscrowScope {
		return DeploymentID{}, false
	}

	did, err := ParseDeploymentID(id.XID)
	return did, err == nil
}
