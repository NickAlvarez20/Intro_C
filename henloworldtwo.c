#include <stdio.h>

#define LOWER 0 /* lower limit of temp table*/
#define UPPER 300 /* upper limit of temp table */
#define STEP 20 /* step size */

main() {
  int fahr;
  for (fahr = UPPER; fahr >= LOWER; fahr = fahr - STEP) {
    printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
  }
  return 0;
}
