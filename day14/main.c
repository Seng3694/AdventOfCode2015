#include <aoc/common.h>
#include <stdio.h>
#include <stdlib.h>
#include <aoc/aoc.h>

#define GOAL_TIME 2503

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

static inline u32 min(const u32 a, const u32 b) {
  return AOC_MIN(a, b);
}

static inline u32 calculate_distance(const reindeer *const r, const u32 time) {
  return (time / (r->duration + r->rest)) * r->speed * r->duration +
         min(time % (r->duration + r->rest), r->duration) * r->speed;
}

static u32 solve_part1(const aoc_vector_reindeer *const reindeers) {
  u32 distance = 0;
  for (size_t i = 0; i < reindeers->length; ++i) {
    const reindeer *const r = &reindeers->items[i];
    const u32 dist = calculate_distance(r, GOAL_TIME);
    if (dist > distance)
      distance = dist;
  }
  return distance;
}

static u32 solve_part2(const aoc_vector_reindeer *const r) {
  u32 scores[r->length];
  for (size_t i = 0; i < r->length; ++i)
    scores[i] = 0;

  size_t index = 0;
  u32 furthest = 0;
  for (u32 i = 1; i <= GOAL_TIME; ++i) {
    furthest = 0;
    for (size_t j = 0; j < r->length; ++j) {
      const u32 dist = calculate_distance(&r->items[j], i);
      if (dist > furthest) {
        furthest = dist;
        index = j;
      }
    }
    scores[index]++;
  }

  u32 highestScore = 0;
  for (size_t i = 0; i < r->length; ++i) {
    if (scores[i] > highestScore)
      highestScore = scores[i];
  }

  return highestScore;
}

int main(void) {
  aoc_vector_reindeer reindeers = {0};
  aoc_vector_reindeer_create(&reindeers, 1 << 4);
  aoc_file_read_lines1("day14/input.txt", (aoc_line_func)parse, &reindeers);

  printf("%u\n%u\n", solve_part1(&reindeers), solve_part2(&reindeers));

  aoc_vector_reindeer_destroy(&reindeers);
}
