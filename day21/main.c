#include <aoc/aoc.h>
#include <aoc/common.h>
#include <stdint.h>
#include <stdio.h>

typedef struct {
  union {
    i32 cost, hp;
  };
  i32 damage, armor;
} stats;

static const stats boss_stats = {.hp = 109, .damage = 8, .armor = 2};
static const i32 player_hp = 100;

static const stats weapons[] = {
    {.cost = 8, .damage = 4, .armor = 0},  // Dagger
    {.cost = 10, .damage = 5, .armor = 0}, // Shortsword
    {.cost = 25, .damage = 6, .armor = 0}, // Warhammer
    {.cost = 40, .damage = 7, .armor = 0}, // Longsword
    {.cost = 74, .damage = 8, .armor = 0}, // Greataxe
};

static const stats armor[] = {
    {.cost = 0, .damage = 0, .armor = 0},   // None
    {.cost = 13, .damage = 0, .armor = 1},  // Leather
    {.cost = 31, .damage = 0, .armor = 2},  // Chainmail
    {.cost = 53, .damage = 0, .armor = 3},  // Splintmail
    {.cost = 75, .damage = 0, .armor = 4},  // Bandedmail
    {.cost = 102, .damage = 0, .armor = 5}, // Platemail
};

static const stats rings[] = {
    {.cost = 0, .damage = 0, .armor = 0},   // None 1
    {.cost = 0, .damage = 0, .armor = 0},   // None 2
    {.cost = 25, .damage = 1, .armor = 0},  // Damage +1
    {.cost = 50, .damage = 2, .armor = 0},  // Damage +2
    {.cost = 100, .damage = 3, .armor = 0}, // Damage +3
    {.cost = 20, .damage = 0, .armor = 1},  // Defense +1
    {.cost = 40, .damage = 0, .armor = 2},  // Defense +2
    {.cost = 80, .damage = 0, .armor = 3},  // Defense +3
};

static inline i32 calculate_damage(const i32 damage, const i32 armor) {
  const i32 result = damage - armor;
  return result <= 0 ? 1 : result;
}

static bool simulate(stats player, stats boss) {
  for (;;) {
    boss.hp -= calculate_damage(player.damage, boss.armor);
    if (boss.hp <= 0)
      return true;
    player.hp -= calculate_damage(boss.damage, player.armor);
    if (player.hp <= 0)
      return false;
  }
  return false; // should never reach
}

static void solve(i32 *const part1, i32 *const part2) {
  i32 minCost = INT32_MAX;
  i32 maxCost = INT32_MIN;

  for (size_t w = 0; w < AOC_ARRAY_SIZE(weapons); ++w) {
    for (size_t a = 0; a < AOC_ARRAY_SIZE(armor); ++a) {
      for (size_t r1 = 0; r1 < AOC_ARRAY_SIZE(rings) - 1; ++r1) {
        for (size_t r2 = r1 + 1; r2 < AOC_ARRAY_SIZE(rings); ++r2) {
          const i32 totalCost =
              weapons[w].cost + armor[a].cost + rings[r1].cost + rings[r2].cost;
          const stats player = {
              .hp = player_hp,
              .damage = weapons[w].damage + armor[a].damage + rings[r1].damage +
                        rings[r2].damage,
              .armor = armor[a].armor + rings[r1].armor + rings[r2].armor,
          };

          const bool won = simulate(player, boss_stats);
          if (won && minCost > totalCost) {
            minCost = totalCost;
          } else if (!won && maxCost < totalCost) {
            maxCost = totalCost;
          }
        }
      }
    }
  }

  *part1 = minCost;
  *part2 = maxCost;
}

int main(void) {
  i32 part1, part2;
  solve(&part1, &part2);
  printf("%d\n%d\n", part1, part2);
}
