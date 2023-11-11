#include <aoc/common.h>
#include <stdio.h>
#include <aoc/aoc.h>
#include <aoc/mem.h>

static const char input[] = "hxbxwxba";
static const i32 input_len = (i32)AOC_ARRAY_SIZE(input) - 1;

static inline void default_sequence(char pwd[input_len], const i32 start) {
  char c = 'a' - 1;
  for (i32 i = start; i < input_len; ++i) {
    c += (i & 1);
    pwd[i] = c;
  }
}

static inline void increase(char pwd[input_len]) {
  for (i32 i = input_len - 1; i > 0; --i) {
    if (pwd[i] == 'z') {
      pwd[i] = 'a';
    } else {
      pwd[i]++;
      break;
    }
  }
}

static inline bool invalid_character_check(char pwd[const input_len]) {
  for (i32 i = 0; i < input_len; ++i) {
    if (pwd[i] == 'i' || pwd[i] == 'o' || pwd[i] == 'l') {
      pwd[i]++;
      i++;
      default_sequence(pwd, i);
      return false;
    }
  }
  return true;
}

static inline bool increasing_pattern_check(char pwd[const input_len]) {
  i8 increasingPatternCount = 1;
  for (i32 i = 0; i < input_len - 1; ++i) {
    if (pwd[i] + 1 == pwd[i + 1]) {
      increasingPatternCount++;
    } else {
      if (increasingPatternCount >= 3)
        break;
      increasingPatternCount = 1;
    }
  }
  return increasingPatternCount >= 3;
}

static inline bool different_pairs_check(char pwd[const input_len]) {
  i8 differentPairsCount = 0;
  char previous = '0';
  for (i32 i = 0; i < input_len - 1; ++i) {
    if (previous != pwd[i] && pwd[i] == pwd[i + 1]) {
      differentPairsCount++;
      previous = pwd[i++];
    }
  }
  return differentPairsCount >= 2;
}

static void next_password(const char pwd[const input_len],
                          char output[const input_len]) {
  aoc_mem_copy(output, pwd, input_len);
  for (;;) {
    increase(output);
    if (invalid_character_check(output) && increasing_pattern_check(output) &&
        different_pairs_check(output))
      break;
  }
}

int main(void) {
  char part1[input_len];
  next_password(input, part1);
  char part2[input_len];
  next_password(part1, part2);
  printf("%.*s\n%.*s\n", input_len, part1, input_len, part2);
}
