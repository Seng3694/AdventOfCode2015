#include <aoc/filesystem.h>
#include <stdio.h>
#include <stdlib.h>
#include <aoc/aoc.h>

typedef enum {
  PROP_CHILDREN,
  PROP_SAMOYEDS,
  PROP_AKITAS,
  PROP_VIZLAS,
  PROP_CARS,
  PROP_PERFUMES,
  PROP_CATS,
  PROP_TREES,
  PROP_POMERANIANS,
  PROP_GOLDFISH,
  PROP_COUNT,
} property;

typedef struct {
  i16 properties[PROP_COUNT];
} aunt;

#define AOC_T aunt
#include <aoc/vector.h>

static void parse_property(char *line, char **out, aunt *const a) {
  switch (line[0]) {
  case 'a': /* akitas */ {
    a->properties[PROP_AKITAS] = (i16)strtoul(line + 8, &line, 10);
    break;
  }
  case 'c': /* cars, cats, children */ {
    switch (line[2]) {
    case 'r': /* cars */ {
      a->properties[PROP_CARS] = (i16)strtoul(line + 6, &line, 10);
      break;
    }
    case 't': /* cats */ {
      a->properties[PROP_CATS] = (i16)strtoul(line + 6, &line, 10);
      break;
    }
    case 'i': /* children */ {
      a->properties[PROP_CHILDREN] = (i16)strtoul(line + 10, &line, 10);
      break;
    }
    }
    break;
  }
  case 'g': /* goldfish */ {
    a->properties[PROP_GOLDFISH] = (i16)strtoul(line + 10, &line, 10);
    break;
  }
  case 'p': /* perfumes, pomeranians */ {
    if (line[1] == 'e')
      a->properties[PROP_PERFUMES] = (i16)strtoul(line + 10, &line, 10);
    else
      a->properties[PROP_POMERANIANS] = (i16)strtoul(line + 13, &line, 10);
    break;
  }
  case 's': /* samoyeds */ {
    a->properties[PROP_SAMOYEDS] = (i16)strtoul(line + 10, &line, 10);
    break;
  }
  case 't': /* trees */ {
    a->properties[PROP_TREES] = (i16)strtoul(line + 7, &line, 10);
    break;
  }
  case 'v': /* vizslas */ {
    a->properties[PROP_VIZLAS] = (i16)strtoul(line + 9, &line, 10);
    break;
  }
  }
  *out = line + 1;
}

static void parse(char *line, void *aunts) {
  // skip name
  while (*line != ':')
    line++;
  line += 2;
  aunt a = {{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}};
  parse_property(line, &line, &a);
  parse_property(line + 1, &line, &a);
  parse_property(line + 1, &line, &a);
  aoc_vector_aunt_push(aunts, a);
}

static const aunt search = {{3, 2, 0, 0, 2, 1, 7, 3, 3, 5}};

static inline u8 prop_score(const aunt *const a, const property prop) {
  return a->properties[prop] == search.properties[prop] ||
         a->properties[prop] == -1;
}

static inline u8 prop_score_gt(const aunt *const a, const property prop) {
  return a->properties[prop] > search.properties[prop] ||
         a->properties[prop] == -1;
}

static inline u8 prop_score_lt(const aunt *const a, const property prop) {
  return a->properties[prop] < search.properties[prop] ||
         a->properties[prop] == -1;
}

typedef u8 (*score_func)(const aunt *const, const property);

static const score_func part1_funcs[] = {
    [PROP_CHILDREN] = prop_score,    [PROP_SAMOYEDS] = prop_score,
    [PROP_AKITAS] = prop_score,      [PROP_VIZLAS] = prop_score,
    [PROP_CARS] = prop_score,        [PROP_PERFUMES] = prop_score,
    [PROP_CATS] = prop_score,        [PROP_TREES] = prop_score,
    [PROP_POMERANIANS] = prop_score, [PROP_GOLDFISH] = prop_score,
};

static const score_func part2_funcs[] = {
    [PROP_CHILDREN] = prop_score,       [PROP_SAMOYEDS] = prop_score,
    [PROP_AKITAS] = prop_score,         [PROP_VIZLAS] = prop_score,
    [PROP_CARS] = prop_score,           [PROP_PERFUMES] = prop_score,
    [PROP_CATS] = prop_score_gt,        [PROP_TREES] = prop_score_gt,
    [PROP_POMERANIANS] = prop_score_lt, [PROP_GOLDFISH] = prop_score_lt,
};

static u16 solve(const aoc_vector_aunt *const aunts,
                 const score_func funcs[const PROP_COUNT]) {
  u16 result = 0;
  u8 maxScore = 0;
  for (size_t i = 0; i < aunts->length; ++i) {
    u8 score = 0;
    for (property p = 0; p < PROP_COUNT; ++p)
      score += funcs[p](&aunts->items[i], p);

    if (score > maxScore) {
      maxScore = score;
      result = i + 1;
    }
  }
  return result;
}

int main(void) {
  aoc_vector_aunt aunts = {0};
  aoc_vector_aunt_create(&aunts, 1 << 9);
  aoc_file_read_lines1("day16/input.txt", parse, &aunts);

  printf("%u\n%u\n", solve(&aunts, part1_funcs), solve(&aunts, part2_funcs));

  aoc_vector_aunt_destroy(&aunts);
}
