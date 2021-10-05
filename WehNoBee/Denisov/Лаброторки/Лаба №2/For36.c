#include <stdio.h>
int main(){
int n,k,summ1,i,promej;
printf("Kakoe choslo posl:");
scanf("%d",&n);
printf("Na chto umnoj:");
scanf("%d",&k);
summ1 = 0;
promej = 0;
for(i = 1;i!=n;i++){
    promej = i * k;
    summ1 = summ1 + promej;
    printf("Promejutok:%d",promej);
printf("Summa:%d",summ1);
return 0;
}
}
