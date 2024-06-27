package v1beta2

import (
	"context"

	"google.golang.org/grpc"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evdtypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staketypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	atypes "github.com/spheronFdn/akash-api-fork/go/node/audit/v1beta3"
	ctypes "github.com/spheronFdn/akash-api-fork/go/node/cert/v1beta3"
	dtypes "github.com/spheronFdn/akash-api-fork/go/node/deployment/v1beta3"
	mtypes "github.com/spheronFdn/akash-api-fork/go/node/market/v1beta4"
	ptypes "github.com/spheronFdn/akash-api-fork/go/node/provider/v1beta3"
)

type sdkQueryClient struct {
	auth     authtypes.QueryClient
	authz    authz.QueryClient
	bank     banktypes.QueryClient
	distr    disttypes.QueryClient
	evidence evdtypes.QueryClient
	feegrant feegranttypes.QueryClient
	gov      govtypes.QueryClient
	mint     minttypes.QueryClient
	params   paramtypes.QueryClient
	slashing slashtypes.QueryClient
	staking  staketypes.QueryClient
	upgrade  upgradetypes.QueryClient
}

type queryClient struct {
	dclient dtypes.QueryClient
	mclient mtypes.QueryClient
	pclient ptypes.QueryClient
	aclient atypes.QueryClient
	cclient ctypes.QueryClient
	sdk     sdkQueryClient
	cctx    sdkclient.Context
}

// NewQueryClient creates new query client instance
func NewQueryClient(cctx sdkclient.Context) QueryClient {
	return newQueryClient(cctx)
}

func newQueryClient(cctx sdkclient.Context) *queryClient {
	return &queryClient{
		dclient: dtypes.NewQueryClient(cctx),
		mclient: mtypes.NewQueryClient(cctx),
		pclient: ptypes.NewQueryClient(cctx),
		aclient: atypes.NewQueryClient(cctx),
		cclient: ctypes.NewQueryClient(cctx),
		sdk: sdkQueryClient{
			auth:     authtypes.NewQueryClient(cctx),
			authz:    authz.NewQueryClient(cctx),
			bank:     banktypes.NewQueryClient(cctx),
			distr:    disttypes.NewQueryClient(cctx),
			evidence: evdtypes.NewQueryClient(cctx),
			feegrant: feegranttypes.NewQueryClient(cctx),
			gov:      govtypes.NewQueryClient(cctx),
			mint:     minttypes.NewQueryClient(cctx),
			params:   paramtypes.NewQueryClient(cctx),
			slashing: slashtypes.NewQueryClient(cctx),
			staking:  staketypes.NewQueryClient(cctx),
			upgrade:  upgradetypes.NewQueryClient(cctx),
		},
		cctx: cctx,
	}
}

func (c *queryClient) ClientContext() sdkclient.Context {
	return c.cctx
}

func (c *queryClient) Deployments(ctx context.Context, in *dtypes.QueryDeploymentsRequest, opts ...grpc.CallOption) (*dtypes.QueryDeploymentsResponse, error) {
	if c.dclient == nil {
		return &dtypes.QueryDeploymentsResponse{}, ErrClientNotFound
	}
	return c.dclient.Deployments(ctx, in, opts...)
}

func (c *queryClient) Deployment(ctx context.Context, in *dtypes.QueryDeploymentRequest, opts ...grpc.CallOption) (*dtypes.QueryDeploymentResponse, error) {
	if c.dclient == nil {
		return &dtypes.QueryDeploymentResponse{}, ErrClientNotFound
	}
	return c.dclient.Deployment(ctx, in, opts...)
}

func (c *queryClient) Group(ctx context.Context, in *dtypes.QueryGroupRequest, opts ...grpc.CallOption) (*dtypes.QueryGroupResponse, error) {
	if c.dclient == nil {
		return &dtypes.QueryGroupResponse{}, ErrClientNotFound
	}
	return c.dclient.Group(ctx, in, opts...)
}

func (c *queryClient) Orders(ctx context.Context, in *mtypes.QueryOrdersRequest, opts ...grpc.CallOption) (*mtypes.QueryOrdersResponse, error) {
	if c.mclient == nil {
		return &mtypes.QueryOrdersResponse{}, ErrClientNotFound
	}
	return c.mclient.Orders(ctx, in, opts...)
}

func (c *queryClient) Order(ctx context.Context, in *mtypes.QueryOrderRequest, opts ...grpc.CallOption) (*mtypes.QueryOrderResponse, error) {
	if c.mclient == nil {
		return &mtypes.QueryOrderResponse{}, ErrClientNotFound
	}
	return c.mclient.Order(ctx, in, opts...)
}

func (c *queryClient) Bids(ctx context.Context, in *mtypes.QueryBidsRequest, opts ...grpc.CallOption) (*mtypes.QueryBidsResponse, error) {
	if c.mclient == nil {
		return &mtypes.QueryBidsResponse{}, ErrClientNotFound
	}
	return c.mclient.Bids(ctx, in, opts...)
}

