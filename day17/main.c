#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <aoc/aoc.h>

static u8 container_sizes[32] = {0};
static u8 count = 0;

static void parse(char *line, void *userData) {
  (void)userData;
  container_sizes[count++] = strtoul(line, NULL, 10);
}

static void solve(u32 *const part1, u32 *const part2) {
  u32 p1 = 0;
  u32 p2 = 0;
  u32 smallest = UINT32_MAX;
  const u32 combinations = (u32)powf(2, (f32)count);
  for (u32 i = 0; i < combinations; ++i) {
    u16 capacity = 0;
    for (u8 j = 0; j < count; ++j) {
      if (AOC_CHECK_BIT(i, j))
        capacity += container_sizes[j];
    }
    if (capacity == 150) {
      const u32 containers = aoc_popcount(i);
      if (containers < smallest) {
        smallest = containers;
        p2 = 0;
      }
      p1++;
      if (containers == smallest)
        p2++;
    }
  }
  *part1 = p1;
  *part2 = p2;
}

int main(void) {
  aoc_file_read_lines1("day17/input.txt", parse, NULL);
  u32 part1, part2;
  solve(&part1, &part2);
  printf("%u\n%u\n", part1, part2);
}
