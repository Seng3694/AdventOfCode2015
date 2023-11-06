#include <stdio.h>
#include <aoc/aoc.h>
#include <aoc/mem.h>
#include <aoc/md5.h>

static inline bool starts_with_5_zeros(const u8 hash[const 16]) {
  return !(hash[0] | hash[1] | (hash[2] & 0xf0));
}

static inline bool starts_with_6_zeros(const u8 hash[const 16]) {
  return !(hash[0] | hash[1] | hash[2]);
}

static i32 append_number(char *dest, i32 num) {
  i32 i = 0;
  while (num != 0) {
    i32 digit = num % 10;
    dest[i++] = digit + '0';
    num /= 10;
  }
  i32 len = i;
  for (i = 0; i < len / 2; i++) {
    char temp = dest[i];
    dest[i] = dest[len - i - 1];
    dest[len - i - 1] = temp;
  }
  return len;
}

static i32 solve(char *input, const size_t length, const i32 start,
                 bool (*starts_with)(const u8 hash[const 16])) {
  u8 hash[16];
  char *numStart = input + length;
  for (i32 i = start;; ++i) {
    const i32 numLength = append_number(numStart, i);
    aoc_md5(input, length + numLength, hash);
    if (starts_with(hash))
      return i;
  }
  return 0; // should never reach
}

int main(void) {
  size_t length = 256;
  char *input = aoc_calloc(1, length);
  aoc_file_read_all1("day04/input.txt", &input, &length);
  const i32 part1 = solve(input, length, 0, starts_with_5_zeros);
  const i32 part2 = solve(input, length, part1 + 1, starts_with_6_zeros);
  printf("%d\n%d\n", part1, part2);
  aoc_free(input);
}