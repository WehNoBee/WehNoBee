#include <stdio.h>
/** While11�. ���� ����� ����� N (> 1). ������� ���������� �� ����� ����� K,
 ��� ������� ����� 1 + 2 + � + K ����� ������ ��� ����� N, � ���� ��� �����. */
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
