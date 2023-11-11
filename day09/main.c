#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <aoc/aoc.h>
#include <aoc/permutation.h>

#define MAX_LOCATIONS 8

typedef struct {
  u32 distances[MAX_LOCATIONS];
} location;

typedef struct __attribute__((packed)) {
  // lookup
  location locations[MAX_LOCATIONS];
  u8 locationCount;

  // this section is for parsing
  i8 start;
  u8 dest;
  char currentId[4];

  // this section is for the solutions
  u32 shortest;
  u32 longest;
} context;

static inline void skip_to_number(char *str, char **out) {
  while (!isdigit(*str))
    str++;
  *out = str;
}

static void parse(char *line, context *const ctx) {
  if (memcmp(line, ctx->currentId, 4) != 0) {
    memcpy(ctx->currentId, line, 4);
    ctx->start++;
    ctx->locationCount = AOC_MAX(ctx->locationCount, ctx->dest);
    ctx->dest = ctx->start + 1;
  }
  skip_to_number(line, &line);
  const u32 distance = strtoul(line, NULL, 10);
  ctx->locations[ctx->start].distances[ctx->dest] = distance;
  ctx->locations[ctx->dest].distances[ctx->start] = distance;
  ctx->dest++;
}

static void calc_distance(const size_t *const indices, const size_t length,
                          context *const ctx) {
  u32 distance = 0;
  for (size_t i = 0; i < length - 1; ++i)
    distance += ctx->locations[indices[i]].distances[indices[i + 1]];
  ctx->shortest = AOC_MIN(ctx->shortest, distance);
  ctx->longest = AOC_MAX(ctx->longest, distance);
}

int main(void) {
  context ctx = {.start = -1, .shortest = UINT32_MAX};
  aoc_file_read_lines1("day09/input.txt", (aoc_line_func)parse, &ctx);
  aoc_permutations(ctx.locationCount, ctx.locationCount,
                   (aoc_permutation_action)calc_distance, &ctx);
  printf("%u\n%u\n", ctx.shortest, ctx.longest);
}
