#include <aoc/aoc.h>
#include <bits/time.h>
#include <complex.h>
#include <stdio.h>
#include <stdlib.h>

#define MAX_INGREDIENTS 4

typedef struct __attribute__((packed)) {
  i8 capacity;
  i8 durability;
  i8 flavor;
  i8 texture;
  i8 calories;
} ingredient;

ingredient ingredients[MAX_INGREDIENTS] = {0};
u8 count;

static void parse(char *line, void *userData) {
  (void)userData;
  while (*line != ' ')
    line++;
  ingredient *const i = &ingredients[count];
  i->capacity = strtol(line + 10, &line, 10);
  i->durability = strtol(line + 13, &line, 10);
  i->flavor = strtol(line + 9, &line, 10);
  i->texture = strtol(line + 10, &line, 10);
  i->calories = strtol(line + 11, &line, 10);
  ++count;
}

static inline i32 zero_clamp(const i32 n) {
  return (n + ((n + (n >> 28)) ^ (n >> 28))) >> 1;
}

static i32 calc_score(const u8 amounts[const MAX_INGREDIENTS]) {
  i32 c = 0, d = 0, f = 0, t = 0;
  for (u8 i = 0; i < count; ++i) {
    c += ingredients[i].capacity * amounts[i];
    d += ingredients[i].durability * amounts[i];
    f += ingredients[i].flavor * amounts[i];
    t += ingredients[i].texture * amounts[i];
  }
  return zero_clamp(c) * zero_clamp(d) * zero_clamp(f) * zero_clamp(t);
}

static inline i32 calc_calories(const u8 amounts[const MAX_INGREDIENTS]) {
  i32 c = 0;
  for (u8 i = 0; i < count; ++i)
    c += ingredients[i].calories * amounts[i];
  return c;
}

static inline i32 calc_score_with_cal(const u8 amounts[const MAX_INGREDIENTS]) {
  return (calc_calories(amounts) <= 500) * calc_score(amounts);
}

// this assumes 4 ingredients and doesn't work with less or more
static i32 solve(i32 (*calc)(const u8[const])) {
  u8 amounts[MAX_INGREDIENTS] = {0};
  i32 bestScore = 0;
  i32 score = 0;
  for (u8 i = 100; i > 0; --i) {
    amounts[0] = i;
    for (u8 j = 0; j <= (100 - i); ++j) {
      amounts[1] = j;
      for (u8 k = 0; k <= (100 - i - j); ++k) {
        amounts[2] = k;
        amounts[3] = (100 - i - j - k);
        score = calc(amounts);
        if (score > bestScore)
          bestScore = score;
      }
    }
  }
  return bestScore;
}

int main(void) {
  aoc_file_read_lines1("day15/input.txt", parse, NULL);
  printf("%d\n%d\n", solve(calc_score), solve(calc_score_with_cal));
}
