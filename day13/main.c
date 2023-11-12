#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <aoc/aoc.h>
#include <aoc/permutation.h>
#include <aoc/mem.h>

#define MAX_PERSONS 10

typedef struct {
  i8 happiness[MAX_PERSONS];
} person;

typedef struct __attribute__((packed)) {
  person persons[MAX_PERSONS];
  u8 personCount;
  i32 happiness;
} context;

static inline void skip_to(char *line, char **out, const char dest) {
  while (*line != dest && *line != '\0')
    line++;
  *out = line + (*line == '\0' ? -1 : 0);
}

static void parse(char *contents, context *const ctx) {
  u8 person = 0;
  u8 neighbor = 0;
  do {
    char current = contents[0];
    while (*contents == current) {
      if (neighbor == person)
        neighbor++;
      skip_to(contents, &contents, ' ');
      const i32 value =
          (contents[7] == 'l' ? -1 : 1) * strtoul(contents + 12, &contents, 10);
      ctx->persons[person].happiness[neighbor++] = (i8)value;
      skip_to(contents, &contents, '\n');
      contents++;
    }
    person++;
    neighbor = 0;
  } while (*contents);

  ctx->personCount = person;
}

static void calc_happiness(const size_t *const indices, const size_t length,
                           context *const ctx) {
  i32 happiness = 0;
  for (size_t i = 0; i < length - 1; ++i) {
    happiness += ctx->persons[indices[i]].happiness[indices[i + 1]];
    happiness += ctx->persons[indices[i + 1]].happiness[indices[i]];
  }

  happiness += ctx->persons[indices[0]].happiness[indices[length - 1]];
  happiness += ctx->persons[indices[length - 1]].happiness[indices[0]];

  if (happiness > ctx->happiness)
    ctx->happiness = happiness;
}

int main(void) {
  context ctx = {.happiness = INT32_MIN};
  char *input = aoc_file_read_all2("day13/input.txt");
  parse(input, &ctx);

  aoc_permutations(ctx.personCount, ctx.personCount,
                   (aoc_permutation_action)calc_happiness, &ctx);
  const i32 part1 = ctx.happiness;

  // add one more person with 0 happiness changes
  ctx.happiness = INT32_MIN;
  for (u8 i = 0; i < ctx.personCount; ++i) {
    ctx.persons[ctx.personCount].happiness[i] = 0;
    ctx.persons[i].happiness[ctx.personCount] = 0;
  }
  ctx.personCount++;

  aoc_permutations(ctx.personCount, ctx.personCount,
                   (aoc_permutation_action)calc_happiness, &ctx);

  printf("%d\n%d\n", part1, ctx.happiness);

  aoc_free(input);
}
