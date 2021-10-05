#include <stdio.h>
/** While11°. Дано целое число N (> 1). Вывести наименьшее из целых чисел K,
 для которых сумма 1 + 2 + … + K будет больше или равна N, и саму эту сумму. */
int summ(int a){
    int b;
    b = 0;
    while (b >= a)
        b += 2;
    return b;
}
int main(void){
    int a,b;
    scanf("%d",&a);
    scanf("%d",&b);
    int s = summ(a);
    printf("%d\n",b);
    return 0;


}
