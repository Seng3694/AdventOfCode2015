#include <stdio.h>
#include <stdlib.h>
#include <aoc/aoc.h>

typedef struct {
  u8 speed, duration, rest;
} reindeer;

#define AOC_T reindeer
#include <aoc/vector.h>

static inline void skip_name(char *line, char **out) {
  while (*line != ' ')
    line++;
  *out = line;
}

static void parse(char *line, aoc_vector_reindeer *const reindeers) {
  skip_name(line, &line);
  reindeer r = {0};
  r.speed = strtoul(line + 9, &line, 10);
  r.duration = strtoul(line + 10, &line, 10);
  r.rest = strtoul(line + 33, NULL, 10);
  aoc_vector_reindeer_push(reindeers, r);
}

static u32 solve_part1(const aoc_vector_reindeer *const reindeers) {
  u32 distance = 0;
  for (size_t i = 0; i < reindeers->length; ++i) {
    const reindeer *const r = &reindeers->items[i];
    const u32 roundTime = (r->duration + r->rest);
    const u32 rounds = 2503 / roundTime;
    const u32 leftover = 2503 % roundTime;
    const u32 dist = rounds * r->speed * r->duration +
                     AOC_MIN(leftover, r->duration) * r->speed;
    if (dist > distance)
      distance = dist;
  }
  return distance;
}

int main(void) {
  aoc_vector_reindeer reindeers = {0};
  aoc_vector_reindeer_create(&reindeers, 1 << 4);
  aoc_file_read_lines1("day14/input.txt", (aoc_line_func)parse, &reindeers);

  const u32 part1 = solve_part1(&reindeers);

  printf("%u\n", part1);

  aoc_vector_reindeer_destroy(&reindeers);
}
