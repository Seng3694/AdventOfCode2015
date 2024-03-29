#include <stdio.h>
#include <stdlib.h>
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

#define AOC_T char
#include <aoc/vector.h>

static inline u32 vec_char_hash(const aoc_vector_char *const vec) {
  return aoc_string_hash1(vec->items, vec->length);
}

typedef struct {
  const aoc_vector_char *vec;
  i64 current;
} vec_iter;

static void back_iter_init(vec_iter *const iter,
                           const aoc_vector_char *const v) {
  iter->vec = v;
  iter->current = v->length;
}

static bool find_next_index_back(vec_iter *const iter, const slice *const s,
                                 size_t *const out) {
  const i64 start = iter->current - (i64)s->len;
  if (start < 0)
    return false;

  for (i64 i = start; i >= 0; --i) {
    for (size_t j = 0; j < s->len; ++j) {
      if (iter->vec->items[i + (i64)j] != s->str[j])
        goto next;
    }
    *out = (size_t)i;
    iter->current = i - 1;
    return true;
  next:;
  }

  return false;
}

static void vec_char_replace_at(aoc_vector_char *const vec,
                                const slice *const search,
                                const slice *const replace, const size_t i) {
  const i64 diff = ((i64)replace->len - (i64)search->len);
  const size_t newLength = vec->length + diff;
  aoc_vector_char_ensure_capacity(vec, newLength);

  if (diff > 0) {
    for (i64 j = vec->length - 1; j > (i64)(i + (i64)search->len - 1); --j)
      vec->items[j + diff] = vec->items[j];
  } else if (diff < 0) {
    for (i64 j = 0; j < (i64)vec->length - (i64)(i + search->len - 1); ++j)
      vec->items[i + j + replace->len - 1] =
          vec->items[i + j + search->len - 1];
  }

  for (size_t j = 0; j < replace->len; ++j)
    vec->items[i + j] = replace->str[j];

  vec->length = newLength;
}

static inline int compare_rule(const rule *const left,
                               const rule *const right) {
  return (int)((i64)right->replace.len - (i64)left->replace.len);
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

static u32 solve_part2(const aoc_vector_rule *const rules,
                       const slice *const input) {
  aoc_vector_char chars = {0};
  aoc_vector_char_create(&chars, 1 << 12);
  aoc_mem_copy(chars.items, input->str, input->len);
  chars.length = input->len;

  u32 steps = 0;
  for (;;) {
    for (size_t i = 0; i < rules->length; ++i) {
      const rule *const r = &rules->items[i];

      vec_iter iter = {0};
      back_iter_init(&iter, &chars);

      size_t j = 0;
      size_t matches = 0;
      while (find_next_index_back(&iter, &r->replace, &j)) {
        vec_char_replace_at(&chars, &r->replace, &r->search, j);
        matches++;
        steps++;
      }

      if (matches > 0)
        break;
    }

    if (chars.length == 1 && chars.items[0] == 'e')
      break;
  }

  aoc_vector_char_destroy(&chars);
  return steps;
}

int main(void) {
  aoc_vector_rule rules = {0};
  aoc_vector_rule_create(&rules, 32);
  char *contents = aoc_file_read_all2("day19/input.txt");
  slice input = {.str = contents, .len = strlen(contents)};
  parse(contents, &rules, &input);
  qsort(rules.items, rules.length, sizeof(rule), (__compar_fn_t)compare_rule);

  printf("%u\n%u\n", solve_part1(&rules, &input), solve_part2(&rules, &input));

  aoc_vector_rule_destroy(&rules);
  aoc_free(contents);
}
