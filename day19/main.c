#include <aoc/string.h>
#include <stdio.h>
#include <string.h>
#include <aoc/aoc.h>
#include <aoc/mem.h>

typedef struct {
  const char *str;
  size_t len;
} slice;

typedef struct {
  slice search;
  slice replace;
} rule;

#define AOC_T rule
#include <aoc/vector.h>

static inline void skip_until(char *line, char **out, const char c) {
  while (*line != c)
    line++;
  *out = line;
}

static rule parse_rule(char *line, char **out) {
  rule r = {0};
  r.search.str = line;
  skip_until(line, &line, ' ');
  r.search.len = line - r.search.str;

  r.replace.str = line + 4;
  skip_until(line + 4, &line, '\n');
  r.replace.len = line - r.replace.str;

  *out = line + 1;
  return r;
}

static void parse(char *contents, aoc_vector_rule *const rules,
                  slice *const input) {
  while (*contents != '\n')
    aoc_vector_rule_push(rules, parse_rule(contents, &contents));

  input->str = contents + 1;
  input->len = strlen(input->str);
}

typedef struct {
  slice data[3];
  u8 count;
} slices;

static inline u32 slices_hash(const slices *const s) {
  u32 hash = 74712713;
  for (u8 i = 0; i < s->count; ++i) {
    for (size_t j = 0; j < s->data[i].len; ++j) {
      hash = (hash << 5) + s->data[i].str[j];
    }
  }
  return hash;
}

static bool slices_equal(const slices *const a, const slices *const b) {
  size_t len1 = 0, len2 = 0;
  for (u8 i = 0; i < a->count; ++i)
    len1 += a->data[i].len;
  for (u8 i = 0; i < b->count; ++i)
    len2 += b->data[i].len;
  if (len1 != len2)
    return false;

  size_t aSliceIndex = 0, bSliceIndex = 0;
  size_t aStrIndex = 0, bStrIndex = 0;
  for (size_t i = 0; i < len1; ++i) {
    if (aStrIndex == a->data[aSliceIndex].len) {
      aStrIndex = 0;
      aSliceIndex++;
    }
    if (bStrIndex == b->data[bSliceIndex].len) {
      bStrIndex = 0;
      bSliceIndex++;
    }
    if (a->data[aSliceIndex].str[aStrIndex] !=
        b->data[bSliceIndex].str[bStrIndex])
      return false;

    aStrIndex++;
    bStrIndex++;
  }

  return true;
}

#define AOC_T slices
#define AOC_T_EMPTY ((slices){0})
#define AOC_T_HASH slices_hash
#define AOC_T_EQUALS slices_equal
#define AOC_BASE2_CAPACITY
#include <aoc/set.h>

static u32 solve_part1(const aoc_vector_rule *const rules,
                       const slice *const input) {
  aoc_set_slices slc = {0};
  aoc_set_slices_create(&slc, 1 << 10);

  for (size_t i = 0; i < rules->length; ++i) {
    const rule *const r = &rules->items[i];
    const size_t len = input->len - r->search.len + 1;
    for (size_t j = 0; j < len; ++j) {
      for (size_t k = 0; k < r->search.len; ++k) {
        if (input->str[j + k] != r->search.str[k]) {
          goto no_match;
        }
      }

      slices s = {0};
      if (j != 0) {
        s.data[s.count].str = input->str;
        s.data[s.count++].len = j;

        s.data[s.count].str = r->replace.str;
        s.data[s.count++].len = r->replace.len;

        if (j != (len - 1)) {
          s.data[s.count].str = input->str + j + r->search.len;
          s.data[s.count++].len = input->len - r->search.len - (j - 1) - 1;
        }
      } else {
        s.data[s.count].str = r->replace.str;
        s.data[s.count++].len = r->replace.len;

        s.data[s.count].str = input->str + r->search.len;
        s.data[s.count++].len = input->len - r->replace.len + 1;
      }

      const u32 hash = slices_hash(&s);
      if (!aoc_set_slices_contains_pre_hashed(&slc, s, hash))
        aoc_set_slices_insert_pre_hashed(&slc, s, hash);

    no_match:;
    }
  }

  aoc_set_slices_destroy(&slc);
  return (u32)slc.count;
}

int main(void) {
  aoc_vector_rule rules = {0};
  aoc_vector_rule_create(&rules, 32);
  char *contents = aoc_file_read_all2("day19/input.txt");
  slice input = {.str = contents, .len = strlen(contents)};
  parse(contents, &rules, &input);

  const u32 part1 = solve_part1(&rules, &input);

  printf("%u\n", part1);

  aoc_vector_rule_destroy(&rules);
  aoc_free(contents);
}
