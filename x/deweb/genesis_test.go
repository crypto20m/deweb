package deweb_test

import (
	"testing"

	keepertest "github.com/deweb-services/deweb/testutil/keeper"
	"github.com/deweb-services/deweb/x/deweb"
	"github.com/deweb-services/deweb/x/deweb/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DewebKeeper(t)
	deweb.InitGenesis(ctx, *k, genesisState)
	got := deweb.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
