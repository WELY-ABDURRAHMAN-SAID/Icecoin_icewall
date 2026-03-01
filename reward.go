package blockchain

func GetMiningReward(blockIndex int, minedAmount int, totalSupply int) int {
	reward := 50
	halvingInterval := 280000

	halvings := blockIndex / halvingInterval
	for i := 0; i < halvings; i++ {
		reward = reward / 2
	}

	if reward <= 0 {
		return 0
	}

	if minedAmount+reward > totalSupply {
		return totalSupply - minedAmount
	}

	return reward
}