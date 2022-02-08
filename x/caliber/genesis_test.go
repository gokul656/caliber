package caliber_test

import (
	"testing"

	keepertest "github.com/malware656/caliber/testutil/keeper"
	"github.com/malware656/caliber/testutil/nullify"
	"github.com/malware656/caliber/x/caliber"
	"github.com/malware656/caliber/x/caliber/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CaliberKeeper(t)
	caliber.InitGenesis(ctx, *k, genesisState)
	got := caliber.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
