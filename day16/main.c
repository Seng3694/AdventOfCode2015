#include <aoc/filesystem.h>
#include <stdio.h>
#include <stdlib.h>
#include <aoc/aoc.h>

typedef struct {
  i16 children;
  i16 cats;
  i16 samoyeds;
  i16 pomeranians;
  i16 akitas;
  i16 vizslas;
  i16 goldfish;
  i16 trees;
  i16 cars;
  i16 perfumes;
} aunt;

#define AOC_T aunt
#include <aoc/vector.h>

static void parse_property(char *line, char **out, aunt *const a) {
  switch (line[0]) {
  case 'a': /* akitas */ {
    a->akitas = (i16)strtoul(line + 8, &line, 10);
    break;
  }
  case 'c': /* cars, cats, children */ {
    switch (line[2]) {
    case 'r': /* cars */ {
      a->cars = (i16)strtoul(line + 6, &line, 10);
      break;
    }
    case 't': /* cats */ {
      a->cats = (i16)strtoul(line + 6, &line, 10);
      break;
    }
    case 'i': /* children */ {
      a->children = (i16)strtoul(line + 10, &line, 10);
      break;
    }
    }
    break;
  }
  case 'g': /* goldfish */ {
    a->goldfish = (i16)strtoul(line + 10, &line, 10);
    break;
  }
  case 'p': /* perfumes, pomeranians */ {
    if (line[1] == 'e')
      a->perfumes = (i16)strtoul(line + 10, &line, 10);
    else
      a->pomeranians = (i16)strtoul(line + 13, &line, 10);
    break;
  }
  case 's': /* samoyeds */ {
    a->samoyeds = (i16)strtoul(line + 10, &line, 10);
    break;
  }
  case 't': /* trees */ {
    a->trees = (i16)strtoul(line + 7, &line, 10);
    break;
  }
  case 'v': /* vizslas */ {
    a->vizslas = (i16)strtoul(line + 9, &line, 10);
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
  aunt a = {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1};
  parse_property(line, &line, &a);
  parse_property(line + 1, &line, &a);
  parse_property(line + 1, &line, &a);
  aoc_vector_aunt_push(aunts, a);
}

static const aunt search = {
    .children = 3,
    .cats = 7,
    .samoyeds = 2,
    .pomeranians = 3,
    .akitas = 0,
    .vizslas = 0,
    .goldfish = 5,
    .trees = 3,
    .cars = 2,
    .perfumes = 1,
};

#define PROP_SCORE(prop) (a[i].prop == search.prop || a[i].prop == -1)
#define PROP_SCORE_GT(prop) (a[i].prop > search.prop || a[i].prop == -1)
#define PROP_SCORE_LT(prop) (a[i].prop < search.prop || a[i].prop == -1)

static u16 solve_part1(const aoc_vector_aunt *const aunts) {
  u16 result = 0;
  u8 maxScore = 0;
  const aunt *const a = aunts->items;
  for (size_t i = 0; i < aunts->length; ++i) {
    u8 score = 0;
    score += PROP_SCORE(children) + PROP_SCORE(cats) + PROP_SCORE(samoyeds) +
             PROP_SCORE(pomeranians) + PROP_SCORE(akitas) +
             PROP_SCORE(vizslas) + PROP_SCORE(goldfish) + PROP_SCORE(trees) +
             PROP_SCORE(cars) + PROP_SCORE(perfumes);
    if (score > maxScore) {
      maxScore = score;
      result = i + 1;
    }
  }
  return result;
}

static u16 solve_part2(const aoc_vector_aunt *const aunts) {
  u16 result = 0;
  u8 maxScore = 0;
  const aunt *const a = aunts->items;
  for (size_t i = 0; i < aunts->length; ++i) {
    u8 score = 0;
    score += PROP_SCORE(children) + PROP_SCORE_GT(cats) + PROP_SCORE(samoyeds) +
             PROP_SCORE_LT(pomeranians) + PROP_SCORE(akitas) +
             PROP_SCORE(vizslas) + PROP_SCORE_LT(goldfish) +
             PROP_SCORE_GT(trees) + PROP_SCORE(cars) + PROP_SCORE(perfumes);
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

  printf("%u\n%u\n", solve_part1(&aunts), solve_part2(&aunts));

  aoc_vector_aunt_destroy(&aunts);
}