func (c *queryClient) Bid(ctx context.Context, in *mtypes.QueryBidRequest, opts ...grpc.CallOption) (*mtypes.QueryBidResponse, error) {
	if c.mclient == nil {
		return &mtypes.QueryBidResponse{}, ErrClientNotFound
	}
	return c.mclient.Bid(ctx, in, opts...)
}

func (c *queryClient) Leases(ctx context.Context, in *mtypes.QueryLeasesRequest, opts ...grpc.CallOption) (*mtypes.QueryLeasesResponse, error) {
	if c.mclient == nil {
		return &mtypes.QueryLeasesResponse{}, ErrClientNotFound
	}
	return c.mclient.Leases(ctx, in, opts...)
}

func (c *queryClient) Lease(ctx context.Context, in *mtypes.QueryLeaseRequest, opts ...grpc.CallOption) (*mtypes.QueryLeaseResponse, error) {
	if c.mclient == nil {
		return &mtypes.QueryLeaseResponse{}, ErrClientNotFound
	}
	return c.mclient.Lease(ctx, in, opts...)
}

func (c *queryClient) Providers(ctx context.Context, in *ptypes.QueryProvidersRequest, opts ...grpc.CallOption) (*ptypes.QueryProvidersResponse, error) {
	if c.pclient == nil {
		return &ptypes.QueryProvidersResponse{}, ErrClientNotFound
	}
	return c.pclient.Providers(ctx, in, opts...)
}

func (c *queryClient) Provider(ctx context.Context, in *ptypes.QueryProviderRequest, opts ...grpc.CallOption) (*ptypes.QueryProviderResponse, error) {
	if c.pclient == nil {
		return &ptypes.QueryProviderResponse{}, ErrClientNotFound
	}
	return c.pclient.Provider(ctx, in, opts...)
}

// AllProvidersAttributes queries all providers
func (c *queryClient) AllProvidersAttributes(ctx context.Context, in *atypes.QueryAllProvidersAttributesRequest, opts ...grpc.CallOption) (*atypes.QueryProvidersResponse, error) {
	if c.pclient == nil {
		return &atypes.QueryProvidersResponse{}, ErrClientNotFound
	}
	return c.aclient.AllProvidersAttributes(ctx, in, opts...)
}

// ProviderAttributes queries all provider signed attributes
func (c *queryClient) ProviderAttributes(ctx context.Context, in *atypes.QueryProviderAttributesRequest, opts ...grpc.CallOption) (*atypes.QueryProvidersResponse, error) {
	if c.pclient == nil {
		return &atypes.QueryProvidersResponse{}, ErrClientNotFound
	}
	return c.aclient.ProviderAttributes(ctx, in, opts...)
}

// ProviderAuditorAttributes queries provider signed attributes by specific validator
func (c *queryClient) ProviderAuditorAttributes(ctx context.Context, in *atypes.QueryProviderAuditorRequest, opts ...grpc.CallOption) (*atypes.QueryProvidersResponse, error) {
	if c.pclient == nil {
		return &atypes.QueryProvidersResponse{}, ErrClientNotFound
	}
	return c.aclient.ProviderAuditorAttributes(ctx, in, opts...)
}

// AuditorAttributes queries all providers signed by this validator
func (c *queryClient) AuditorAttributes(ctx context.Context, in *atypes.QueryAuditorAttributesRequest, opts ...grpc.CallOption) (*atypes.QueryProvidersResponse, error) {
	if c.aclient == nil {
		return &atypes.QueryProvidersResponse{}, ErrClientNotFound
	}
	return c.aclient.AuditorAttributes(ctx, in, opts...)
}

func (c *queryClient) Certificates(ctx context.Context, in *ctypes.QueryCertificatesRequest, opts ...grpc.CallOption) (*ctypes.QueryCertificatesResponse, error) {
	if c.cclient == nil {
		return &ctypes.QueryCertificatesResponse{}, ErrClientNotFound
	}
	return c.cclient.Certificates(ctx, in, opts...)
}

func (c *queryClient) Auth() authtypes.QueryClient {
	return c.sdk.auth
}

func (c *queryClient) Authz() authz.QueryClient {
	return c.sdk.authz
}

func (c *queryClient) Bank() banktypes.QueryClient {
	return c.sdk.bank
}

func (c *queryClient) Distribution() disttypes.QueryClient {
	return c.sdk.distr
}

func (c *queryClient) Evidence() evdtypes.QueryClient {
	return c.sdk.evidence
}

func (c *queryClient) Feegrant() feegranttypes.QueryClient {
	return c.sdk.feegrant
}

func (c *queryClient) Gov() govtypes.QueryClient {
	return c.sdk.gov
}

func (c *queryClient) Mint() minttypes.QueryClient {
	return c.sdk.mint
}

func (c *queryClient) Params() paramtypes.QueryClient {
	return c.sdk.params
}

func (c *queryClient) Slashing() slashtypes.QueryClient {
	return c.sdk.slashing
}

func (c *queryClient) Staking() staketypes.QueryClient {
	return c.sdk.staking
}

func (c *queryClient) Upgrade() upgradetypes.QueryClient {
	return c.sdk.upgrade
}
