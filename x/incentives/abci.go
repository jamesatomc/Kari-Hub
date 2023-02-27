package incentives

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	kariutils "hub/utils"

	"hub/x/incentives/keeper"
	"hub/x/incentives/types"
)

// BeginBlocker distributes block rewards to validators who have signed the
// previous block.
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	ids, totalBlockReward := k.ReleaseBlockReward(ctx, req.LastCommitInfo.Votes)

	if !totalBlockReward.IsZero() {
		k.Logger(ctx).Info(
			"released incentives",
			"ids", kariutils.UintArrayToString(ids, ","),
			"amount", totalBlockReward.String(),
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIncentivesReleased,
			sdk.NewAttribute(types.AttributeKeySchedules, kariutils.UintArrayToString(ids, ",")),
			sdk.NewAttribute(sdk.AttributeKeyAmount, totalBlockReward.String()),
		),
	)
}
