package main

import (
	"fmt"
	"math"
)

/*
--- Day 22: Wizard Simulator 20XX ---

Little Henry Case decides that defeating bosses with swords and stuff is boring. Now he's playing the game with a wizard. Of course, he gets stuck on another boss and needs your help again.

In this version, combat still proceeds with the player and the boss taking alternating turns. The player still goes first. Now, however, you don't get any equipment; instead, you must choose one of your spells to cast. The first character at or below 0 hit points loses.

Since you're a wizard, you don't get to wear armor, and you can't attack normally. However, since you do magic damage, your opponent's armor is ignored, and so the boss effectively has zero armor as well. As before, if armor (from a spell, in this case) would reduce damage below 1, it becomes 1 instead - that is, the boss' attacks always deal at least 1 damage.

On each of your turns, you must select one of your spells to cast. If you cannot afford to cast any spell, you lose. Spells cost mana; you start with 500 mana, but have no maximum limit. You must have enough mana to cast a spell, and its cost is immediately deducted when you cast it. Your spells are Magic Missile, Drain, Shield, Poison, and Recharge.

    Magic Missile costs 53 mana. It instantly does 4 damage.
    Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
    Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
    Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
    Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.

Effects all work the same way. Effects apply at the start of both the player's turns and the boss' turns. Effects are created with a timer (the number of turns they last); at the start of each turn, after they apply any effect they have, their timer is decreased by one. If this decreases the timer to zero, the effect ends. You cannot cast a spell that would start an effect which is already active. However, effects can be started on the same turn they end.

For example, suppose the player has 10 hit points and 250 mana, and that the boss has 13 hit points and 8 damage:

-- Player turn --
- Player has 10 hit points, 0 armor, 250 mana
- Boss has 13 hit points
Player casts Poison.

-- Boss turn --
- Player has 10 hit points, 0 armor, 77 mana
- Boss has 13 hit points
Poison deals 3 damage; its timer is now 5.
Boss attacks for 8 damage.

-- Player turn --
- Player has 2 hit points, 0 armor, 77 mana
- Boss has 10 hit points
Poison deals 3 damage; its timer is now 4.
Player casts Magic Missile, dealing 4 damage.

-- Boss turn --
- Player has 2 hit points, 0 armor, 24 mana
- Boss has 3 hit points
Poison deals 3 damage. This kills the boss, and the player wins.

Now, suppose the same initial conditions, except that the boss has 14 hit points instead:

-- Player turn --
- Player has 10 hit points, 0 armor, 250 mana
- Boss has 14 hit points
Player casts Recharge.

-- Boss turn --
- Player has 10 hit points, 0 armor, 21 mana
- Boss has 14 hit points
Recharge provides 101 mana; its timer is now 4.
Boss attacks for 8 damage!

-- Player turn --
- Player has 2 hit points, 0 armor, 122 mana
- Boss has 14 hit points
Recharge provides 101 mana; its timer is now 3.
Player casts Shield, increasing armor by 7.

-- Boss turn --
- Player has 2 hit points, 7 armor, 110 mana
- Boss has 14 hit points
Shield's timer is now 5.
Recharge provides 101 mana; its timer is now 2.
Boss attacks for 8 - 7 = 1 damage!

-- Player turn --
- Player has 1 hit point, 7 armor, 211 mana
- Boss has 14 hit points
Shield's timer is now 4.
Recharge provides 101 mana; its timer is now 1.
Player casts Drain, dealing 2 damage, and healing 2 hit points.

-- Boss turn --
- Player has 3 hit points, 7 armor, 239 mana
- Boss has 12 hit points
Shield's timer is now 3.
Recharge provides 101 mana; its timer is now 0.
Recharge wears off.
Boss attacks for 8 - 7 = 1 damage!

-- Player turn --
- Player has 2 hit points, 7 armor, 340 mana
- Boss has 12 hit points
Shield's timer is now 2.
Player casts Poison.

-- Boss turn --
- Player has 2 hit points, 7 armor, 167 mana
- Boss has 12 hit points
Shield's timer is now 1.
Poison deals 3 damage; its timer is now 5.
Boss attacks for 8 - 7 = 1 damage!

-- Player turn --
- Player has 1 hit point, 7 armor, 167 mana
- Boss has 9 hit points
Shield's timer is now 0.
Shield wears off, decreasing armor by 7.
Poison deals 3 damage; its timer is now 4.
Player casts Magic Missile, dealing 4 damage.

-- Boss turn --
- Player has 1 hit point, 0 armor, 114 mana
- Boss has 2 hit points
Poison deals 3 damage. This kills the boss, and the player wins.

You start with 50 hit points and 500 mana points. The boss's actual stats are in your puzzle input. What is the least amount of mana you can spend and still win the fight? (Do not include mana recharge effects as "spending" negative mana.)

--- Part Two ---

On the next run through the game, you increase the difficulty to hard.

At the start of each player turn (before any other effects apply), you lose 1 hit point. If this brings you to or below 0 hit points, you lose.

With the same starting stats for you and the boss, what is the least amount of mana you can spend and still win the fight?

*/

