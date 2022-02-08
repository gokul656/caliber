package keeper_test

import (
	"testing"

	testkeeper "github.com/malware656/caliber/testutil/keeper"
	"github.com/malware656/caliber/x/caliber/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CaliberKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
