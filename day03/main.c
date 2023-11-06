#include <stdio.h>
#include <aoc/aoc.h>
#include <aoc/mem.h>

#define AOC_T i32
#include <aoc/point.h>
typedef aoc_point2_i32 point;

#define AOC_T point
#define AOC_T_EMPTY ((point){INT32_MIN, INT32_MIN})
#define AOC_T_HASH aoc_point2_i32_hash
#define AOC_T_EQUALS aoc_point2_i32_equals
#define AOC_BASE2_CAPACITY
#include <aoc/set.h>

static aoc_set_point visits = {0};

static size_t solve(const char *const input, const size_t length,
                    const i32 santaCount) {
  aoc_set_point_clear(&visits);
  aoc_set_point_insert(&visits, (point){0, 0});
  for (i32 s = 0; s < santaCount; ++s) {
    point p = {0, 0};
    for (size_t i = s; i < length; i += santaCount) {
      // clang-format off
      switch(input[i]) {
      case '<': p.x--; break;
      case '>': p.x++; break;
      case '^': p.y--; break;
      case 'v': p.y++; break;
      }
      // clang-format on
      u32 hash = aoc_point2_i32_hash(&p);
      if (!aoc_set_point_contains_pre_hashed(&visits, p, hash))
        aoc_set_point_insert_pre_hashed(&visits, p, hash);
    }
  }
  return visits.count;
}

int main(void) {
  char *input = NULL;
  size_t length = 0;
  aoc_file_read_all1("day03/input.txt", &input, &length);
  aoc_set_point_create(&visits, 1 << 12);
  printf("%zu\n%zu\n", solve(input, length, 1), solve(input, length, 2));
  aoc_set_point_destroy(&visits);
  aoc_free(input);
}