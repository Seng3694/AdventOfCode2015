#include <aoc/filesystem.h>
#include <stdio.h>
#include <string.h>
#include <aoc/aoc.h>

typedef struct {
  size_t visual, memory, encoded;
} context;

static void parse(char *line, context *ctx) {
  size_t totalLength = strlen(line);
  aoc_string_trim_right(line, &totalLength);
  ctx->visual += totalLength;
  ctx->encoded += 6; //  "\"\""

  line++; // "
  while (*line != '"') {
    switch (*line) {
    case '\\':
      switch (line[1]) {
      case 'x':
        line += 4;
        ctx->encoded += 4;
        break;
      case '\\':
      case '"':
        line += 2;
        ctx->encoded += 3;
        break;
      }
      break;
    default:
      line++;
      break;
    }
    ctx->memory += 1;
    ctx->encoded += 1;
  }
}

int main(void) {
  context ctx = {0};
  aoc_file_read_lines1("day08/input.txt", (aoc_line_func)parse, &ctx);
  printf("%zu\n", ctx.visual - ctx.memory);
  printf("%zu\n", ctx.encoded - ctx.visual);
}
