package main

import (
	"github.com/orolol/dkp/utils"
)

func genericApplyEffect(player *utils.PlayerInGame, opponent *utils.PlayerInGame, effects []utils.Effect, game *utils.Game) {
	for _, e := range effects {
		if e.Target == "Player" || e.Target == "Both" || e.Target == "" {
			utils.ApplyEffect(player, e, game)
		}
		if e.Target == "Opponent" || e.Target == "Both" {
			utils.ApplyEffect(opponent, e, game)
		}
	}
}

func genericApplyCosts(player *utils.PlayerInGame, costs []utils.Cost) {
	for _, c := range costs {
		utils.ApplyCost(player, c)
	}
}
