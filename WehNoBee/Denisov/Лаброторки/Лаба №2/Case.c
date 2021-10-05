#include<stdio.h>
int main(){
char st;
printf("Vvedite chislo dniya\n");
st = getchar();
switch(st){
case '1':
    printf("Ponedelnik\n");
    break;
case '2':
    printf("Vtornik\n");
    break;
case '3':
    printf("Sreda\n");
    break;
case '4':
    printf("Chetverg\n");
    break;
case '5':
    printf("Pyatnica\n");
    break;
case '6':
    printf("Subbota\n");
    break;
case '7':
    printf("Voskresenie\n");
    break;
}
return 0;
}
