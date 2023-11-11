#include <aoc/common.h>
#include <aoc/aoc.h>
#include <stdio.h>

#define AOC_T char
#include <aoc/vector.h>

static inline void say_convert(aoc_vector_char *const input,
                               aoc_vector_char *const output) {
  aoc_vector_char_clear(output);
  const char *n = input->items;
  u8 count = 0;
  for (size_t i = 0; i < input->length; ++i) {
    count++;
    // no bounds check necessary because array is big enough
    if (n[i] != n[i + 1]) {
      aoc_vector_char_push(output, count + '0');
      aoc_vector_char_push(output, n[i]);
      count = 0;
    }
  }
}

static size_t solve(aoc_vector_char *input, aoc_vector_char *output, i8 count) {
  // use back buffer swapping to avoid copying all the time
  while (count-- > 0) {
    say_convert(input, output);
    aoc_vector_char *tmp = input;
    input = output;
    output = tmp;
  }
  // assumes count is always dividable by 2
  return input->length;
}

int main(void) {
  const char input[] = "3113322113";
  aoc_vector_char numbers = {0};
  aoc_vector_char buffer = {0};
  aoc_vector_char_create(&numbers, 1 << 23);
  aoc_vector_char_create(&buffer, 1 << 23);
  aoc_mem_copy(numbers.items, input, AOC_ARRAY_SIZE(input));
  numbers.length = AOC_ARRAY_SIZE(input) - 1;

  const size_t part1 = solve(&numbers, &buffer, 40);
  const size_t part2 = solve(&numbers, &buffer, 10);

  printf("%zu\n%zu\n", part1, part2);

  aoc_vector_char_destroy(&numbers);
  aoc_vector_char_destroy(&buffer);
}
