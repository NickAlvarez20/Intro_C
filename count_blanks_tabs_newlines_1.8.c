#include <stdio.h>

int main(void) {
  int c, nl, nt, nb;
  nl = nt = nb = 0;
  while ((c = getchar()) != EOF) {
    if (c == '\n') {
      ++nl;
    }
    if (c == '\t') {
      ++nt;
    }
    if (c == ' ') {
      ++nb;
    }
  }
  printf("%d %d %d\n", nl, nt, nb);
  return 0;
}