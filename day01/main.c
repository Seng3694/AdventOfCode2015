#include <aoc/aoc.h>
#include <aoc/mem.h>
#include <stdio.h>

static i32 solve_part1(const char *input) {
  i32 solution = 0;
  while (*input)
    solution += *input++ == '(' ? 1 : -1;
  return solution;
}

static i32 solve_part2(const char *input) {
  for (i32 i = 0, floor = 0;; ++i) {
    floor += input[i] == '(' ? 1 : -1;
    if (floor == -1)
      return i + 1;
  }
  return 0; // should not reach
}

int main(void) {
  char *input = aoc_file_read_all2("day01/input.txt");
  printf("%d\n%d\n", solve_part1(input), solve_part2(input));
  aoc_free(input);
}
