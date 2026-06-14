#include <stdio.h>

/* count characters in input; 1st version */

int main(void){
  double nc;

  for (nc = 0; getchar() != EOF; ++nc)
    ; // null statement

  printf("%.0f\n", nc);
  return 0;
}


/* create a program that creates an array, allocates memory for it, and fills it with the numbers 1 to 100, and then frees the memory*/

// int main(void){
//   int *array = malloc(100 * sizeof(int));
//   if (array == NULL){
//     printf("Error: Failed to allocate memory\n");
//     return 1;
//   }
//   for (int i = 0; i < 100; i++){
//     array[i] = i + 1;
//   }
    
  /* use the array here if you want, e.g. printf("%d\n", array[0]); */

//   free(array);
//   return 0;
// }