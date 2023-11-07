#include <stdio.h>
#include <stdlib.h>
#include <aoc/aoc.h>

#define GRID_SIZE 1000
static i32 grid[GRID_SIZE * GRID_SIZE] = {0};
static i32 brightness[GRID_SIZE * GRID_SIZE] = {0};

typedef enum {
  INSTRUCTION_TYPE_TURN_OFF,
  INSTRUCTION_TYPE_TURN_ON,
  INSTRUCTION_TYPE_TOGGLE,
} instruction_type;

static const size_t parse_advance[] = {9, 8, 7};
static const i32 instruction_results[][2] = {{0, 0}, {1, 1}, {1, 0}};
static const i32 brightness_changes[] = {-1, 1, 2};

typedef struct {
  instruction_type type;
  i32 from_x, from_y, to_x, to_y;
} instruction;

#define AOC_T instruction
#include <aoc/vector.h>

static void parse(char *line, void *userData) {
  instruction instr = {0};
  // clang-format off
  switch (line[6]) {
  case ' ': instr.type = INSTRUCTION_TYPE_TOGGLE; break;
  case 'f': instr.type = INSTRUCTION_TYPE_TURN_OFF; break;
  case 'n': instr.type = INSTRUCTION_TYPE_TURN_ON; break;
  }
  // clang-format on
  instr.from_x = strtoul(line + parse_advance[instr.type], &line, 10);
  instr.from_y = strtoul(line + 1, &line, 10);
  instr.to_x = strtoul(line + 9, &line, 10);
  instr.to_y = strtoul(line + 1, NULL, 10);
  aoc_vector_instruction_push(userData, instr);
}

static inline i32 lower_clamp(const i32 n) {
  return (n + ((n + (n >> 28)) ^ (n >> 28))) >> 1;
}

static inline void run_instruction(const instruction *const instr) {
  i32 *g = &grid[instr->from_y * GRID_SIZE];
  i32 *b = &brightness[instr->from_y * GRID_SIZE];
  for (i32 y = instr->from_y; y <= instr->to_y; ++y) {
    for (i32 x = instr->from_x; x <= instr->to_x; ++x) {
      g[x] = instruction_results[instr->type][g[x]];
      b[x] = lower_clamp(b[x] + brightness_changes[instr->type]);
    }
    g += GRID_SIZE;
    b += GRID_SIZE;
  }
}

static void solve(const aoc_vector_instruction *const instructions) {
  for (size_t i = 0; i < instructions->length; ++i)
    run_instruction(&instructions->items[i]);
  i32 activeLights = 0;
  i32 totalBrightness = 0;
  for (i32 i = 0; i < GRID_SIZE * GRID_SIZE; ++i) {
    activeLights += grid[i];
    totalBrightness += brightness[i];
  }
  printf("%d\n%d\n", activeLights, totalBrightness);
}

int main(void) {
  aoc_vector_instruction instructions = {0};
  aoc_vector_instruction_create(&instructions, 1 << 9);
  aoc_file_read_lines1("day06/input.txt", parse, &instructions);
  solve(&instructions);
  aoc_vector_instruction_destroy(&instructions);
}
