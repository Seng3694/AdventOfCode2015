#include <stdio.h>
#include <aoc/aoc.h>
#include <aoc/mem.h>

#define GRID_SIZE ((100 * 100) >> 3)

static inline void turn_on_light(u8 *const lights, const u32 i) {
  lights[i >> 3] |= (1 << (i & 7));
}

static inline void turn_off_light(u8 *const lights, const u32 i) {
  lights[i >> 3] &= ~(1 << (i & 7));
}

static inline bool light_state(const u8 *const lights, const u32 i) {
  return AOC_CHECK_BIT(lights[i >> 3], (i & 7));
}

static inline void set_light(const u8 *const lights, u8 *const buffer,
                             const u32 i, const u8 adjacent) {
  if (light_state(lights, i)) {
    if (adjacent != 2 && adjacent != 3)
      turn_off_light(buffer, i);
    else
      turn_on_light(buffer, i);
  } else {
    if (adjacent == 3)
      turn_on_light(buffer, i);
    else
      turn_off_light(buffer, i);
  }
}

typedef struct {
  u8 *grid;
  u8 y;
} context;

static inline void parse(char *line, context *const ctx) {
  for (u8 x = 0; x < 100; ++x)
    if (line[x] == '#')
      turn_on_light(ctx->grid, ctx->y * 100 + x);
  ctx->y++;
}

static inline u8 count_tl_corner(const u8 *const grid) {
  return light_state(grid, 1) + light_state(grid, 100) + light_state(grid, 101);
}

static inline u8 count_tr_corner(const u8 *const grid) {
  return light_state(grid, 98) + light_state(grid, 198) +
         light_state(grid, 199);
}

static inline u8 count_bl_corner(const u8 *const grid) {
  return light_state(grid, 800) + light_state(grid, 801) +
         light_state(grid, 901);
}

static inline u8 count_br_corner(const u8 *const grid) {
  return light_state(grid, 898) + light_state(grid, 899) +
         light_state(grid, 998);
}

static inline u8 count_left(const u8 *const grid, const u32 i) {
  return light_state(grid, i - 100) + light_state(grid, i - 100 + 1) +
         light_state(grid, i + 1) + light_state(grid, i + 100) +
         light_state(grid, i + 100 + 1);
}

static inline u8 count_right(const u8 *const grid, const u32 i) {
  return light_state(grid, i - 100) + light_state(grid, i - 100 - 1) +
         light_state(grid, i - 1) + light_state(grid, i + 100) +
         light_state(grid, i + 100 - 1);
}

static inline u8 count_top(const u8 *const grid, const u32 i) {
  return light_state(grid, i - 1) + light_state(grid, i + 1) +
         light_state(grid, i + 100 - 1) + light_state(grid, i + 100) +
         light_state(grid, i + 100 + 1);
}

static inline u8 count_bottom(const u8 *const grid, const u32 i) {
  return light_state(grid, i - 1) + light_state(grid, i + 1) +
         light_state(grid, i - 100 - 1) + light_state(grid, i - 100) +
         light_state(grid, i - 100 + 1);
}

static inline u8 count_center(const u8 *const grid, const u32 i) {
  return light_state(grid, i - 100 - 1) + light_state(grid, i - 100) +
         light_state(grid, i - 100 + 1) + light_state(grid, i - 1) +
         light_state(grid, i + 1) + light_state(grid, i + 100 - 1) +
         light_state(grid, i + 100) + light_state(grid, i + 100 + 1);
}

static void update(const u8 *const grid, u8 *const buffer, const bool part2) {
  // corners
  if (!part2) {
    set_light(grid, buffer, 0, count_tl_corner(grid));
    set_light(grid, buffer, 99, count_tr_corner(grid));
    set_light(grid, buffer, 9900, count_bl_corner(grid));
    set_light(grid, buffer, 9999, count_br_corner(grid));
  }

  // left and right
  for (u8 y = 1; y < 99; ++y) {
    set_light(grid, buffer, y * 100, count_left(grid, y * 100));
    set_light(grid, buffer, y * 100 + 99, count_right(grid, y * 100 + 99));
  }

  // top and bottom
  for (u8 x = 1; x < 99; ++x) {
    set_light(grid, buffer, x, count_top(grid, x));
    set_light(grid, buffer, 9900 + x, count_bottom(grid, 9900 + x));
  }

  // center
  for (u8 y = 1; y < 99; ++y)
    for (u8 x = 1; x < 99; ++x)
      set_light(grid, buffer, y * 100 + x, count_center(grid, y * 100 + x));
}

static inline u32 count_lights(const u8 *const grid) {
  u32 on = 0;
  for (u32 i = 0; i < GRID_SIZE; ++i)
    on += aoc_popcount(grid[i]);
  return on;
}

static inline void turn_on_corners(u8 *const grid) {
  turn_on_light(grid, 0);
  turn_on_light(grid, 99);
  turn_on_light(grid, 9900);
  turn_on_light(grid, 9999);
}

static u32 solve(const u8 *const grid, const bool part2) {
  u8 front[GRID_SIZE];
  u8 back[GRID_SIZE];
  aoc_mem_copy(front, grid, GRID_SIZE * sizeof(u8));
  if (part2) {
    turn_on_corners(front);
    turn_on_corners(back);
  }
  for (u8 i = 0; i < 50; ++i) {
    update(front, back, part2);
    update(back, front, part2);
  }
  return count_lights(front);
}

int main(void) {
  u8 grid[GRID_SIZE] = {0};
  context ctx = {.grid = grid};
  aoc_file_read_lines1("day18/input.txt", (aoc_line_func)parse, &ctx);
  printf("%u\n%u\n", solve(grid, false), solve(grid, true));
}
