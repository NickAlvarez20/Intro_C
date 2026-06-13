#include <stdio.h>

/* copy input to output; 1st version */

int main(void) {
    int c;

    while (printf("%d ", (c = getchar()) != EOF), c != EOF)
        putchar(c);
    putchar('\n');
    return 0; 
}