#include <assert.h>
#include <stdlib.h>
#include <stdio.h>
#include <aoc/aoc.h>
#include <aoc/mem.h>
#include <string.h>

static inline bool is_integer(const char c) {
  return c == '-' || (c >= '0' && c <= '9');
}

static i32 solve_part1(char *input) {
  i32 sum = 0;
  while (*input) {
    if (is_integer(*input))
      sum += strtol(input, &input, 10);
    input++;
  }
  return sum;
}

static inline void skip_string(char *input, char **out) {
  assert(*input == '"');
  input++;
  while (*input != '"')
    input++;
  *out = input + 1;
}

static inline void skip_current_object(char *input, char **out) {
  // leaves the current object
  i32 depth = 1;
  while (*input != '}' || depth != 0) {
    input++;
    switch (*input) {
    case '{':
      ++depth;
      break;
    case '}':
      --depth;
      break;
    }
  };
  *out = input;
}

static i32 json_object(char *input, char **out);
static i32 json_array(char *input, char **out);

static i32 json_array(char *input, char **out) {
  i32 sum = 0;
  assert(*input == '[');
  input++;

  for (;;) {
    switch (*input) {
    case '"':
      skip_string(input, &input);
      break;
    case ',':
      input++;
      break;
    case '{':
      sum += json_object(input, &input);
      break;
    case '[':
      sum += json_array(input, &input);
      break;
    case ']':
      goto done;
    default:
      sum += strtol(input, &input, 10);
      break;
    }
  }

done:
  *out = input + 1;
  return sum;
}

static i32 json_object(char *input, char **out) {
  i32 sum = 0;
  assert(*input == '{');

  while (*input != '}') {
    skip_string(input + 1, &input); // name
    input += 1;                     // :

    switch (*input) {
    case '"':
      // if it contains the word "red" skip the whole group
      if (memcmp(input + 1, "red\"", 4) == 0) {
        sum = 0;
        skip_current_object(input, &input);
        goto end;
      }
      skip_string(input, &input);
      break;
    case '{':
      sum += json_object(input, &input);
      break;
    case '[':
      sum += json_array(input, &input);
      break;
    default:
      sum += strtol(input, &input, 10);
      break;
    }

    assert(*input == ',' || *input == '}');
  }

end:
  assert(*input == '}');
  *out = input + 1;
  return sum;
}

static inline i32 solve_part2(char *input) {
  return json_object(input, &input);
}

int main(void) {
  char *input = aoc_file_read_all2("day12/input.txt");
  printf("%d\n%d\n", solve_part1(input), solve_part2(input));
  aoc_free(input);
}
