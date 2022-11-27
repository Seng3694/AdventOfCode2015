package main

import (
	"aocutil"
	"math"
)

type Stats struct {
	hp, damage, armor int
}

type Equipment struct {
	cost, damage, armor int
}

func calculate_damage(damageStat, targetArmorStat int) int {
	result := damageStat - targetArmorStat
	if result <= 0 { // deals at least 1 damage
		return 1
	} else {
		return result
	}
}

func simulate_battle(player, boss *Stats) {
	turn := 0 //take turns attacking, player always starts first
	for player.hp > 0 && boss.hp > 0 {
		if turn%2 == 0 { // player attack
			boss.hp -= calculate_damage(player.damage, boss.armor)
		} else { // boss attack
			player.hp -= calculate_damage(boss.damage, player.armor)
		}
		turn++
	}
}

func main() {
	bossStats := Stats{
		hp:     109,
		damage: 8,
		armor:  2,
	}

	playerHp := 100

	weapons := []Equipment{
		{cost: 8, damage: 4, armor: 0},  // Dagger
		{cost: 10, damage: 5, armor: 0}, // Shortsword
		{cost: 25, damage: 6, armor: 0}, // Warhammer
		{cost: 40, damage: 7, armor: 0}, // Longsword
		{cost: 74, damage: 8, armor: 0}, // Greataxe
	}

	armor := []Equipment{
		{cost: 0, damage: 0, armor: 0},   // None
		{cost: 13, damage: 0, armor: 1},  // Leather
		{cost: 31, damage: 0, armor: 2},  // Chainmail
		{cost: 53, damage: 0, armor: 3},  // Splintmail
		{cost: 75, damage: 0, armor: 4},  // Bandedmail
		{cost: 102, damage: 0, armor: 5}, // Platemail
	}

	rings := []Equipment{
		{cost: 0, damage: 0, armor: 0},   // None 1
		{cost: 0, damage: 0, armor: 0},   // None 2
		{cost: 25, damage: 1, armor: 0},  // Damage +1
		{cost: 50, damage: 2, armor: 0},  // Damage +2
		{cost: 100, damage: 3, armor: 0}, // Damage +3
		{cost: 20, damage: 0, armor: 1},  // Defense +1
		{cost: 40, damage: 0, armor: 2},  // Defense +2
		{cost: 80, damage: 0, armor: 3},  // Defense +3
	}

	minCost := math.MaxInt
	maxCost := math.MinInt
	for _, w := range weapons {
		for _, a := range armor {
			for ri1, r1 := range rings {
				for ri2, r2 := range rings {
					if ri1 == ri2 { //can't buy the same ring twice
						continue
					}
					totalCost := w.cost + a.cost + r1.cost + r2.cost
					player := Stats{
						hp:     playerHp,
						damage: w.damage + a.damage + r1.damage + r2.damage,
						armor:  w.armor + a.armor + r1.armor + r2.armor,
					}
					boss := bossStats
					simulate_battle(&player, &boss)

					if player.hp > 0 && minCost > totalCost {
						minCost = totalCost
					} else if boss.hp > 0 && maxCost < totalCost {
						maxCost = totalCost
					}
				}
			}
		}
	}

	aocutil.AOCFinish(minCost, maxCost)
}
