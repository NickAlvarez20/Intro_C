#include <stdio.h>

int main(void){
    // Write a program to copy its input to its output, replacing each string of one or more blanks by a single blank.
    int c, last_char = '\0';
    while ((c = getchar()) != EOF){
        if (c == ' ' && last_char == ' '){
            continue;
        }
        last_char = c;
        putchar(c);
    }
    return 0;
}

