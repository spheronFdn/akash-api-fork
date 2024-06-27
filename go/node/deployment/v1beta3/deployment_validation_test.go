package v1beta3_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	types "github.com/spheronFdn/akash-api-fork/go/node/deployment/v1beta3"
	akashtypes "github.com/spheronFdn/akash-api-fork/go/node/types/v1beta3"
	tutil "github.com/spheronFdn/akash-api-fork/go/testutil"
	testutil "github.com/spheronFdn/akash-api-fork/go/testutil/v1beta3"
)

const (
	regexInvalidUnitBoundaries = `^.*invalid unit count|CPU|GPU|memory|storage \(\d+ > 0 > \d+ fails\)$`
)

func TestZeroValueGroupSpec(t *testing.T) {
	did := testutil.DeploymentID(t)

	dgroup := testutil.DeploymentGroup(t, did, uint32(6))
	gspec := dgroup.GroupSpec

	t.Run("assert nominal test success", func(t *testing.T) {
		err := gspec.ValidateBasic()
		require.NoError(t, err)
	})
}

func TestZeroValueGroupSpecs(t *testing.T) {
	did := testutil.DeploymentID(t)
	dgroups := testutil.DeploymentGroups(t, did, uint32(6))
	gspecs := make([]types.GroupSpec, 0)
	for _, d := range dgroups {
		gspecs = append(gspecs, d.GroupSpec)
	}

	t.Run("assert nominal test success", func(t *testing.T) {
		err := types.ValidateDeploymentGroups(gspecs)
		require.NoError(t, err)
	})

	gspecZeroed := make([]types.GroupSpec, len(gspecs))
	gspecZeroed = append(gspecZeroed, gspecs...)
	t.Run("assert error for zero value bid duration", func(t *testing.T) {
		err := types.ValidateDeploymentGroups(gspecZeroed)
		require.Error(t, err)
	})
}

func TestEmptyGroupSpecIsInvalid(t *testing.T) {
	err := types.ValidateDeploymentGroups(make([]types.GroupSpec, 0))
	require.Equal(t, types.ErrInvalidGroups, err)
}

func validSimpleGroupSpec() types.GroupSpec {
	resources := make(types.ResourceUnits, 1)
	resources[0] = types.ResourceUnit{
		Resources: akashtypes.Resources{
			ID: 1,
			CPU: &akashtypes.CPU{
				Units: akashtypes.ResourceValue{
					Val: sdk.NewInt(10),
				},
				Attributes: nil,
			},
			GPU: &akashtypes.GPU{
				Units: akashtypes.ResourceValue{
					Val: sdk.NewInt(0),
				},
				Attributes: nil,
			},
			Memory: &akashtypes.Memory{
				Quantity: akashtypes.ResourceValue{
					Val: sdk.NewIntFromUint64(types.GetValidationConfig().Unit.Min.Memory),
				},
				Attributes: nil,
			},
			Storage: akashtypes.Volumes{
				akashtypes.Storage{
					Quantity: akashtypes.ResourceValue{
						Val: sdk.NewIntFromUint64(types.GetValidationConfig().Unit.Min.Storage),
					},
					Attributes: nil,
				},
			},
			Endpoints: akashtypes.Endpoints{},
		},
		Count: 1,
		Price: sdk.NewInt64DecCoin(tutil.CoinDenom, 1),
	}
	return types.GroupSpec{
		Name:         "testGroup",
		Requirements: akashtypes.PlacementRequirements{},
		Resources:    resources,
	}
}

func validSimpleGroupSpecs() []types.GroupSpec {
	result := make([]types.GroupSpec, 1)
	result[0] = validSimpleGroupSpec()

	return result
}

func TestSimpleGroupSpecIsValid(t *testing.T) {
	groups := validSimpleGroupSpecs()
	err := types.ValidateDeploymentGroups(groups)
	require.NoError(t, err)
}

func TestDuplicateSimpleGroupSpecIsInvalid(t *testing.T) {
	groups := validSimpleGroupSpecs()
	groupsDuplicate := make([]types.GroupSpec, 2)
	groupsDuplicate[0] = groups[0]
	groupsDuplicate[1] = groups[0]
	err := types.ValidateDeploymentGroups(groupsDuplicate)
	require.Error(t, err) // TODO - specific error
	require.Regexp(t, "^.*duplicate.*$", err)
}

func TestGroupWithZeroCount(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Count = 0
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithZeroCPU(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].CPU.Units.Val = sdk.NewInt(0)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithZeroMemory(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Memory.Quantity.Val = sdk.NewInt(0)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithZeroStorage(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Storage[0].Quantity.Val = sdk.NewInt(0)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, regexInvalidUnitBoundaries, err)
}

func TestGroupWithNilCPU(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].CPU = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit CPU.*$", err)
}

func TestGroupWithNilGPU(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].GPU = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit GPU.*$", err)
}

func TestGroupWithNilMemory(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Memory = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit memory.*$", err)
}

func TestGroupWithNilStorage(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Storage = nil
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid unit storage.*$", err)
}

func TestGroupWithInvalidPrice(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Price = sdk.DecCoin{}
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid price object.*$", err)
}

func TestGroupWithNegativePrice(t *testing.T) {
	group := validSimpleGroupSpec()
	group.Resources[0].Price.Amount = sdk.NewDec(-1)
	err := group.ValidateBasic()
	require.Error(t, err)
	require.Regexp(t, "^.*invalid price object.*$", err)
}
