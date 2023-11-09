#include <stdio.h>
#include <aoc/aoc.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

typedef struct {
  char data[2];
} identifier;

typedef enum {
  op_not_r,
  op_assign_i,
  op_assign_r,
  op_and_ir,
  op_and_rr,
  op_or_ir,
  op_or_rr,
  op_rshift_ri,
  op_lshift_ri,
} opcode;

typedef union {
  identifier reg;
  u16 imm;
} operand;

typedef struct {
  opcode code;
  operand op1;
  operand op2;
} instruction;

static inline u32 identifier_hash(const identifier *const i) {
  return ((36591911u * (u32)i->data[0]) << 5u ^ (97337263u * (u32)i->data[1])
                                                    << 5u);
}

static inline bool identifier_equals(const identifier *const a,
                                     const identifier *const b) {
  return (a->data[0] == b->data[0] && a->data[1] == b->data[1]);
}

#define AOC_KEY_T identifier
#define AOC_KEY_T_NAME id
#define AOC_VALUE_T instruction
#define AOC_VALUE_T_NAME instr
#define AOC_KEY_T_EMPTY ((identifier){0})
#define AOC_KEY_T_HASH identifier_hash
#define AOC_KEY_T_EQUALS identifier_equals
#define AOC_BASE2_CAPACITY
#include <aoc/map.h>

static identifier parse_identifier(char *line, char **out) {
  i8 len = 0;
  identifier i = {0};
  while (!isspace(*line) && *line != '\0') {
    i.data[len++] = *line;
    line++;
  }
  *out = line;
  return i;
}

static void parse(char *line, void *map) {
  identifier key = {0};
  instruction instr = {0};
  if (line[0] == 'N') /* NOT */ {
    instr.code = op_not_r;
    instr.op1.reg = parse_identifier(line + 4, &line);
    key = parse_identifier(line + 4, &line);
  } else {
    bool firstOpIsIdentifier = true;
    if (line[0] >= 'a' && line[0] <= 'z') {
      instr.op1.reg = parse_identifier(line, &line);
    } else /* 0-9 */ {
      instr.op1.imm = strtoul(line, &line, 10);
      firstOpIsIdentifier = false;
    }
    line++;
    switch (line[0]) {
    case '-':
      instr.code = op_assign_i + firstOpIsIdentifier;
      key = parse_identifier(line + 3, &line);
      break;
    case 'R':
      instr.code = op_rshift_ri;
      instr.op2.imm = strtoul(line + 7, &line, 10);
      key = parse_identifier(line + 4, &line);
      break;
    case 'L':
      instr.code = op_lshift_ri;
      instr.op2.imm = strtoul(line + 7, &line, 10);
      key = parse_identifier(line + 4, &line);
      break;
    case 'O':
      instr.code = op_or_ir + firstOpIsIdentifier;
      instr.op2.reg = parse_identifier(line + 3, &line);
      key = parse_identifier(line + 4, &line);
      break;
    case 'A':
      instr.code = op_and_ir + firstOpIsIdentifier;
      instr.op2.reg = parse_identifier(line + 4, &line);
      key = parse_identifier(line + 4, &line);
      break;
    }
  }
  aoc_map_id_instr_put(map, key, instr);
}

static u16 evaluate(aoc_map_id_instr *const map, const identifier id) {
  u16 value = 0;
  instruction instr = {0};
  aoc_map_id_instr_get(map, id, &instr);
  switch (instr.code) {
  case op_not_r:
    value = ~(evaluate(map, instr.op1.reg));
    break;
  case op_assign_i:
    value = instr.op1.imm;
    break;
  case op_assign_r:
    value = evaluate(map, instr.op1.reg);
    break;
  case op_and_ir:
    value = instr.op1.imm & evaluate(map, instr.op2.reg);
    break;
  case op_and_rr:
    value = evaluate(map, instr.op1.reg) & evaluate(map, instr.op2.reg);
    break;
  case op_or_ir:
    value = instr.op1.imm | evaluate(map, instr.op2.reg);
    break;
  case op_or_rr:
    value = evaluate(map, instr.op1.reg) | evaluate(map, instr.op2.reg);
    break;
  case op_rshift_ri:
    value = evaluate(map, instr.op1.reg) >> instr.op2.imm;
    break;
  case op_lshift_ri:
    value = evaluate(map, instr.op1.reg) << instr.op2.imm;
    break;
  }

  // overwrite instruction with an assignment instruction so it doesn't have to
  // be evaluated again
  aoc_map_id_instr_put(map, id,
                       (instruction){
                           .code = op_assign_i,
                           .op1.imm = value,
                       });
  return value;
}

int main(void) {
  aoc_map_id_instr map = {0};
  aoc_map_id_instr_create(&map, 1 << 9);
  aoc_file_read_lines1("day07/input.txt", parse, &map);

  aoc_map_id_instr clone = {0};
  aoc_map_id_instr_duplicate(&clone, &map);

  const u16 part1 = evaluate(&map, (identifier){.data = "a"});

  // overwrite wire b with the result of part1 and restart
  aoc_map_id_instr_put(&clone, (identifier){.data = "b"},
                       (instruction){
                           .code = op_assign_i,
                           .op1.imm = part1,
                       });
  const u16 part2 = evaluate(&clone, (identifier){.data = "a"});

  printf("%u\n%u\n", part1, part2);

  aoc_map_id_instr_destroy(&map);
  aoc_map_id_instr_destroy(&clone);
}
