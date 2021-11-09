package keeper_test

import (
	"time"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/cosmos/interchain-security/app"
	"github.com/cosmos/interchain-security/x/ccv/parent/types"
)

func (suite *KeeperTestSuite) TestParams() {
	expParams := types.DefaultParams()

	params := suite.parentChain.App.(*app.App).ParentKeeper.GetParams(suite.parentChain.GetContext())
	suite.Require().Equal(expParams, params)

	newParams := types.NewParams(ibctmtypes.NewClientState("", ibctmtypes.DefaultTrustLevel, 0, 0,
		time.Second*40, clienttypes.Height{}, commitmenttypes.GetSDKSpecs(), []string{"ibc", "upgradedIBCState"}, true, false))
	suite.parentChain.App.(*app.App).ParentKeeper.SetParams(suite.parentChain.GetContext(), newParams)
	params = suite.parentChain.App.(*app.App).ParentKeeper.GetParams(suite.parentChain.GetContext())
	suite.Require().Equal(newParams, params)
}
