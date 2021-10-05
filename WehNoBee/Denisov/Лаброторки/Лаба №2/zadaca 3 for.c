/*
#include <stdio.h>
int main()
{
int a,b,i;
i = 0;
scanf("%d%d", &a,&b);
for(;a < b;)
{
    a = a + 1;
    i = i + a;
}
printf("%d", a);
return 0;
}
*/
#include <stdio.h>
int main(void){
 int a,b,i=0;
 printf("A=");
 scanf("%d",&a);
 printf("B=");
 scanf("%d",&b);
 for(;a <= b; a = a + 1)
{
    i += a;
}
printf("%d", i);
return 0;
 }
