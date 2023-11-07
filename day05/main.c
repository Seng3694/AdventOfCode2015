#include <stdio.h>
#include <string.h>
#include <aoc/aoc.h>

typedef struct {
  i32 part1;
  i32 part2;
  size_t length;
} context;

static const u16 bad_strings[] = {0x6162, 0x6364, 0x7071, 0x7879};

static void begin(char *line, void *userData) {
  context *ctx = userData;
  ctx->length = strlen(line);
}

static void process(char *line, void *userData, size_t lineNum) {
  (void)lineNum;
  context *ctx = userData;
  i8 vowelCount = 0;
  for (size_t i = 0; i < ctx->length; ++i) {
    vowelCount += (line[i] == 'a' || line[i] == 'e' || line[i] == 'i' ||
                   line[i] == 'o' || line[i] == 'u');
  }
  bool hasDoubleLetter = false;
  char previous = line[0];
  for (size_t i = 1; i < ctx->length; ++i) {
    if (line[i] == previous) {
      hasDoubleLetter = true;
      break;
    }
    previous = line[i];
  }
  bool hasNoBadString = true;
  for (size_t i = 0; i < ctx->length - 1; ++i) {
    const u16 c = ((u16)line[i] << 8) | (u16)line[i + 1];
    if (c == bad_strings[0] || c == bad_strings[1] || c == bad_strings[2] ||
        c == bad_strings[3]) {
      hasNoBadString = false;
      break;
    }
  }
  bool hasPairTwice = false;
  for (size_t i = 0; i < ctx->length - 3; ++i) {
    for (size_t j = i + 2; j < ctx->length; ++j) {
      if (line[i] == line[j] && line[i + 1] == line[j + 1]) {
        hasPairTwice = true;
        goto leave_has_pair_twice;
      }
    }
  }
leave_has_pair_twice:;

  bool hasSandwich = false;
  for (size_t i = 0; i < ctx->length - 2; ++i) {
    if (line[i] == line[i + 2] && line[i] != line[i + 1]) {
      hasSandwich = true;
      break;
    }
  }

  ctx->part1 += (vowelCount >= 3 && hasDoubleLetter && hasNoBadString);
  ctx->part2 += (hasPairTwice && hasSandwich);
}

static void end(char *line, void *userData, size_t lineNum) {
  (void)line;
  (void)userData;
  (void)lineNum;
}

int main(void) {
  context ctx = {0};
  aoc_file_read_lines3("day05/input.txt", begin, process, end, &ctx);
  printf("%d\n%d\n", ctx.part1, ctx.part2);
}
