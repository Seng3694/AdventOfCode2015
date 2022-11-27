package main

import (
	"aocutil"
	"math"
)

const (
	BATTLE_STATE_ONGOING int = iota
	BATTLE_STATE_WON
	BATTLE_STATE_LOST
)

type BossStats struct {
	hp, damage  int
	poisonTimer int
}

type PlayerStats struct {
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

type Spell struct {
	cost, timer, value int
}

var (
	spells = []Spell{
		{cost: 53, timer: 0, value: 4},    // Magic Missile
		{cost: 73, timer: 0, value: 2},    // Drain
		{cost: 113, timer: 6, value: 7},   // Shield
		{cost: 173, timer: 6, value: 3},   // Poison
		{cost: 229, timer: 5, value: 101}, // Recharge
	}
)

func can_apply_spell(id int, player *PlayerStats, boss *BossStats) bool {
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

func cast_spell(id int, player *PlayerStats, boss *BossStats) {
	switch id {
	case SPELL_MAGIC_MISSILE:
		boss.hp -= spells[id].value
	case SPELL_DRAIN: //not sure if there is maxhp. "heal" implies it
		player.hp = aocutil.Clamp(player.hp+spells[id].value, 0, player.maxHp)
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

func update_effects(player *PlayerStats, boss *BossStats) {
	player.shieldTimer = aocutil.Clamp(player.shieldTimer-1, -1, math.MaxInt)
	player.rechargeTimer = aocutil.Clamp(player.rechargeTimer-1, -1, math.MaxInt)
	boss.poisonTimer = aocutil.Clamp(boss.poisonTimer-1, -1, math.MaxInt)

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

func set_battle_state(player *PlayerStats, boss *BossStats, ignoreMana bool) int {
	if boss.hp <= 0 {
		return BATTLE_STATE_WON
	} else if player.hp <= 0 || (!ignoreMana && player.mana < spells[SPELL_MAGIC_MISSILE].cost) {
		return BATTLE_STATE_LOST
	} else {
		return BATTLE_STATE_ONGOING
	}
}

type BfsData struct {
	spellID, manaSpent, battleState int
	player                          PlayerStats
	boss                            BossStats
}

func player_turn(data *BfsData, isHardmode bool) int {
	if isHardmode {
		data.player.hp--
		data.battleState = set_battle_state(&data.player, &data.boss, true)
		if data.battleState == BATTLE_STATE_LOST {
			return BATTLE_STATE_LOST
		}
	}

	update_effects(&data.player, &data.boss)
	//check for battlestate. boss could die from poison or player has no more mana
	data.battleState = set_battle_state(&data.player, &data.boss, false)
	if data.battleState != BATTLE_STATE_ONGOING {
		return data.battleState
	}

	cast_spell(data.spellID, &data.player, &data.boss)
	data.manaSpent += spells[data.spellID].cost

	data.battleState = set_battle_state(&data.player, &data.boss, true)
	if data.battleState == BATTLE_STATE_WON {
		return BATTLE_STATE_WON
	}

	return BATTLE_STATE_ONGOING
}

func boss_turn(data *BfsData) int {
	update_effects(&data.player, &data.boss)
	//check for battlestate. boss could die from poison
	data.battleState = set_battle_state(&data.player, &data.boss, true)
	if data.battleState == BATTLE_STATE_WON {
		return BATTLE_STATE_WON
	}

	damage := aocutil.Clamp(data.boss.damage-data.player.armor, 1, data.player.hp)
	data.player.hp -= damage
	data.battleState = set_battle_state(&data.player, &data.boss, true)
	if data.battleState == BATTLE_STATE_LOST {
		return BATTLE_STATE_LOST
	}

	return BATTLE_STATE_ONGOING
}

func day22_solution(isHardmode bool) int {
	bossStats := BossStats{
		hp:          55,
		damage:      8,
		poisonTimer: -1,
	}

	playerStats := PlayerStats{
		maxHp:         50,
		hp:            50,
		mana:          500,
		armor:         0,
		shieldTimer:   -1,
		rechargeTimer: -1,
	}
	minManaSpent := math.MaxInt

	for k := 0; k < SPELL_COUNT; k++ {
		bfsSpellQueue := make([]BfsData, 0, 64)
		//starting spell. build bfs tree from here
		bfsSpellQueue = append(bfsSpellQueue, BfsData{
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

				state := player_turn(&item, isHardmode)
				if item.manaSpent > minManaSpent || state == BATTLE_STATE_LOST {
					continue
				} else if state == BATTLE_STATE_WON {
					if item.manaSpent < minManaSpent {
						minManaSpent = item.manaSpent
					}
					continue
				}

				state = boss_turn(&item)
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
						if can_apply_spell(i, &item.player, &item.boss) {
							bfsSpellQueue = append(bfsSpellQueue, BfsData{
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

func main() {
	aocutil.AOCFinish(day22_solution(false), day22_solution(true))
}
