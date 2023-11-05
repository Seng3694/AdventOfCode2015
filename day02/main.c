#include <stdio.h>
#include <stdlib.h>
#include <aoc/aoc.h>

typedef struct {
  i32 part1, part2;
} solutions;

static i32 calculate_total_area(const i32 w, const i32 h, const i32 l) {
  const i32 area1 = l * w;
  const i32 area2 = w * h;
  const i32 area3 = h * l;
  i32 slack = AOC_MIN(area1, area2);
  slack = AOC_MIN(slack, area3);
  return 2 * area1 + 2 * area2 + 2 * area3 + slack;
}

static i32 calculate_ribbon_length(i32 w, i32 h, i32 l) {
  i32 temp = w;
  if (w > h) {
    w = h;
    h = temp;
  }
  if (h > l) {
    temp = h;
    h = l;
    l = temp;
    if (w > h) {
      temp = w;
      w = h;
      h = temp;
    }
  }
  return 2 * w + 2 * h + w * h * l;
}

static void solve(char *str, void *userData) {
  i32 w = strtol(str, &str, 10);
  i32 h = strtol(str + 1, &str, 10);
  i32 l = strtol(str + 1, NULL, 10);
  solutions *s = userData;
  s->part1 += calculate_total_area(w, h, l);
  s->part2 += calculate_ribbon_length(w, h, l);
}

int main(void) {
  solutions s = {0};
  aoc_file_read_lines1("day02/input.txt", solve, &s);
  printf("%d\n%d\n", s.part1, s.part2);
}
