package ethash

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// CLOMinerReward - Block reward in wei for successfully mining a block upward for Callisto Network
	CLOMinerReward = new(big.Int).Mul(big.NewInt(420), big.NewInt(1e+18))
	// CLOTreasuryReward - Block reward in wei for successfully mining a block upward for Callisto Network
	CLOTreasuryReward = new(big.Int).Mul(big.NewInt(120), big.NewInt(1e+18))
	// CLOStakeReward - Block reward in wei for successfully mining a block upward for Callisto Network
	CLOStakeReward = new(big.Int).Mul(big.NewInt(60), big.NewInt(1e+18))
	// CLOHF1TreasuryReward - Block reward in wei for successfully mining a block upward for Callisto Network
	CLOHF1TreasuryReward = new(big.Int).Mul(big.NewInt(60), big.NewInt(1e+18))
	// CLOHF1StakeReward - Block reward in wei for successfully mining a block upward for Callisto Network
	CLOHF1StakeReward = new(big.Int).Mul(big.NewInt(120), big.NewInt(1e+18))
	// CLOTreasuryAddress - Treasury Address
	CLOTreasuryAddress = common.HexToAddress("0x74682Fc32007aF0b6118F259cBe7bCCC21641600")
	// CLOStakeAddress - Stake Address before HF1
	CLOStakeAddress = common.HexToAddress("0x3c06f218Ce6dD8E2c535a8925A2eDF81674984D9")
	// CLOHF1StakeAddress - Stake Address HF1
	CLOHF1StakeAddress = common.HexToAddress("0xd813419749b3c2cdc94a2f9cfcf154113264a9d6")
	// CLOMonetaryPolicyMinerReward - Monetary Policy blocks and miner rewards
	CLOMonetaryPolicyMinerReward = map[*big.Int]*big.Int{
		big.NewInt(2750001): calcBigNumber(234),
		big.NewInt(4250001): calcBigNumber(129.6),
		big.NewInt(5750001): calcBigNumber(71.28),
		big.NewInt(7250001): calcBigNumber(38.88),
	}
	// CLOMonetaryPolicyTreasury - Monetary Policy blocks and treasury rewards
	CLOMonetaryPolicyTreasury = map[*big.Int]*big.Int{
		big.NewInt(2750001): calcBigNumber(36),
		big.NewInt(4250001): calcBigNumber(21.6),
		big.NewInt(5750001): calcBigNumber(12.96),
		big.NewInt(7250001): calcBigNumber(7.77),
	}
	// CLOMonetaryPolicyStake - Monetary Policy blocks and stake rewards
	CLOMonetaryPolicyStake = map[*big.Int]*big.Int{
		big.NewInt(2750001): calcBigNumber(90),
		big.NewInt(4250001): calcBigNumber(64.8),
		big.NewInt(5750001): calcBigNumber(45.36),
		big.NewInt(7250001): calcBigNumber(31.1),
	}
)

func calcBigNumber(reward float64) *big.Int {
	bigReward := new(big.Float).Mul(big.NewFloat(reward), big.NewFloat(1e+18))
	bigRewardInt := new(big.Int)
	bigReward.Int(bigRewardInt)
	return bigRewardInt
}