const (
	BATTLE_STATE_ONGOING int = iota
	BATTLE_STATE_WON
	BATTLE_STATE_LOST
)

type Day22BossStats struct {
	hp, damage  int
	poisonTimer int
}

type Day22PlayerStats struct {
	maxHp, hp, mana, armor     int
	shieldTimer, rechargeTimer int
}

const (
	SPELL_MAGIC_MISSILE int = iota
	SPELL_DRAIN
	SPELL_SHIELD
	SPELL_POISON
	SPELL_RECHARGE
	SPELL_COUNT
)

type Day22Spell struct {
	cost, timer, value int
}

var (
	spells = []Day22Spell{
		{cost: 53, timer: 0, value: 4},    // Magic Missile
		{cost: 73, timer: 0, value: 2},    // Drain
		{cost: 113, timer: 6, value: 7},   // Shield
		{cost: 173, timer: 6, value: 3},   // Poison
		{cost: 229, timer: 5, value: 101}, // Recharge
	}
)

func day22_clamp(value, min, max int) int {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func day22_can_apply_spell(id int, player *Day22PlayerStats, boss *Day22BossStats) bool {
	manaNextTurn := player.mana
	if player.rechargeTimer > 0 {
		manaNextTurn += spells[SPELL_RECHARGE].value
	}
	if manaNextTurn < spells[id].cost {
		return false
	}
	switch id {
	case SPELL_MAGIC_MISSILE:
		return true
	case SPELL_DRAIN:
		return true
	case SPELL_SHIELD: //for timers -1 because there will be one more tick
		if player.shieldTimer-1 > 0 {
			return false
		}
		return true
	case SPELL_POISON:
		if boss.poisonTimer-1 > 0 {
			return false
		}
		return true
	case SPELL_RECHARGE:
		if player.rechargeTimer-1 > 0 {
			return false
		}
		return true
	}
	return true
}

func day22_cast_spell(id int, player *Day22PlayerStats, boss *Day22BossStats) {
	switch id {
	case SPELL_MAGIC_MISSILE:
		boss.hp -= spells[id].value
	case SPELL_DRAIN: //not sure if there is maxhp. "heal" implies it
		player.hp = day22_clamp(player.hp+spells[id].value, 0, player.maxHp)
		boss.hp -= spells[id].value
	case SPELL_SHIELD:
		player.shieldTimer = spells[id].timer
		player.armor += spells[id].value
	case SPELL_POISON:
		boss.poisonTimer = spells[id].timer
	case SPELL_RECHARGE:
		player.rechargeTimer = spells[id].timer
	}
	player.mana -= spells[id].cost
}

func day22_update_effects(player *Day22PlayerStats, boss *Day22BossStats) {
	player.shieldTimer = day22_clamp(player.shieldTimer-1, -1, math.MaxInt)
	player.rechargeTimer = day22_clamp(player.rechargeTimer-1, -1, math.MaxInt)
	boss.poisonTimer = day22_clamp(boss.poisonTimer-1, -1, math.MaxInt)

	if player.shieldTimer == 0 { //if it just ran out then reduce armor
		player.armor -= spells[SPELL_SHIELD].value
		//set to -1 so it won't reduce armor next turn
		player.shieldTimer = -1
	}
	if player.rechargeTimer > -1 {
		player.mana += spells[SPELL_RECHARGE].value
	}
	if boss.poisonTimer > -1 {
		boss.hp -= spells[SPELL_POISON].value
	}
}

func day22_set_battle_state(player *Day22PlayerStats, boss *Day22BossStats, ignoreMana bool) int {
	if boss.hp <= 0 {
		return BATTLE_STATE_WON
	} else if player.hp <= 0 || (!ignoreMana && player.mana < spells[SPELL_MAGIC_MISSILE].cost) {
		return BATTLE_STATE_LOST
	} else {
		return BATTLE_STATE_ONGOING
	}
}

type Day22BfsData struct {
	spellID, manaSpent, battleState int
	player                          Day22PlayerStats
	boss                            Day22BossStats
}

func day22_player_turn(data *Day22BfsData, isHardmode bool) int {
	if isHardmode {
		data.player.hp--
		data.battleState = day22_set_battle_state(&data.player, &data.boss, true)
		if data.battleState == BATTLE_STATE_LOST {
			return BATTLE_STATE_LOST
		}
	}

	day22_update_effects(&data.player, &data.boss)
	//check for battlestate. boss could die from poison or player has no more mana
	data.battleState = day22_set_battle_state(&data.player, &data.boss, false)
	if data.battleState != BATTLE_STATE_ONGOING {
		return data.battleState
	}

	day22_cast_spell(data.spellID, &data.player, &data.boss)
	data.manaSpent += spells[data.spellID].cost

	data.battleState = day22_set_battle_state(&data.player, &data.boss, true)
	if data.battleState == BATTLE_STATE_WON {
		return BATTLE_STATE_WON
	}

	return BATTLE_STATE_ONGOING
}

func day22_boss_turn(data *Day22BfsData) int {
	day22_update_effects(&data.player, &data.boss)
	//check for battlestate. boss could die from poison
	data.battleState = day22_set_battle_state(&data.player, &data.boss, true)
	if data.battleState == BATTLE_STATE_WON {
		return BATTLE_STATE_WON
	}

	damage := day22_clamp(data.boss.damage-data.player.armor, 1, data.player.hp)
	data.player.hp -= damage
	data.battleState = day22_set_battle_state(&data.player, &data.boss, true)
	if data.battleState == BATTLE_STATE_LOST {
		return BATTLE_STATE_LOST
	}

	return BATTLE_STATE_ONGOING
}

func day22_solution(isHardmode bool) int {
	bossStats := Day22BossStats{
		hp:          55,
		damage:      8,
		poisonTimer: -1,
	}

	playerStats := Day22PlayerStats{
		maxHp:         50,
		hp:            50,
		mana:          500,
		armor:         0,
		shieldTimer:   -1,
		rechargeTimer: -1,
	}
	minManaSpent := math.MaxInt

	for k := 0; k < SPELL_COUNT; k++ {
		bfsSpellQueue := make([]Day22BfsData, 0, 64)
		//starting spell. build bfs tree from here
		bfsSpellQueue = append(bfsSpellQueue, Day22BfsData{
			spellID:     k,
			battleState: BATTLE_STATE_ONGOING,
			manaSpent:   0,
			player:      playerStats,
			boss:        bossStats})

		for len(bfsSpellQueue) > 0 {
			queueLength := len(bfsSpellQueue)

			for s := 0; s < queueLength; s++ {
				item := bfsSpellQueue[0]          //peek
				bfsSpellQueue = bfsSpellQueue[1:] //pop

				state := day22_player_turn(&item, isHardmode)
				if item.manaSpent > minManaSpent || state == BATTLE_STATE_LOST {
					continue
				} else if state == BATTLE_STATE_WON {
					if item.manaSpent < minManaSpent {
						minManaSpent = item.manaSpent
					}
					continue
				}

				state = day22_boss_turn(&item)
				if state == BATTLE_STATE_WON {
					if item.manaSpent < minManaSpent {
						minManaSpent = item.manaSpent
					}
					continue
				} else if state == BATTLE_STATE_LOST {
					continue
				}

				if item.battleState == BATTLE_STATE_ONGOING {
					for i := range spells {
						if day22_can_apply_spell(i, &item.player, &item.boss) {
							bfsSpellQueue = append(bfsSpellQueue, Day22BfsData{
								spellID:   i,
								player:    item.player,
								boss:      item.boss,
								manaSpent: item.manaSpent,
							})
						}
					}
				}
			}
		}
	}

	return minManaSpent
}

func day22() (string, string) {
	return fmt.Sprint(day22_solution(false)), fmt.Sprint(day22_solution(true))
}
