#include <stdio.h>
#define INPUT 34000000
#define MAX 1000000

static int solve(const int steps, const int multiplier) {
  int presents[MAX] = {0};
  for (int i = 1; i < MAX; i++) {
    int j = i;
    int step = 0;
    while (j < MAX && step < steps) {
      presents[j] += multiplier * i;
      j += i;
      step++;
    }
  }

  for (int i = 0; i < MAX; i++) {
    if (presents[i] >= INPUT)
      return i;
  }
  return -1;
}

int main(void) {
  printf("%d\n%d\n", solve(MAX, 10), solve(50, 11));
}
