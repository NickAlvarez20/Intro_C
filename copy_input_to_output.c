#include <stdio.h>

/* copy input to output; 1st version */

/* LEARNING (Track B — optional, C/K&R):
 * Exercise 1.7: Count spaces, tabs, and newlines in input.
 * Exercise 1.8: Write a program to copy input to output, replacing each
 *               string of one or more blanks by a single blank.
 * Exercise 1.9: Write a program to copy input to output, replacing each
 *               sequence of one or more newline characters by a single newline.
 * See LEARNING.md Track B for verify commands.
 */

int main(void) {
    int c;

    while (printf("%d ", (c = getchar()) != EOF), c != EOF)
        putchar(c);
    putchar('\n');
    return 0; 
}